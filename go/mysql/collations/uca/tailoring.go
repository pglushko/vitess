/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package uca

type WeightPatch struct {
	Codepoint rune
	Patch     []uint16
}

type Reorder struct {
	FromMin, FromMax uint16
	ToMin, ToMax     uint16
}

const MaxCodepoint = 0x10FFFF + 1
const CodepointsPerPage = 256
const Pages = MaxCodepoint / CodepointsPerPage
const MaxCollationElementsPerCodepoint = 8

type patcher interface {
	alloc(original *[]uint16, patches []WeightPatch) []uint16
	apply(page []uint16, offset int, weights []uint16)
}

type patcherUCA900 struct{}

func (patcherUCA900) alloc(original *[]uint16, patches []WeightPatch) []uint16 {
	var maxWeights int
	for _, p := range patches {
		if len(p.Patch)%3 != 0 {
			panic("len(p.Patch)%3")
		}
		if len(p.Patch) > maxWeights {
			maxWeights = len(p.Patch)
		}
	}

	minLenForPage := maxWeights*CodepointsPerPage + CodepointsPerPage
	if original == nil {
		return make([]uint16, minLenForPage)
	}
	if len(*original) > minLenForPage {
		minLenForPage = len(*original)
	}
	newPage := make([]uint16, minLenForPage)
	copy(newPage, *original)
	return newPage
}

func (patcherUCA900) apply(page []uint16, offset int, weights []uint16) {
	var weightcount = len(weights) / 3
	var ce int
	for ce < weightcount {
		page[(ce*3+1)*CodepointsPerPage+offset] = weights[ce*3+0]
		page[(ce*3+2)*CodepointsPerPage+offset] = weights[ce*3+1]
		page[(ce*3+3)*CodepointsPerPage+offset] = weights[ce*3+2]
		ce++
	}

	for ce < int(page[offset]) {
		page[(ce*3+1)*CodepointsPerPage+offset] = 0x0
		page[(ce*3+2)*CodepointsPerPage+offset] = 0x0
		page[(ce*3+3)*CodepointsPerPage+offset] = 0x0
		ce++
	}

	page[offset] = uint16(weightcount)
}

type patcherUCALegacy struct{}

func (patcherUCALegacy) alloc(original *[]uint16, patches []WeightPatch) []uint16 {
	var newStride int
	for _, p := range patches {
		if len(p.Patch) > newStride {
			newStride = len(p.Patch)
		}
	}

	minLenForPage := 1 + newStride*CodepointsPerPage
	if original == nil {
		return make([]uint16, minLenForPage)
	}

	if len(*original) >= minLenForPage {
		newPage := make([]uint16, len(*original))
		copy(newPage, *original)
		return newPage
	}

	originalStride := int((*original)[0])
	if originalStride >= newStride {
		panic("mismatch in originalStride calculation?")
	}

	newPage := make([]uint16, minLenForPage)
	for i := 0; i < CodepointsPerPage; i++ {
		for j := 0; j < originalStride; j++ {
			newPage[1+i*newStride] = (*original)[1+i*originalStride]
		}
	}
	return newPage
}

func (patcherUCALegacy) apply(page []uint16, offset int, weights []uint16) {
	stride := int(page[0])
	var ce int
	for ce < len(weights) {
		page[1+offset*stride+ce] = weights[ce]
		ce++
	}
	for ce < stride {
		page[1+offset*stride+ce] = 0x0
		ce++
	}
}

func applyTailoring(patcher patcher, base WeightTable, patches []WeightPatch) WeightTable {
	if len(patches) == 0 {
		return base
	}

	result := make(WeightTable, len(base))
	copy(result, base)

	groups := make(map[int][]WeightPatch)
	for _, patch := range patches {
		p, _ := pageOffset(patch.Codepoint)
		groups[p] = append(groups[p], patch)
	}

	for p, pps := range groups {
		page := patcher.alloc(result[p], pps)

		for _, patch := range pps {
			_, off := pageOffset(patch.Codepoint)
			patcher.apply(page, off, patch.Patch)
		}

		result[p] = &page
	}

	return result
}

type parametricT struct {
	upperCaseFirst bool
	reorder        []Reorder
	reorderMin     uint16
	reorderMax     uint16
}

func (p *parametricT) adjust(level int, weight uint16) uint16 {
	if p == nil {
		return weight
	}
	switch level {
	case 0:
		if weight >= p.reorderMin && weight <= p.reorderMax {
			for _, reorder := range p.reorder {
				if weight >= reorder.FromMin && weight <= reorder.FromMax {
					return weight - reorder.FromMin + reorder.ToMin
				}
			}
		}
	case 2:
		// see: https://unicode.org/reports/tr35/tr35-collation.html#Case_Untailored
		if p.upperCaseFirst && weight < ' ' {
			switch weight {
			case 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0E, 0x11, 0x12, 0x1D:
				return weight | 0x100
			default:
				return weight | 0x300
			}
		}
	}
	return weight
}

func newParametricTailoring(reorder []Reorder, upperCaseFirst bool) *parametricT {
	if len(reorder) == 0 && !upperCaseFirst {
		return nil
	}

	t := &parametricT{
		upperCaseFirst: upperCaseFirst,
		reorder:        reorder,
		reorderMin:     ^uint16(0),
		reorderMax:     0,
	}

	for _, r := range reorder {
		if r.FromMin < t.reorderMin {
			t.reorderMin = r.FromMin
		}
		if r.FromMax > t.reorderMax {
			t.reorderMax = r.FromMax
		}
	}

	return t
}
