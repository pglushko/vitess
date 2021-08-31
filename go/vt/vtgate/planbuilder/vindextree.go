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

package planbuilder

import (
	"vitess.io/vitess/go/vt/sqlparser"
	"vitess.io/vitess/go/vt/vtgate/semantics"
)

type vindexTree struct {
	table routeTable
}

func (v vindexTree) tableID() semantics.TableSet {
	panic("implement me")
}

func (v vindexTree) clone() queryTree {
	panic("implement me")
}

func (v vindexTree) cost() int {
	panic("implement me")
}

func (v vindexTree) pushOutputColumns(names []*sqlparser.ColName, semTable *semantics.SemTable) ([]int, error) {
	panic("implement me")
}
