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

package sqlparser

import (
	"strings"

	"vitess.io/vitess/go/sqltypes"
)

// formatFast formats the node.
func (node *Select) formatFast(buf *TrackedBuffer) {
	buf.WriteString("select ")
	node.Comments.formatFast(buf)

	if node.Distinct {
		buf.WriteString(DistinctStr)
	}
	if node.Cache != nil {
		if *node.Cache {
			buf.WriteString(SQLCacheStr)
		} else {
			buf.WriteString(SQLNoCacheStr)
		}
	}
	if node.StraightJoinHint {
		buf.WriteString(StraightJoinHint)
	}
	if node.SQLCalcFoundRows {
		buf.WriteString(SQLCalcFoundRowsStr)
	}

	node.SelectExprs.formatFast(buf)
	buf.WriteString(" from ")

	node.From.formatFast(buf)

	node.Where.formatFast(buf)

	node.GroupBy.formatFast(buf)

	node.Having.formatFast(buf)

	node.OrderBy.formatFast(buf)

	node.Limit.formatFast(buf)
	buf.WriteString(node.Lock.ToString())
	node.Into.formatFast(buf)

}

// formatFast formats the node.
func (node *ParenSelect) formatFast(buf *TrackedBuffer) {
	buf.WriteByte('(')
	node.Select.formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *Union) formatFast(buf *TrackedBuffer) {
	node.FirstStatement.formatFast(buf)
	for _, us := range node.UnionSelects {
		us.formatFast(buf)
	}
	node.OrderBy.formatFast(buf)
	node.Limit.formatFast(buf)
	buf.WriteString(node.Lock.ToString())
}

// formatFast formats the node.
func (node *UnionSelect) formatFast(buf *TrackedBuffer) {
	if node.Distinct {
		buf.WriteByte(' ')
		buf.WriteString(UnionStr)
		buf.WriteByte(' ')
		node.Statement.formatFast(buf)
	} else {
		buf.WriteByte(' ')
		buf.WriteString(UnionAllStr)
		buf.WriteByte(' ')
		node.Statement.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *VStream) formatFast(buf *TrackedBuffer) {
	buf.WriteString("vstream ")
	node.Comments.formatFast(buf)
	node.SelectExpr.formatFast(buf)
	buf.WriteString(" from ")
	node.Table.formatFast(buf)

}

// formatFast formats the node.
func (node *Stream) formatFast(buf *TrackedBuffer) {
	buf.WriteString("stream ")
	node.Comments.formatFast(buf)
	node.SelectExpr.formatFast(buf)
	buf.WriteString(" from ")
	node.Table.formatFast(buf)

}

// formatFast formats the node.
func (node *Insert) formatFast(buf *TrackedBuffer) {
	switch node.Action {
	case InsertAct:
		buf.WriteString(InsertStr)
		buf.WriteByte(' ')

		node.Comments.formatFast(buf)
		buf.WriteString(node.Ignore.ToString())
		buf.WriteString("into ")

		node.Table.formatFast(buf)

		node.Partitions.formatFast(buf)

		node.Columns.formatFast(buf)
		buf.WriteByte(' ')

		node.Rows.formatFast(buf)

		node.OnDup.formatFast(buf)

	case ReplaceAct:
		buf.WriteString(ReplaceStr)
		buf.WriteByte(' ')

		node.Comments.formatFast(buf)
		buf.WriteString(node.Ignore.ToString())
		buf.WriteString("into ")

		node.Table.formatFast(buf)

		node.Partitions.formatFast(buf)

		node.Columns.formatFast(buf)
		buf.WriteByte(' ')

		node.Rows.formatFast(buf)

		node.OnDup.formatFast(buf)

	default:
		buf.WriteString("Unkown Insert Action")
		buf.WriteByte(' ')

		node.Comments.formatFast(buf)
		buf.WriteString(node.Ignore.ToString())
		buf.WriteString("into ")

		node.Table.formatFast(buf)

		node.Partitions.formatFast(buf)

		node.Columns.formatFast(buf)
		buf.WriteByte(' ')

		node.Rows.formatFast(buf)

		node.OnDup.formatFast(buf)

	}

}

// formatFast formats the node.
func (node *Update) formatFast(buf *TrackedBuffer) {
	buf.WriteString("update ")
	node.Comments.formatFast(buf)
	buf.WriteString(node.Ignore.ToString())
	node.TableExprs.formatFast(buf)
	buf.WriteString(" set ")

	node.Exprs.formatFast(buf)

	node.Where.formatFast(buf)

	node.OrderBy.formatFast(buf)

	node.Limit.formatFast(buf)

}

// formatFast formats the node.
func (node *Delete) formatFast(buf *TrackedBuffer) {
	buf.WriteString("delete ")
	node.Comments.formatFast(buf)
	if node.Ignore {
		buf.WriteString("ignore ")
	}
	if node.Targets != nil {
		node.Targets.formatFast(buf)
		buf.WriteByte(' ')
	}
	buf.WriteString("from ")
	node.TableExprs.formatFast(buf)
	node.Partitions.formatFast(buf)
	node.Where.formatFast(buf)
	node.OrderBy.formatFast(buf)
	node.Limit.formatFast(buf)
}

// formatFast formats the node.
func (node *Set) formatFast(buf *TrackedBuffer) {
	buf.WriteString("set ")
	node.Comments.formatFast(buf)
	node.Exprs.formatFast(buf)
}

// formatFast formats the node.
func (node *SetTransaction) formatFast(buf *TrackedBuffer) {
	if node.Scope == ImplicitScope {
		buf.WriteString("set ")
		node.Comments.formatFast(buf)
		buf.WriteString("transaction ")
	} else {
		buf.WriteString("set ")
		node.Comments.formatFast(buf)
		buf.WriteString(node.Scope.ToString())
		buf.WriteString(" transaction ")
	}

	for i, char := range node.Characteristics {
		if i > 0 {
			buf.WriteString(", ")
		}
		char.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *DropDatabase) formatFast(buf *TrackedBuffer) {
	exists := ""
	if node.IfExists {
		exists = "if exists "
	}
	buf.WriteString(DropStr)
	buf.WriteString(" database ")
	node.Comments.formatFast(buf)
	buf.WriteString(exists)
	buf.WriteString(node.DBName)
}

// formatFast formats the node.
func (node *Flush) formatFast(buf *TrackedBuffer) {
	buf.WriteString(FlushStr)
	if node.IsLocal {
		buf.WriteString(" local")
	}
	if len(node.FlushOptions) != 0 {
		for i, option := range node.FlushOptions {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(option)
		}
	} else {
		buf.WriteString(" tables")
		if len(node.TableNames) != 0 {
			buf.WriteByte(' ')
			node.TableNames.formatFast(buf)
		}
		if node.ForExport {
			buf.WriteString(" for export")
		}
		if node.WithLock {
			buf.WriteString(" with read lock")
		}
	}
}

// formatFast formats the node.
func (node *AlterVschema) formatFast(buf *TrackedBuffer) {
	switch node.Action {
	case CreateVindexDDLAction:
		buf.WriteString("alter vschema create vindex ")
		node.Table.formatFast(buf)
		buf.WriteByte(' ')
		node.VindexSpec.formatFast(buf)
	case DropVindexDDLAction:
		buf.WriteString("alter vschema drop vindex ")
		node.Table.formatFast(buf)
	case AddVschemaTableDDLAction:
		buf.WriteString("alter vschema add table ")
		node.Table.formatFast(buf)
	case DropVschemaTableDDLAction:
		buf.WriteString("alter vschema drop table ")
		node.Table.formatFast(buf)
	case AddColVindexDDLAction:
		buf.WriteString("alter vschema on ")
		node.Table.formatFast(buf)
		buf.WriteString(" add vindex ")
		node.VindexSpec.Name.formatFast(buf)
		buf.WriteString(" (")
		for i, col := range node.VindexCols {
			if i != 0 {
				buf.WriteString(", ")
				col.formatFast(buf)
			} else {
				col.formatFast(buf)
			}
		}
		buf.WriteByte(')')
		if node.VindexSpec.Type.String() != "" {
			buf.WriteByte(' ')
			node.VindexSpec.formatFast(buf)
		}
	case DropColVindexDDLAction:
		buf.WriteString("alter vschema on ")
		node.Table.formatFast(buf)
		buf.WriteString(" drop vindex ")
		node.VindexSpec.Name.formatFast(buf)
	case AddSequenceDDLAction:
		buf.WriteString("alter vschema add sequence ")
		node.Table.formatFast(buf)
	case AddAutoIncDDLAction:
		buf.WriteString("alter vschema on ")
		node.Table.formatFast(buf)
		buf.WriteString(" add auto_increment ")
		node.AutoIncSpec.formatFast(buf)
	default:
		buf.WriteString(node.Action.ToString())
		buf.WriteString(" table ")
		node.Table.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *RevertMigration) formatFast(buf *TrackedBuffer) {
	buf.WriteString("revert vitess_migration '")
	buf.WriteString(node.UUID)
	buf.WriteByte('\'')
}

// formatFast formats the node.
func (node *OptLike) formatFast(buf *TrackedBuffer) {
	buf.WriteString("like ")
	node.LikeTable.formatFast(buf)
}

// formatFast formats the node.
func (node *PartitionSpec) formatFast(buf *TrackedBuffer) {
	switch node.Action {
	case ReorganizeAction:
		buf.WriteString(ReorganizeStr)
		buf.WriteByte(' ')
		for i, n := range node.Names {
			if i != 0 {
				buf.WriteString(", ")
			}
			n.formatFast(buf)
		}
		buf.WriteString(" into (")
		for i, pd := range node.Definitions {
			if i != 0 {
				buf.WriteString(", ")
			}
			pd.formatFast(buf)
		}
		buf.WriteByte(')')
	case AddAction:
		buf.WriteString(AddStr)
		buf.WriteString(" (")
		node.Definitions[0].formatFast(buf)
		buf.WriteByte(')')
	case DropAction:
		buf.WriteString(DropPartitionStr)
		buf.WriteByte(' ')
		for i, n := range node.Names {
			if i != 0 {
				buf.WriteString(", ")
			}
			n.formatFast(buf)
		}
	case DiscardAction:
		buf.WriteString(DiscardStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
		buf.WriteString(" tablespace")
	case ImportAction:
		buf.WriteString(ImportStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
		buf.WriteString(" tablespace")
	case TruncateAction:
		buf.WriteString(TruncatePartitionStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
	case CoalesceAction:
		buf.WriteString(CoalesceStr)
		buf.WriteByte(' ')
		node.Number.formatFast(buf)
	case ExchangeAction:
		buf.WriteString(ExchangeStr)
		buf.WriteByte(' ')
		node.Names[0].formatFast(buf)
		buf.WriteString(" with table ")
		node.TableName.formatFast(buf)
		if node.WithoutValidation {
			buf.WriteString(" without validation")
		}
	case AnalyzeAction:
		buf.WriteString(AnalyzePartitionStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
	case CheckAction:
		buf.WriteString(CheckStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
	case OptimizeAction:
		buf.WriteString(OptimizeStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
	case RebuildAction:
		buf.WriteString(RebuildStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
	case RepairAction:
		buf.WriteString(RepairStr)
		buf.WriteByte(' ')
		if node.IsAll {
			buf.WriteString("all")
		} else {
			prefix := ""
			for _, n := range node.Names {
				buf.WriteString(prefix)
				n.formatFast(buf)
				prefix = ", "
			}
		}
	case RemoveAction:
		buf.WriteString(RemoveStr)
	case UpgradeAction:
		buf.WriteString(UpgradeStr)
	default:
		panic("unimplemented")
	}
}

// formatFast formats the node
func (node *PartitionDefinition) formatFast(buf *TrackedBuffer) {
	if !node.Maxvalue {
		buf.WriteString("partition ")
		node.Name.formatFast(buf)
		buf.WriteString(" values less than (")
		node.Limit.formatFast(buf)
		buf.WriteByte(')')
	} else {
		buf.WriteString("partition ")
		node.Name.formatFast(buf)
		buf.WriteString(" values less than (maxvalue)")
	}
}

// formatFast formats the node.
func (ts *TableSpec) formatFast(buf *TrackedBuffer) {
	buf.WriteString("(\n")
	for i, col := range ts.Columns {
		if i == 0 {
			buf.WriteByte('\t')
			col.formatFast(buf)
		} else {
			buf.WriteString(",\n\t")
			col.formatFast(buf)
		}
	}
	for _, idx := range ts.Indexes {
		buf.WriteString(",\n\t")
		idx.formatFast(buf)
	}
	for _, c := range ts.Constraints {
		buf.WriteString(",\n\t")
		c.formatFast(buf)
	}

	buf.WriteString("\n)")
	for i, opt := range ts.Options {
		if i != 0 {
			buf.WriteString(",\n ")
		}
		buf.WriteByte(' ')
		buf.WriteString(opt.Name)
		if opt.String != "" {
			buf.WriteByte(' ')
			buf.WriteString(opt.String)
		} else if opt.Value != nil {
			buf.WriteByte(' ')
			opt.Value.formatFast(buf)
		} else {
			buf.WriteString(" (")
			opt.Tables.formatFast(buf)
			buf.WriteByte(')')
		}
	}
}

// formatFast formats the node.
func (col *ColumnDefinition) formatFast(buf *TrackedBuffer) {
	col.Name.formatFast(buf)
	buf.WriteByte(' ')
	(&col.Type).formatFast(buf)
}

// formatFast returns a canonical string representation of the type and all relevant options
func (ct *ColumnType) formatFast(buf *TrackedBuffer) {
	buf.WriteString(ct.Type)

	if ct.Length != nil && ct.Scale != nil {
		buf.WriteByte('(')
		ct.Length.formatFast(buf)
		buf.WriteByte(',')
		ct.Scale.formatFast(buf)
		buf.WriteByte(')')

	} else if ct.Length != nil {
		buf.WriteByte('(')
		ct.Length.formatFast(buf)
		buf.WriteByte(')')
	}

	if ct.EnumValues != nil {
		buf.WriteByte('(')
		buf.WriteString(strings.Join(ct.EnumValues, ", "))
		buf.WriteByte(')')
	}

	if ct.Unsigned {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[UNSIGNED])
	}
	if ct.Zerofill {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[ZEROFILL])
	}
	if ct.Charset != "" {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[CHARACTER])
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[SET])
		buf.WriteByte(' ')
		buf.WriteString(ct.Charset)
	}
	if ct.Collate != "" {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[COLLATE])
		buf.WriteByte(' ')
		buf.WriteString(ct.Collate)
	}
	if ct.Options.NotNull {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[NOT])
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[NULL])
	}
	if ct.Options.Default != nil {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[DEFAULT])
		buf.WriteByte(' ')
		ct.Options.Default.formatFast(buf)
	}
	if ct.Options.OnUpdate != nil {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[ON])
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[UPDATE])
		buf.WriteByte(' ')
		ct.Options.OnUpdate.formatFast(buf)
	}
	if ct.Options.Autoincrement {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[AUTO_INCREMENT])
	}
	if ct.Options.Comment != nil {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[COMMENT_KEYWORD])
		buf.WriteByte(' ')
		ct.Options.Comment.formatFast(buf)
	}
	if ct.Options.KeyOpt == colKeyPrimary {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[PRIMARY])
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[KEY])
	}
	if ct.Options.KeyOpt == colKeyUnique {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[UNIQUE])
	}
	if ct.Options.KeyOpt == colKeyUniqueKey {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[UNIQUE])
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[KEY])
	}
	if ct.Options.KeyOpt == colKeySpatialKey {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[SPATIAL])
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[KEY])
	}
	if ct.Options.KeyOpt == colKeyFulltextKey {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[FULLTEXT])
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[KEY])
	}
	if ct.Options.KeyOpt == colKey {
		buf.WriteByte(' ')
		buf.WriteString(keywordStrings[KEY])
	}
}

// formatFast formats the node.
func (idx *IndexDefinition) formatFast(buf *TrackedBuffer) {
	idx.Info.formatFast(buf)
	buf.WriteString(" (")
	for i, col := range idx.Columns {
		if i != 0 {
			buf.WriteString(", ")
			col.Column.formatFast(buf)
		} else {
			col.Column.formatFast(buf)
		}
		if col.Length != nil {
			buf.WriteByte('(')
			col.Length.formatFast(buf)
			buf.WriteByte(')')
		}
		if col.Direction == DescOrder {
			buf.WriteString(" desc")
		}
	}
	buf.WriteByte(')')

	for _, opt := range idx.Options {
		buf.WriteByte(' ')
		buf.WriteString(opt.Name)
		if opt.String != "" {
			buf.WriteByte(' ')
			buf.WriteString(opt.String)
		} else {
			buf.WriteByte(' ')
			opt.Value.formatFast(buf)
		}
	}
}

// formatFast formats the node.
func (ii *IndexInfo) formatFast(buf *TrackedBuffer) {
	if !ii.ConstraintName.IsEmpty() {
		buf.WriteString("constraint ")
		ii.ConstraintName.formatFast(buf)
		buf.WriteByte(' ')
	}
	if ii.Primary {
		buf.WriteString(ii.Type)
	} else {
		buf.WriteString(ii.Type)
		if !ii.Name.IsEmpty() {
			buf.WriteByte(' ')
			ii.Name.formatFast(buf)
		}
	}
}

// formatFast formats the node.
func (node *AutoIncSpec) formatFast(buf *TrackedBuffer) {
	node.Column.formatFast(buf)
	buf.WriteByte(' ')
	buf.WriteString("using ")
	node.Sequence.formatFast(buf)
}

// formatFast formats the node. The "CREATE VINDEX" preamble was formatted in
// the containing DDL node Format, so this just prints the type, any
// parameters, and optionally the owner
func (node *VindexSpec) formatFast(buf *TrackedBuffer) {
	buf.WriteString("using ")
	node.Type.formatFast(buf)

	numParams := len(node.Params)
	if numParams != 0 {
		buf.WriteString(" with ")
		for i, p := range node.Params {
			if i != 0 {
				buf.WriteString(", ")
			}
			p.formatFast(buf)
		}
	}
}

// formatFast formats the node.
func (node VindexParam) formatFast(buf *TrackedBuffer) {
	buf.WriteString(node.Key.String())
	buf.WriteByte('=')
	buf.WriteString(node.Val)
}

// formatFast formats the node.
func (c *ConstraintDefinition) formatFast(buf *TrackedBuffer) {
	if c.Name != "" {
		buf.WriteString("constraint ")
		buf.WriteString(c.Name)
		buf.WriteByte(' ')
	}
	c.Details.Format(buf)
}

// formatFast formats the node.
func (a ReferenceAction) formatFast(buf *TrackedBuffer) {
	switch a {
	case Restrict:
		buf.WriteString("restrict")
	case Cascade:
		buf.WriteString("cascade")
	case NoAction:
		buf.WriteString("no action")
	case SetNull:
		buf.WriteString("set null")
	case SetDefault:
		buf.WriteString("set default")
	}
}

// formatFast formats the node.
func (f *ForeignKeyDefinition) formatFast(buf *TrackedBuffer) {
	buf.WriteString("foreign key ")
	f.Source.formatFast(buf)
	buf.WriteString(" references ")
	f.ReferencedTable.formatFast(buf)
	buf.WriteByte(' ')
	f.ReferencedColumns.formatFast(buf)
	if f.OnDelete != DefaultAction {
		buf.WriteString(" on delete ")
		f.OnDelete.formatFast(buf)
	}
	if f.OnUpdate != DefaultAction {
		buf.WriteString(" on update ")
		f.OnUpdate.formatFast(buf)
	}
}

// formatFast formats the node.
func (c *CheckConstraintDefinition) formatFast(buf *TrackedBuffer) {
	buf.WriteString("check (")
	c.Expr.formatFast(buf)
	buf.WriteByte(')')
	if !c.Enforced {
		buf.WriteString(" not enforced")
	}
}

// formatFast formats the node.
func (node *Show) formatFast(buf *TrackedBuffer) {
	node.Internal.formatFast(buf)
}

// formatFast formats the node.
func (node *ShowLegacy) formatFast(buf *TrackedBuffer) {
	nodeType := strings.ToLower(node.Type)
	if (nodeType == "tables" || nodeType == "columns" || nodeType == "fields" || nodeType == "index" || nodeType == "keys" || nodeType == "indexes" ||
		nodeType == "databases" || nodeType == "schemas" || nodeType == "keyspaces" || nodeType == "vitess_keyspaces" || nodeType == "vitess_shards" || nodeType == "vitess_tablets") && node.ShowTablesOpt != nil {
		opt := node.ShowTablesOpt
		if node.Extended != "" {
			buf.WriteString("show ")
			buf.WriteString(node.Extended)
			buf.WriteString(nodeType)
		} else {
			buf.WriteString("show ")
			buf.WriteString(opt.Full)
			buf.WriteString(nodeType)
		}
		if (nodeType == "columns" || nodeType == "fields") && node.HasOnTable() {
			buf.WriteString(" from ")
			node.OnTable.formatFast(buf)
		}
		if (nodeType == "index" || nodeType == "keys" || nodeType == "indexes") && node.HasOnTable() {
			buf.WriteString(" from ")
			node.OnTable.formatFast(buf)
		}
		if opt.DbName != "" {
			buf.WriteString(" from ")
			buf.WriteString(opt.DbName)
		}
		opt.Filter.formatFast(buf)
		return
	}
	if node.Scope == ImplicitScope {
		buf.WriteString("show ")
		buf.WriteString(nodeType)
	} else {
		buf.WriteString("show ")
		buf.WriteString(node.Scope.ToString())
		buf.WriteByte(' ')
		buf.WriteString(nodeType)
	}
	if node.HasOnTable() {
		buf.WriteString(" on ")
		node.OnTable.formatFast(buf)
	}
	if nodeType == "collation" && node.ShowCollationFilterOpt != nil {
		buf.WriteString(" where ")
		node.ShowCollationFilterOpt.formatFast(buf)
	}
	if nodeType == "charset" && node.ShowTablesOpt != nil {
		node.ShowTablesOpt.Filter.formatFast(buf)
	}
	if node.HasTable() {
		buf.WriteByte(' ')
		node.Table.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *ShowFilter) formatFast(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	if node.Like != "" {
		buf.WriteString(" like ")
		buf.WriteString(encodeSQLString(node.Like))
	} else {
		buf.WriteString(" where ")
		node.Filter.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *Use) formatFast(buf *TrackedBuffer) {
	if node.DBName.v != "" {
		buf.WriteString("use ")
		node.DBName.formatFast(buf)
	} else {
		buf.WriteString("use")
	}
}

// formatFast formats the node.
func (node *Commit) formatFast(buf *TrackedBuffer) {
	buf.WriteString("commit")
}

// formatFast formats the node.
func (node *Begin) formatFast(buf *TrackedBuffer) {
	buf.WriteString("begin")
}

// formatFast formats the node.
func (node *Rollback) formatFast(buf *TrackedBuffer) {
	buf.WriteString("rollback")
}

// formatFast formats the node.
func (node *SRollback) formatFast(buf *TrackedBuffer) {
	buf.WriteString("rollback to ")
	node.Name.formatFast(buf)
}

// formatFast formats the node.
func (node *Savepoint) formatFast(buf *TrackedBuffer) {
	buf.WriteString("savepoint ")
	node.Name.formatFast(buf)
}

// formatFast formats the node.
func (node *Release) formatFast(buf *TrackedBuffer) {
	buf.WriteString("release savepoint ")
	node.Name.formatFast(buf)
}

// formatFast formats the node.
func (node *ExplainStmt) formatFast(buf *TrackedBuffer) {
	format := ""
	switch node.Type {
	case EmptyType:
	case AnalyzeType:
		format = AnalyzeStr + " "
	default:
		format = "format = " + node.Type.ToString() + " "
	}
	buf.WriteString("explain ")
	buf.WriteString(format)
	node.Statement.formatFast(buf)
}

// formatFast formats the node.
func (node *ExplainTab) formatFast(buf *TrackedBuffer) {
	buf.WriteString("explain ")
	node.Table.formatFast(buf)
	if node.Wild != "" {
		buf.WriteByte(' ')
		buf.WriteString(node.Wild)
	}
}

// formatFast formats the node.
func (node *CallProc) formatFast(buf *TrackedBuffer) {
	buf.WriteString("call ")
	node.Name.formatFast(buf)
	buf.WriteByte('(')
	node.Params.formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *OtherRead) formatFast(buf *TrackedBuffer) {
	buf.WriteString("otherread")
}

// formatFast formats the node.
func (node *OtherAdmin) formatFast(buf *TrackedBuffer) {
	buf.WriteString("otheradmin")
}

// formatFast formats the node.
func (node Comments) formatFast(buf *TrackedBuffer) {
	for _, c := range node {
		buf.WriteString(c)
		buf.WriteByte(' ')
	}
}

// formatFast formats the node.
func (node SelectExprs) formatFast(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node *StarExpr) formatFast(buf *TrackedBuffer) {
	if !node.TableName.IsEmpty() {
		node.TableName.formatFast(buf)
		buf.WriteByte('.')
	}
	buf.WriteByte('*')
}

// formatFast formats the node.
func (node *AliasedExpr) formatFast(buf *TrackedBuffer) {
	node.Expr.formatFast(buf)
	if !node.As.IsEmpty() {
		buf.WriteString(" as ")
		node.As.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *Nextval) formatFast(buf *TrackedBuffer) {
	buf.WriteString("next ")
	node.Expr.formatFast(buf)
	buf.WriteString(" values")
}

// formatFast formats the node.
func (node Columns) formatFast(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	prefix := "("
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
	buf.WriteString(")")
}

// formatFast formats the node
func (node Partitions) formatFast(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	prefix := " partition ("
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
	buf.WriteString(")")
}

// formatFast formats the node.
func (node TableExprs) formatFast(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node *AliasedTableExpr) formatFast(buf *TrackedBuffer) {
	node.Expr.formatFast(buf)
	node.Partitions.formatFast(buf)
	if !node.As.IsEmpty() {
		buf.WriteString(" as ")
		node.As.formatFast(buf)
	}
	if node.Hints != nil {
		// Hint node provides the space padding.
		node.Hints.formatFast(buf)
	}
}

// formatFast formats the node.
func (node TableNames) formatFast(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node TableName) formatFast(buf *TrackedBuffer) {
	if node.IsEmpty() {
		return
	}
	if !node.Qualifier.IsEmpty() {
		node.Qualifier.formatFast(buf)
		buf.WriteByte('.')
	}
	node.Name.formatFast(buf)
}

// formatFast formats the node.
func (node *ParenTableExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteByte('(')
	node.Exprs.formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node JoinCondition) formatFast(buf *TrackedBuffer) {
	if node.On != nil {
		buf.WriteString(" on ")
		node.On.formatFast(buf)
	}
	if node.Using != nil {
		buf.WriteString(" using ")
		node.Using.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *JoinTableExpr) formatFast(buf *TrackedBuffer) {
	node.LeftExpr.formatFast(buf)
	buf.WriteByte(' ')
	buf.WriteString(node.Join.ToString())
	buf.WriteByte(' ')
	node.RightExpr.formatFast(buf)
	node.Condition.formatFast(buf)
}

// formatFast formats the node.
func (node *IndexHints) formatFast(buf *TrackedBuffer) {
	buf.WriteByte(' ')
	buf.WriteString(node.Type.ToString())
	buf.WriteString("index ")
	if len(node.Indexes) == 0 {
		buf.WriteString("()")
	} else {
		prefix := "("
		for _, n := range node.Indexes {
			buf.WriteString(prefix)
			n.formatFast(buf)
			prefix = ", "
		}
		buf.WriteByte(')')
	}
}

// formatFast formats the node.
func (node *Where) formatFast(buf *TrackedBuffer) {
	if node == nil || node.Expr == nil {
		return
	}
	buf.WriteByte(' ')
	buf.WriteString(node.Type.ToString())
	buf.WriteByte(' ')
	node.Expr.formatFast(buf)
}

// formatFast formats the node.
func (node Exprs) formatFast(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node *AndExpr) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Left, true)
	buf.WriteString(" and ")
	buf.printExpr(node, node.Right, false)
}

// formatFast formats the node.
func (node *OrExpr) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Left, true)
	buf.WriteString(" or ")
	buf.printExpr(node, node.Right, false)
}

// formatFast formats the node.
func (node *XorExpr) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Left, true)
	buf.WriteString(" xor ")
	buf.printExpr(node, node.Right, false)
}

// formatFast formats the node.
func (node *NotExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("not ")
	buf.printExpr(node, node.Expr, true)
}

// formatFast formats the node.
func (node *ComparisonExpr) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Left, true)
	buf.WriteByte(' ')
	buf.WriteString(node.Operator.ToString())
	buf.WriteByte(' ')
	buf.printExpr(node, node.Right, false)
	if node.Escape != nil {
		buf.WriteString(" escape ")
		buf.printExpr(node, node.Escape, true)
	}
}

// formatFast formats the node.
func (node *RangeCond) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Left, true)
	buf.WriteByte(' ')
	buf.WriteString(node.Operator.ToString())
	buf.WriteByte(' ')
	buf.printExpr(node, node.From, true)
	buf.WriteString(" and ")
	buf.printExpr(node, node.To, false)
}

// formatFast formats the node.
func (node *IsExpr) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Expr, true)
	buf.WriteByte(' ')
	buf.WriteString(node.Operator.ToString())
}

// formatFast formats the node.
func (node *ExistsExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("exists ")
	buf.printExpr(node, node.Subquery, true)
}

// formatFast formats the node.
func (node *Literal) formatFast(buf *TrackedBuffer) {
	switch node.Type {
	case StrVal:
		sqltypes.MakeTrusted(sqltypes.VarBinary, node.Bytes()).EncodeSQL(buf)
	case IntVal, FloatVal, HexNum:
		buf.WriteString(node.Val)
	case HexVal:
		buf.WriteString("X'")
		buf.WriteString(node.Val)
		buf.WriteByte('\'')
	case BitVal:
		buf.WriteString("B'")
		buf.WriteString(node.Val)
		buf.WriteByte('\'')
	default:
		panic("unexpected")
	}
}

// formatFast formats the node.
func (node Argument) formatFast(buf *TrackedBuffer) {
	buf.WriteArg(string(node))
}

// formatFast formats the node.
func (node *NullVal) formatFast(buf *TrackedBuffer) {
	buf.WriteString("null")
}

// formatFast formats the node.
func (node BoolVal) formatFast(buf *TrackedBuffer) {
	if node {
		buf.WriteString("true")
	} else {
		buf.WriteString("false")
	}
}

// formatFast formats the node.
func (node *ColName) formatFast(buf *TrackedBuffer) {
	if !node.Qualifier.IsEmpty() {
		node.Qualifier.formatFast(buf)
		buf.WriteByte('.')
	}
	node.Name.formatFast(buf)
}

// formatFast formats the node.
func (node ValTuple) formatFast(buf *TrackedBuffer) {
	buf.WriteByte('(')
	Exprs(node).formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *Subquery) formatFast(buf *TrackedBuffer) {
	buf.WriteByte('(')
	node.Select.formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *DerivedTable) formatFast(buf *TrackedBuffer) {
	buf.WriteByte('(')
	node.Select.formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node ListArg) formatFast(buf *TrackedBuffer) {
	buf.WriteArg(string(node))
}

// formatFast formats the node.
func (node *BinaryExpr) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Left, true)
	buf.WriteByte(' ')
	buf.WriteString(node.Operator.ToString())
	buf.WriteByte(' ')
	buf.printExpr(node, node.Right, false)
}

// formatFast formats the node.
func (node *UnaryExpr) formatFast(buf *TrackedBuffer) {
	if _, unary := node.Expr.(*UnaryExpr); unary {
		// They have same precedence so parenthesis is not required.
		buf.WriteString(node.Operator.ToString())
		buf.WriteByte(' ')
		buf.printExpr(node, node.Expr, true)
		return
	}
	buf.WriteString(node.Operator.ToString())
	buf.printExpr(node, node.Expr, true)
}

// formatFast formats the node.
func (node *IntervalExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("interval ")
	buf.printExpr(node, node.Expr, true)
	buf.WriteByte(' ')
	buf.WriteString(node.Unit)
}

// formatFast formats the node.
func (node *TimestampFuncExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString(node.Name)
	buf.WriteByte('(')
	buf.WriteString(node.Unit)
	buf.WriteString(", ")
	buf.printExpr(node, node.Expr1, true)
	buf.WriteString(", ")
	buf.printExpr(node, node.Expr2, true)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *CurTimeFuncExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString(node.Name.String())
	buf.WriteByte('(')
	buf.printExpr(node, node.Fsp, true)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *CollateExpr) formatFast(buf *TrackedBuffer) {
	buf.printExpr(node, node.Expr, true)
	buf.WriteString(" collate ")
	buf.WriteString(node.Charset)
}

// formatFast formats the node.
func (node *FuncExpr) formatFast(buf *TrackedBuffer) {
	var distinct string
	if node.Distinct {
		distinct = "distinct "
	}
	if !node.Qualifier.IsEmpty() {
		node.Qualifier.formatFast(buf)
		buf.WriteByte('.')
	}
	// Function names should not be back-quoted even
	// if they match a reserved word, only if they contain illegal characters
	funcName := node.Name.String()

	if containEscapableChars(funcName, NoAt) {
		writeEscapedString(buf, funcName)
	} else {
		buf.WriteString(funcName)
	}
	buf.WriteByte('(')
	buf.WriteString(distinct)
	node.Exprs.formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node
func (node *GroupConcatExpr) formatFast(buf *TrackedBuffer) {
	if node.Distinct {
		buf.WriteString("group_concat(")
		buf.WriteString(DistinctStr)
		node.Exprs.formatFast(buf)
		node.OrderBy.formatFast(buf)
		buf.WriteString(node.Separator)
		node.Limit.formatFast(buf)
		buf.WriteByte(')')
	} else {
		buf.WriteString("group_concat(")
		node.Exprs.formatFast(buf)
		node.OrderBy.formatFast(buf)
		buf.WriteString(node.Separator)
		node.Limit.formatFast(buf)
		buf.WriteByte(')')
	}
}

// formatFast formats the node.
func (node *ValuesFuncExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("values(")
	buf.printExpr(node, node.Name, true)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *SubstrExpr) formatFast(buf *TrackedBuffer) {
	var val SQLNode
	if node.Name != nil {
		val = node.Name
	} else {
		val = node.StrVal
	}

	if node.To == nil {
		buf.WriteString("substr(")
		val.formatFast(buf)
		buf.WriteString(", ")
		buf.printExpr(node, node.From, true)
		buf.WriteByte(')')
	} else {
		buf.WriteString("substr(")
		val.formatFast(buf)
		buf.WriteString(", ")
		buf.printExpr(node, node.From, true)
		buf.WriteString(", ")
		buf.printExpr(node, node.To, true)
		buf.WriteByte(')')
	}
}

// formatFast formats the node.
func (node *ConvertExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("convert(")
	buf.printExpr(node, node.Expr, true)
	buf.WriteString(", ")
	node.Type.formatFast(buf)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *ConvertUsingExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("convert(")
	buf.printExpr(node, node.Expr, true)
	buf.WriteString(" using ")
	buf.WriteString(node.Type)
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *ConvertType) formatFast(buf *TrackedBuffer) {
	buf.WriteString(node.Type)
	if node.Length != nil {
		buf.WriteByte('(')
		node.Length.formatFast(buf)
		if node.Scale != nil {
			buf.WriteString(", ")
			node.Scale.formatFast(buf)
		}
		buf.WriteByte(')')
	}
	if node.Charset != "" {
		buf.WriteString(node.Operator.ToString())
		buf.WriteByte(' ')
		buf.WriteString(node.Charset)
	}
}

// formatFast formats the node
func (node *MatchExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("match(")
	node.Columns.formatFast(buf)
	buf.WriteString(") against (")
	buf.printExpr(node, node.Expr, true)
	buf.WriteString(node.Option.ToString())
	buf.WriteByte(')')
}

// formatFast formats the node.
func (node *CaseExpr) formatFast(buf *TrackedBuffer) {
	buf.WriteString("case ")
	if node.Expr != nil {
		buf.printExpr(node, node.Expr, true)
		buf.WriteByte(' ')
	}
	for _, when := range node.Whens {
		when.formatFast(buf)
		buf.WriteByte(' ')
	}
	if node.Else != nil {
		buf.WriteString("else ")
		buf.printExpr(node, node.Else, true)
		buf.WriteByte(' ')
	}
	buf.WriteString("end")
}

// formatFast formats the node.
func (node *Default) formatFast(buf *TrackedBuffer) {
	buf.WriteString("default")
	if node.ColName != "" {
		buf.WriteString("(")
		formatID(buf, node.ColName, NoAt)
		buf.WriteString(")")
	}
}

// formatFast formats the node.
func (node *When) formatFast(buf *TrackedBuffer) {
	buf.WriteString("when ")
	node.Cond.formatFast(buf)
	buf.WriteString(" then ")
	node.Val.formatFast(buf)
}

// formatFast formats the node.
func (node GroupBy) formatFast(buf *TrackedBuffer) {
	prefix := " group by "
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node OrderBy) formatFast(buf *TrackedBuffer) {
	prefix := " order by "
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node *Order) formatFast(buf *TrackedBuffer) {
	if node, ok := node.Expr.(*NullVal); ok {
		buf.printExpr(node, node, true)
		return
	}
	if node, ok := node.Expr.(*FuncExpr); ok {
		if node.Name.Lowered() == "rand" {
			buf.printExpr(node, node, true)
			return
		}
	}

	node.Expr.formatFast(buf)
	buf.WriteByte(' ')
	buf.WriteString(node.Direction.ToString())
}

// formatFast formats the node.
func (node *Limit) formatFast(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	buf.WriteString(" limit ")
	if node.Offset != nil {
		node.Offset.formatFast(buf)
		buf.WriteString(", ")
	}
	node.Rowcount.formatFast(buf)
}

// formatFast formats the node.
func (node Values) formatFast(buf *TrackedBuffer) {
	prefix := "values "
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node UpdateExprs) formatFast(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node *UpdateExpr) formatFast(buf *TrackedBuffer) {
	node.Name.formatFast(buf)
	buf.WriteString(" = ")
	node.Expr.formatFast(buf)
}

// formatFast formats the node.
func (node SetExprs) formatFast(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node.
func (node *SetExpr) formatFast(buf *TrackedBuffer) {
	if node.Scope != ImplicitScope {
		buf.WriteString(node.Scope.ToString())
		buf.WriteString(" ")
	}
	// We don't have to backtick set variable names.
	switch {
	case node.Name.EqualString("charset") || node.Name.EqualString("names"):
		buf.WriteString(node.Name.String())
		buf.WriteByte(' ')
		node.Expr.formatFast(buf)
	case node.Name.EqualString(TransactionStr):
		literal := node.Expr.(*Literal)
		buf.WriteString(node.Name.String())
		buf.WriteByte(' ')
		buf.WriteString(strings.ToLower(string(literal.Val)))
	default:
		node.Name.formatFast(buf)
		buf.WriteString(" = ")
		node.Expr.formatFast(buf)
	}
}

// formatFast formats the node.
func (node OnDup) formatFast(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	buf.WriteString(" on duplicate key update ")
	UpdateExprs(node).formatFast(buf)
}

// formatFast formats the node.
func (node ColIdent) formatFast(buf *TrackedBuffer) {
	for i := NoAt; i < node.at; i++ {
		buf.WriteByte('@')
	}
	formatID(buf, node.val, node.at)
}

// formatFast formats the node.
func (node TableIdent) formatFast(buf *TrackedBuffer) {
	formatID(buf, node.v, NoAt)
}

// formatFast formats the node.
func (node IsolationLevel) formatFast(buf *TrackedBuffer) {
	buf.WriteString("isolation level ")
	switch node {
	case ReadUncommitted:
		buf.WriteString(ReadUncommittedStr)
	case ReadCommitted:
		buf.WriteString(ReadCommittedStr)
	case RepeatableRead:
		buf.WriteString(RepeatableReadStr)
	case Serializable:
		buf.WriteString(SerializableStr)
	default:
		buf.WriteString("Unknown Isolation level value")
	}
}

// formatFast formats the node.
func (node AccessMode) formatFast(buf *TrackedBuffer) {
	if node == ReadOnly {
		buf.WriteString(TxReadOnly)
	} else {
		buf.WriteString(TxReadWrite)
	}
}

// formatFast formats the node.
func (node *Load) formatFast(buf *TrackedBuffer) {
	buf.WriteString("AST node missing for Load type")
}

// formatFast formats the node.
func (node *ShowBasic) formatFast(buf *TrackedBuffer) {
	buf.WriteString("show")
	if node.Full {
		buf.WriteString(" full")
	}
	buf.WriteString(node.Command.ToString())
	if !node.Tbl.IsEmpty() {
		buf.WriteString(" from ")
		node.Tbl.formatFast(buf)
	}
	if node.DbName != "" {
		buf.WriteString(" from ")
		buf.WriteString(node.DbName)
	}
	node.Filter.formatFast(buf)
}

// formatFast formats the node.
func (node *ShowCreate) formatFast(buf *TrackedBuffer) {
	buf.WriteString("show")
	buf.WriteString(node.Command.ToString())
	buf.WriteByte(' ')
	node.Op.formatFast(buf)
}

// formatFast formats the node.
func (node *SelectInto) formatFast(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	buf.WriteString(node.Type.ToString())
	buf.WriteString(node.FileName)
	if node.Charset != "" {
		buf.WriteString(" character set ")
		buf.WriteString(node.Charset)
	}
	buf.WriteString(node.FormatOption)
	buf.WriteString(node.ExportOption)
	buf.WriteString(node.Manifest)
	buf.WriteString(node.Overwrite)
}

// formatFast formats the node.
func (node *CreateDatabase) formatFast(buf *TrackedBuffer) {
	buf.WriteString("create database ")
	node.Comments.formatFast(buf)
	if node.IfNotExists {
		buf.WriteString("if not exists ")
	}
	buf.WriteString(node.DBName)
	if node.CreateOptions != nil {
		for _, createOption := range node.CreateOptions {
			if createOption.IsDefault {
				buf.WriteString(" default")
			}
			buf.WriteString(createOption.Type.ToString())
			buf.WriteString(" " + createOption.Value)
		}
	}
}

// formatFast formats the node.
func (node *AlterDatabase) formatFast(buf *TrackedBuffer) {
	buf.WriteString("alter database")
	if node.DBName != "" {
		buf.WriteByte(' ')
		buf.WriteString(node.DBName)
	}
	if node.UpdateDataDirectory {
		buf.WriteString(" upgrade data directory name")
	}
	if node.AlterOptions != nil {
		for _, createOption := range node.AlterOptions {
			if createOption.IsDefault {
				buf.WriteString(" default")
			}
			buf.WriteString(createOption.Type.ToString())
			buf.WriteString(" " + createOption.Value)
		}
	}
}

// formatFast formats the node.
func (node *CreateTable) formatFast(buf *TrackedBuffer) {
	buf.WriteString("create ")
	if node.Temp {
		buf.WriteString("temporary ")
	}
	buf.WriteString("table ")

	if node.IfNotExists {
		buf.WriteString("if not exists ")
	}
	node.Table.formatFast(buf)

	if node.OptLike != nil {
		buf.WriteByte(' ')
		node.OptLike.formatFast(buf)
	}
	if node.TableSpec != nil {
		buf.WriteByte(' ')
		node.TableSpec.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *CreateView) formatFast(buf *TrackedBuffer) {
	buf.WriteString("create")
	if node.IsReplace {
		buf.WriteString(" or replace")
	}
	if node.Algorithm != "" {
		buf.WriteString(" algorithm = ")
		buf.WriteString(node.Algorithm)
	}
	if node.Definer != "" {
		buf.WriteString(" definer = ")
		buf.WriteString(node.Definer)
	}
	if node.Security != "" {
		buf.WriteString(" sql security ")
		buf.WriteString(node.Security)
	}
	buf.WriteString(" view ")
	node.ViewName.formatFast(buf)
	node.Columns.formatFast(buf)
	buf.WriteString(" as ")
	node.Select.formatFast(buf)
	if node.CheckOption != "" {
		buf.WriteString(" with ")
		buf.WriteString(node.CheckOption)
		buf.WriteString(" check option")
	}
}

// formatFast formats the LockTables node.
func (node *LockTables) formatFast(buf *TrackedBuffer) {
	buf.WriteString("lock tables ")
	node.Tables[0].Table.formatFast(buf)
	buf.WriteByte(' ')
	buf.WriteString(node.Tables[0].Lock.ToString())
	for i := 1; i < len(node.Tables); i++ {
		buf.WriteString(", ")
		node.Tables[i].Table.formatFast(buf)
		buf.WriteByte(' ')
		buf.WriteString(node.Tables[i].Lock.ToString())
	}
}

// formatFast formats the UnlockTables node.
func (node *UnlockTables) formatFast(buf *TrackedBuffer) {
	buf.WriteString("unlock tables")
}

// formatFast formats the node.
func (node *AlterView) formatFast(buf *TrackedBuffer) {
	buf.WriteString("alter")
	if node.Algorithm != "" {
		buf.WriteString(" algorithm = ")
		buf.WriteString(node.Algorithm)
	}
	if node.Definer != "" {
		buf.WriteString(" definer = ")
		buf.WriteString(node.Definer)
	}
	if node.Security != "" {
		buf.WriteString(" sql security ")
		buf.WriteString(node.Security)
	}
	buf.WriteString(" view ")
	node.ViewName.formatFast(buf)
	node.Columns.formatFast(buf)
	buf.WriteString(" as ")
	node.Select.formatFast(buf)
	if node.CheckOption != "" {
		buf.WriteString(" with ")
		buf.WriteString(node.CheckOption)
		buf.WriteString(" check option")
	}
}

// formatFast formats the node.
func (node *DropTable) formatFast(buf *TrackedBuffer) {
	temp := ""
	if node.Temp {
		temp = " temporary"
	}
	exists := ""
	if node.IfExists {
		exists = " if exists"
	}
	buf.WriteString("drop")
	buf.WriteString(temp)
	buf.WriteString(" table")
	buf.WriteString(exists)
	buf.WriteByte(' ')
	node.FromTables.formatFast(buf)
}

// formatFast formats the node.
func (node *DropView) formatFast(buf *TrackedBuffer) {
	exists := ""
	if node.IfExists {
		exists = " if exists"
	}
	buf.WriteString("drop view")
	buf.WriteString(exists)
	buf.WriteByte(' ')
	node.FromTables.formatFast(buf)
}

// formatFast formats the AlterTable node.
func (node *AlterTable) formatFast(buf *TrackedBuffer) {
	buf.WriteString("alter table ")
	node.Table.formatFast(buf)
	prefix := ""
	for i, option := range node.AlterOptions {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(' ')
		option.formatFast(buf)
		if node.PartitionSpec != nil && node.PartitionSpec.Action != RemoveAction {
			prefix = ","
		}
	}
	if node.PartitionSpec != nil {
		buf.WriteString(prefix)
		buf.WriteByte(' ')
		node.PartitionSpec.formatFast(buf)
	}
}

// formatFast formats the node.
func (node *AddConstraintDefinition) formatFast(buf *TrackedBuffer) {
	buf.WriteString("add ")
	node.ConstraintDefinition.formatFast(buf)
}

// formatFast formats the node.
func (node *AddIndexDefinition) formatFast(buf *TrackedBuffer) {
	buf.WriteString("add ")
	node.IndexDefinition.formatFast(buf)
}

// formatFast formats the node.
func (node *AddColumns) formatFast(buf *TrackedBuffer) {

	if len(node.Columns) == 1 {
		buf.WriteString("add column ")
		node.Columns[0].formatFast(buf)
		if node.First != nil {
			buf.WriteString(" first ")
			node.First.formatFast(buf)
		}
		if node.After != nil {
			buf.WriteString(" after ")
			node.After.formatFast(buf)
		}
	} else {
		for i, col := range node.Columns {
			if i == 0 {
				buf.WriteString("add column (")
				col.formatFast(buf)
			} else {
				buf.WriteString(", ")
				col.formatFast(buf)
			}
		}
		buf.WriteString(")")
	}
}

// formatFast formats the node.
func (node AlgorithmValue) formatFast(buf *TrackedBuffer) {
	buf.WriteString("algorithm = ")
	buf.WriteString(string(node))
}

// formatFast formats the node
func (node *AlterColumn) formatFast(buf *TrackedBuffer) {
	if node.DropDefault {
		buf.WriteString("alter column ")
		node.Column.formatFast(buf)
		buf.WriteString(" drop default")
	} else {
		buf.WriteString("alter column ")
		node.Column.formatFast(buf)
		buf.WriteString(" set default")
		buf.WriteByte(' ')
		node.DefaultVal.formatFast(buf)
	}
}

// formatFast formats the node
func (node *ChangeColumn) formatFast(buf *TrackedBuffer) {
	buf.WriteString("change column ")
	node.OldColumn.formatFast(buf)
	buf.WriteByte(' ')
	node.NewColDefinition.formatFast(buf)
	if node.First != nil {
		buf.WriteString(" first ")
		node.First.formatFast(buf)
	}
	if node.After != nil {
		buf.WriteString(" after ")
		node.After.formatFast(buf)
	}
}

// formatFast formats the node
func (node *ModifyColumn) formatFast(buf *TrackedBuffer) {
	buf.WriteString("modify column ")
	node.NewColDefinition.formatFast(buf)
	if node.First != nil {
		buf.WriteString(" first ")
		node.First.formatFast(buf)
	}
	if node.After != nil {
		buf.WriteString(" after ")
		node.After.formatFast(buf)
	}
}

// formatFast formats the node
func (node *AlterCharset) formatFast(buf *TrackedBuffer) {
	buf.WriteString("convert to character set ")
	buf.WriteString(node.CharacterSet)
	if node.Collate != "" {
		buf.WriteString(" collate ")
		buf.WriteString(node.Collate)
	}
}

// formatFast formats the node
func (node *KeyState) formatFast(buf *TrackedBuffer) {
	if node.Enable {
		buf.WriteString("enable keys")
	} else {
		buf.WriteString("disable keys")
	}

}

// formatFast formats the node
func (node *TablespaceOperation) formatFast(buf *TrackedBuffer) {
	if node.Import {
		buf.WriteString("import tablespace")
	} else {
		buf.WriteString("discard tablespace")
	}
}

// formatFast formats the node
func (node *DropColumn) formatFast(buf *TrackedBuffer) {
	buf.WriteString("drop column ")
	node.Name.formatFast(buf)
}

// formatFast formats the node
func (node *DropKey) formatFast(buf *TrackedBuffer) {
	buf.WriteString("drop ")
	buf.WriteString(node.Type.ToString())
	if node.Name != "" {
		buf.WriteByte(' ')
		buf.WriteString(node.Name)
	}
}

// formatFast formats the node
func (node *Force) formatFast(buf *TrackedBuffer) {
	buf.WriteString("force")
}

// formatFast formats the node
func (node *LockOption) formatFast(buf *TrackedBuffer) {
	buf.WriteString("lock ")
	buf.WriteString(node.Type.ToString())
}

// formatFast formats the node
func (node *OrderByOption) formatFast(buf *TrackedBuffer) {
	buf.WriteString("order by ")
	prefix := ""
	for _, n := range node.Cols {
		buf.WriteString(prefix)
		n.formatFast(buf)
		prefix = ", "
	}
}

// formatFast formats the node
func (node *RenameTableName) formatFast(buf *TrackedBuffer) {
	buf.WriteString("rename ")
	node.Table.formatFast(buf)
}

// formatFast formats the node
func (node *RenameIndex) formatFast(buf *TrackedBuffer) {
	buf.WriteString("rename index ")
	buf.WriteString(node.OldName)
	buf.WriteString(" to ")
	buf.WriteString(node.NewName)
}

// formatFast formats the node
func (node *Validation) formatFast(buf *TrackedBuffer) {
	if node.With {
		buf.WriteString("with validation")
	} else {
		buf.WriteString("without validation")
	}
}

// formatFast formats the node
func (node TableOptions) formatFast(buf *TrackedBuffer) {
	for i, option := range node {
		if i != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(option.Name)
		if option.String != "" {
			buf.WriteByte(' ')
			buf.WriteString(option.String)
		} else if option.Value != nil {
			buf.WriteByte(' ')
			option.Value.formatFast(buf)
		} else {
			buf.WriteString(" (")
			option.Tables.formatFast(buf)
			buf.WriteByte(')')
		}
	}
}

// formatFast formats the node
func (node *TruncateTable) formatFast(buf *TrackedBuffer) {
	buf.WriteString("truncate table ")
	node.Table.formatFast(buf)
}

// formatFast formats the node.
func (node *RenameTable) formatFast(buf *TrackedBuffer) {
	buf.WriteString("rename table")
	prefix := " "
	for _, pair := range node.TablePairs {
		buf.WriteString(prefix)
		pair.FromTable.formatFast(buf)
		buf.WriteString(" to ")
		pair.ToTable.formatFast(buf)
		prefix = ", "
	}
}
