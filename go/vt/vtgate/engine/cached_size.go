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
// Code generated by Sizegen. DO NOT EDIT.

package engine

import (
	"math"
	"reflect"
	"unsafe"
)

type cachedObject interface {
	CachedSize(alloc bool) int64
}

func (cached *AggregateParams) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(48)
	}
	// field Alias string
	size += int64(len(cached.Alias))
	// field Expr vitess.io/vitess/go/vt/sqlparser.Expr
	if cc, ok := cached.Expr.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *AlterVSchema) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field AlterVschemaDDL *vitess.io/vitess/go/vt/sqlparser.AlterVschema
	size += cached.AlterVschemaDDL.CachedSize(true)
	return size
}
func (cached *Concatenate) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	// field Sources []vitess.io/vitess/go/vt/vtgate/engine.Primitive
	{
		size += int64(cap(cached.Sources)) * int64(16)
		for _, elem := range cached.Sources {
			if cc, ok := elem.(cachedObject); ok {
				size += cc.CachedSize(true)
			}
		}
	}
	return size
}
func (cached *DBDDL) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *DDL) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(59)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field SQL string
	size += int64(len(cached.SQL))
	// field DDL vitess.io/vitess/go/vt/sqlparser.DDLStatement
	if cc, ok := cached.DDL.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field NormalDDL *vitess.io/vitess/go/vt/vtgate/engine.Send
	size += cached.NormalDDL.CachedSize(true)
	// field OnlineDDL *vitess.io/vitess/go/vt/vtgate/engine.OnlineDDL
	size += cached.OnlineDDL.CachedSize(true)
	return size
}
func (cached *DML) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(144)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Query string
	size += int64(len(cached.Query))
	// field Vindex vitess.io/vitess/go/vt/vtgate/vindexes.SingleColumn
	if cc, ok := cached.Vindex.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Values []vitess.io/vitess/go/sqltypes.PlanValue
	{
		size += int64(cap(cached.Values)) * int64(88)
		for _, elem := range cached.Values {
			size += elem.CachedSize(false)
		}
	}
	// field KsidVindex vitess.io/vitess/go/vt/vtgate/vindexes.SingleColumn
	if cc, ok := cached.KsidVindex.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Table *vitess.io/vitess/go/vt/vtgate/vindexes.Table
	size += cached.Table.CachedSize(true)
	// field OwnedVindexQuery string
	size += int64(len(cached.OwnedVindexQuery))
	return size
}
func (cached *Delete) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(144)
	}
	// field DML vitess.io/vitess/go/vt/vtgate/engine.DML
	size += cached.DML.CachedSize(false)
	return size
}
func (cached *Distinct) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field Source vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Source.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Generate) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(112)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field Query string
	size += int64(len(cached.Query))
	// field Values vitess.io/vitess/go/sqltypes.PlanValue
	size += cached.Values.CachedSize(false)
	return size
}
func (cached *GroupByParams) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field Expr vitess.io/vitess/go/vt/sqlparser.Expr
	if cc, ok := cached.Expr.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Insert) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(144)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field Query string
	size += int64(len(cached.Query))
	// field VindexValues []vitess.io/vitess/go/sqltypes.PlanValue
	{
		size += int64(cap(cached.VindexValues)) * int64(88)
		for _, elem := range cached.VindexValues {
			size += elem.CachedSize(false)
		}
	}
	// field Table *vitess.io/vitess/go/vt/vtgate/vindexes.Table
	size += cached.Table.CachedSize(true)
	// field Generate *vitess.io/vitess/go/vt/vtgate/engine.Generate
	size += cached.Generate.CachedSize(true)
	// field Prefix string
	size += int64(len(cached.Prefix))
	// field Mid []string
	{
		size += int64(cap(cached.Mid)) * int64(16)
		for _, elem := range cached.Mid {
			size += int64(len(elem))
		}
	}
	// field Suffix string
	size += int64(len(cached.Suffix))
	return size
}

//go:nocheckptr
func (cached *Join) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(72)
	}
	// field Left vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Left.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Right vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Right.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Cols []int
	{
		size += int64(cap(cached.Cols)) * int64(8)
	}
	// field Vars map[string]int
	if cached.Vars != nil {
		size += int64(48)
		hmap := reflect.ValueOf(cached.Vars)
		numBuckets := int(math.Pow(2, float64((*(*uint8)(unsafe.Pointer(hmap.Pointer() + uintptr(9)))))))
		numOldBuckets := (*(*uint16)(unsafe.Pointer(hmap.Pointer() + uintptr(10))))
		size += int64(numOldBuckets * 208)
		if len(cached.Vars) > 0 || numBuckets > 1 {
			size += int64(numBuckets * 208)
		}
		for k := range cached.Vars {
			size += int64(len(k))
		}
	}
	return size
}
func (cached *Limit) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(192)
	}
	// field Count vitess.io/vitess/go/sqltypes.PlanValue
	size += cached.Count.CachedSize(false)
	// field Offset vitess.io/vitess/go/sqltypes.PlanValue
	size += cached.Offset.CachedSize(false)
	// field Input vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Input.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Lock) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(40)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Query string
	size += int64(len(cached.Query))
	return size
}
func (cached *MStream) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(40)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field TableName string
	size += int64(len(cached.TableName))
	return size
}
func (cached *MemorySort) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(136)
	}
	// field UpperLimit vitess.io/vitess/go/sqltypes.PlanValue
	size += cached.UpperLimit.CachedSize(false)
	// field OrderBy []vitess.io/vitess/go/vt/vtgate/engine.OrderByParams
	{
		size += int64(cap(cached.OrderBy)) * int64(32)
	}
	// field Input vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Input.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *MergeSort) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(49)
	}
	// field Primitives []vitess.io/vitess/go/vt/vtgate/engine.StreamExecutor
	{
		size += int64(cap(cached.Primitives)) * int64(16)
		for _, elem := range cached.Primitives {
			if cc, ok := elem.(cachedObject); ok {
				size += cc.CachedSize(true)
			}
		}
	}
	// field OrderBy []vitess.io/vitess/go/vt/vtgate/engine.OrderByParams
	{
		size += int64(cap(cached.OrderBy)) * int64(32)
	}
	return size
}
func (cached *OnlineDDL) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(64)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field DDL vitess.io/vitess/go/vt/sqlparser.DDLStatement
	if cc, ok := cached.DDL.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field SQL string
	size += int64(len(cached.SQL))
	// field DDLStrategySetting *vitess.io/vitess/go/vt/schema.DDLStrategySetting
	size += cached.DDLStrategySetting.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *OrderedAggregate) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(80)
	}
	// field Aggregates []vitess.io/vitess/go/vt/vtgate/engine.AggregateParams
	{
		size += int64(cap(cached.Aggregates)) * int64(48)
		for _, elem := range cached.Aggregates {
			size += elem.CachedSize(false)
		}
	}
	// field GroupByKeys []vitess.io/vitess/go/vt/vtgate/engine.GroupByParams
	{
		size += int64(cap(cached.GroupByKeys)) * int64(32)
		for _, elem := range cached.GroupByKeys {
			size += elem.CachedSize(false)
		}
	}
	// field Input vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Input.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Plan) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(120)
	}
	// field Original string
	size += int64(len(cached.Original))
	// field Instructions vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Instructions.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field BindVarNeeds *vitess.io/vitess/go/vt/sqlparser.BindVarNeeds
	size += cached.BindVarNeeds.CachedSize(true)
	// field Warnings []*vitess.io/vitess/go/vt/proto/query.QueryWarning
	{
		size += int64(cap(cached.Warnings)) * int64(8)
		for _, elem := range cached.Warnings {
			size += elem.CachedSize(true)
		}
	}
	return size
}
func (cached *Projection) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(64)
	}
	// field Cols []string
	{
		size += int64(cap(cached.Cols)) * int64(16)
		for _, elem := range cached.Cols {
			size += int64(len(elem))
		}
	}
	// field Exprs []vitess.io/vitess/go/vt/vtgate/evalengine.Expr
	{
		size += int64(cap(cached.Exprs)) * int64(16)
		for _, elem := range cached.Exprs {
			if cc, ok := elem.(cachedObject); ok {
				size += cc.CachedSize(true)
			}
		}
	}
	// field Input vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Input.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *PulloutSubquery) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(72)
	}
	// field SubqueryResult string
	size += int64(len(cached.SubqueryResult))
	// field HasValues string
	size += int64(len(cached.HasValues))
	// field Subquery vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Subquery.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Underlying vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Underlying.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *RenameFields) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(64)
	}
	// field Cols []string
	{
		size += int64(cap(cached.Cols)) * int64(16)
		for _, elem := range cached.Cols {
			size += int64(len(elem))
		}
	}
	// field Indices []int
	{
		size += int64(cap(cached.Indices)) * int64(8)
	}
	// field Input vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Input.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *ReplaceVariables) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field Input vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Input.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *RevertMigration) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(48)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field Stmt *vitess.io/vitess/go/vt/sqlparser.RevertMigration
	size += cached.Stmt.CachedSize(true)
	// field Query string
	size += int64(len(cached.Query))
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Route) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(224)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Query string
	size += int64(len(cached.Query))
	// field TableName string
	size += int64(len(cached.TableName))
	// field FieldQuery string
	size += int64(len(cached.FieldQuery))
	// field Vindex vitess.io/vitess/go/vt/vtgate/vindexes.SingleColumn
	if cc, ok := cached.Vindex.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Values []vitess.io/vitess/go/sqltypes.PlanValue
	{
		size += int64(cap(cached.Values)) * int64(88)
		for _, elem := range cached.Values {
			size += elem.CachedSize(false)
		}
	}
	// field OrderBy []vitess.io/vitess/go/vt/vtgate/engine.OrderByParams
	{
		size += int64(cap(cached.OrderBy)) * int64(32)
	}
	// field SysTableTableSchema []vitess.io/vitess/go/vt/vtgate/evalengine.Expr
	{
		size += int64(cap(cached.SysTableTableSchema)) * int64(16)
		for _, elem := range cached.SysTableTableSchema {
			if cc, ok := elem.(cachedObject); ok {
				size += cc.CachedSize(true)
			}
		}
	}
	// field SysTableTableName []vitess.io/vitess/go/vt/vtgate/evalengine.Expr
	{
		size += int64(cap(cached.SysTableTableName)) * int64(16)
		for _, elem := range cached.SysTableTableName {
			if cc, ok := elem.(cachedObject); ok {
				size += cc.CachedSize(true)
			}
		}
	}
	return size
}
func (cached *Rows) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(48)
	}
	// field rows [][]vitess.io/vitess/go/sqltypes.Value
	{
		size += int64(cap(cached.rows)) * int64(24)
		for _, elem := range cached.rows {
			{
				size += int64(cap(elem)) * int64(32)
				for _, elem := range elem {
					size += elem.CachedSize(false)
				}
			}
		}
	}
	// field fields []*vitess.io/vitess/go/vt/proto/query.Field
	{
		size += int64(cap(cached.fields)) * int64(8)
		for _, elem := range cached.fields {
			size += elem.CachedSize(true)
		}
	}
	return size
}
func (cached *SQLCalcFoundRows) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field LimitPrimitive vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.LimitPrimitive.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field CountPrimitive vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.CountPrimitive.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Send) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(44)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Query string
	size += int64(len(cached.Query))
	return size
}
func (cached *SessionPrimitive) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *Set) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(40)
	}
	// field Ops []vitess.io/vitess/go/vt/vtgate/engine.SetOp
	{
		size += int64(cap(cached.Ops)) * int64(16)
		for _, elem := range cached.Ops {
			if cc, ok := elem.(cachedObject); ok {
				size += cc.CachedSize(true)
			}
		}
	}
	// field Input vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Input.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Subquery) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(40)
	}
	// field Cols []int
	{
		size += int64(cap(cached.Cols)) * int64(8)
	}
	// field Subquery vitess.io/vitess/go/vt/vtgate/engine.Primitive
	if cc, ok := cached.Subquery.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *SysVarCheckAndIgnore) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(56)
	}
	// field Name string
	size += int64(len(cached.Name))
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Expr string
	size += int64(len(cached.Expr))
	return size
}
func (cached *SysVarIgnore) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field Name string
	size += int64(len(cached.Name))
	// field Expr string
	size += int64(len(cached.Expr))
	return size
}
func (cached *SysVarReservedConn) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(56)
	}
	// field Name string
	size += int64(len(cached.Name))
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Expr string
	size += int64(len(cached.Expr))
	return size
}
func (cached *SysVarSetAware) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field Name string
	size += int64(len(cached.Name))
	// field Expr vitess.io/vitess/go/vt/vtgate/evalengine.Expr
	if cc, ok := cached.Expr.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}

//go:nocheckptr
func (cached *Update) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(152)
	}
	// field DML vitess.io/vitess/go/vt/vtgate/engine.DML
	size += cached.DML.CachedSize(false)
	// field ChangedVindexValues map[string]*vitess.io/vitess/go/vt/vtgate/engine.VindexValues
	if cached.ChangedVindexValues != nil {
		size += int64(48)
		hmap := reflect.ValueOf(cached.ChangedVindexValues)
		numBuckets := int(math.Pow(2, float64((*(*uint8)(unsafe.Pointer(hmap.Pointer() + uintptr(9)))))))
		numOldBuckets := (*(*uint16)(unsafe.Pointer(hmap.Pointer() + uintptr(10))))
		size += int64(numOldBuckets * 208)
		if len(cached.ChangedVindexValues) > 0 || numBuckets > 1 {
			size += int64(numBuckets * 208)
		}
		for k, v := range cached.ChangedVindexValues {
			size += int64(len(k))
			size += v.CachedSize(true)
		}
	}
	return size
}
func (cached *UpdateTarget) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field Target string
	size += int64(len(cached.Target))
	return size
}
func (cached *UserDefinedVariable) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field Name string
	size += int64(len(cached.Name))
	// field Expr vitess.io/vitess/go/vt/vtgate/evalengine.Expr
	if cc, ok := cached.Expr.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *VStream) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(64)
	}
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field TargetDestination vitess.io/vitess/go/vt/key.Destination
	if cc, ok := cached.TargetDestination.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field TableName string
	size += int64(len(cached.TableName))
	// field Position string
	size += int64(len(cached.Position))
	return size
}
func (cached *VindexFunc) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(160)
	}
	// field Fields []*vitess.io/vitess/go/vt/proto/query.Field
	{
		size += int64(cap(cached.Fields)) * int64(8)
		for _, elem := range cached.Fields {
			size += elem.CachedSize(true)
		}
	}
	// field Cols []int
	{
		size += int64(cap(cached.Cols)) * int64(8)
	}
	// field Vindex vitess.io/vitess/go/vt/vtgate/vindexes.SingleColumn
	if cc, ok := cached.Vindex.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	// field Value vitess.io/vitess/go/sqltypes.PlanValue
	size += cached.Value.CachedSize(false)
	return size
}

//go:nocheckptr
func (cached *VindexValues) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field PvMap map[string]vitess.io/vitess/go/sqltypes.PlanValue
	if cached.PvMap != nil {
		size += int64(48)
		hmap := reflect.ValueOf(cached.PvMap)
		numBuckets := int(math.Pow(2, float64((*(*uint8)(unsafe.Pointer(hmap.Pointer() + uintptr(9)))))))
		numOldBuckets := (*(*uint16)(unsafe.Pointer(hmap.Pointer() + uintptr(10))))
		size += int64(numOldBuckets * 848)
		if len(cached.PvMap) > 0 || numBuckets > 1 {
			size += int64(numBuckets * 848)
		}
		for k, v := range cached.PvMap {
			size += int64(len(k))
			size += v.CachedSize(false)
		}
	}
	return size
}

//go:nocheckptr
func (cached *shardRoute) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field query string
	size += int64(len(cached.query))
	// field rs *vitess.io/vitess/go/vt/srvtopo.ResolvedShard
	size += cached.rs.CachedSize(true)
	// field bv map[string]*vitess.io/vitess/go/vt/proto/query.BindVariable
	if cached.bv != nil {
		size += int64(48)
		hmap := reflect.ValueOf(cached.bv)
		numBuckets := int(math.Pow(2, float64((*(*uint8)(unsafe.Pointer(hmap.Pointer() + uintptr(9)))))))
		numOldBuckets := (*(*uint16)(unsafe.Pointer(hmap.Pointer() + uintptr(10))))
		size += int64(numOldBuckets * 208)
		if len(cached.bv) > 0 || numBuckets > 1 {
			size += int64(numBuckets * 208)
		}
		for k, v := range cached.bv {
			size += int64(len(k))
			size += v.CachedSize(true)
		}
	}
	return size
}
