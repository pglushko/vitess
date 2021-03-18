package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"path"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

func main() {
	err := load(os.Args[1])
	if err != nil {
		panic(err)
	}
}

func load(packageName string) error {
	config := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(config, packageName)
	if err != nil {
		return fmt.Errorf("error loading package %s: %w", packageName, err)
	}
	for _, pkg := range pkgs {
		if pkg.Name == "sqlparser" {
			rewriter := &Rewriter{pkg: pkg}
			err := rewriter.Rewrite()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Rewriter struct {
	pkg     *packages.Package
	astExpr *types.Interface
}

func (r *Rewriter) Rewrite() error {
	scope := r.pkg.Types.Scope()
	exprT := scope.Lookup("Expr").(*types.TypeName)
	exprN := exprT.Type().(*types.Named).Underlying()
	r.astExpr = exprN.(*types.Interface)

	for i, file := range r.pkg.GoFiles {
		dirname, filename := path.Split(file)
		if filename == "ast_format.go" {
			syntax := r.pkg.Syntax[i]
			astutil.Apply(syntax, r.replaceAstfmtCalls, nil)

			f, err := os.Create(path.Join(dirname, "ast_format_fast.go"))
			if err != nil {
				return err
			}
			printer.Fprint(f, r.pkg.Fset, syntax)
			f.Close()
		}
	}
	return nil
}

func (r *Rewriter) replaceAstfmtCalls(cursor *astutil.Cursor) bool {
	switch v := cursor.Node().(type) {
	case *ast.Comment:
		v.Text = strings.ReplaceAll(v.Text, " Format ", " formatFast ")
	case *ast.FuncDecl:
		if v.Name.Name == "Format" {
			v.Name.Name = "formatFast"
		}
	case *ast.ExprStmt:
		if call, ok := v.X.(*ast.CallExpr); ok {
			if r.isPrintfCall(call) {
				return r.rewriteAstPrintf(cursor, call)
			}
		}
	}
	return true
}

func (r *Rewriter) isPrintfCall(n *ast.CallExpr) bool {
	s, ok := n.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id := s.Sel
	if id != nil && !r.pkg.TypesInfo.Types[id].IsType() {
		return id.Name == "astPrintf"
	}
	return false
}

func (r *Rewriter) rewriteLiteral(rcv ast.Expr, method string, arg ast.Expr) ast.Stmt {
	expr := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   rcv,
			Sel: &ast.Ident{Name: method},
		},
		Args: []ast.Expr{arg},
	}
	return &ast.ExprStmt{X: expr}
}

func (r *Rewriter) rewriteAstPrintf(cursor *astutil.Cursor, expr *ast.CallExpr) bool {
	callexpr := expr.Fun.(*ast.SelectorExpr)
	lit := expr.Args[1].(*ast.BasicLit)
	format, _ := strconv.Unquote(lit.Value)

	end := len(format)
	fieldnum := 0
	for i := 0; i < end; {
		lasti := i
		for i < end && format[i] != '%' {
			i++
		}
		if i > lasti {
			var arg ast.Expr
			var method string
			var lit = format[lasti:i]

			if len(lit) == 1 {
				method = "WriteByte"
				arg = &ast.BasicLit{
					Kind:  token.CHAR,
					Value: strconv.QuoteRune(rune(lit[0])),
				}
			} else {
				method = "WriteString"
				arg = &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote(lit),
				}
			}

			cursor.InsertBefore(r.rewriteLiteral(callexpr.X, method, arg))
		}
		if i >= end {
			break
		}
		i++ // '%'
		token := format[i]
		switch token {
		case 'c':
			cursor.InsertBefore(r.rewriteLiteral(callexpr.X, "WriteByte", expr.Args[2+fieldnum]))
		case 's':
			cursor.InsertBefore(r.rewriteLiteral(callexpr.X, "WriteString", expr.Args[2+fieldnum]))
		case 'l', 'r', 'v':
			leftExpr := expr.Args[0]
			leftExprT := r.pkg.TypesInfo.Types[leftExpr].Type

			rightExpr := expr.Args[2+fieldnum]
			rightExprT := r.pkg.TypesInfo.Types[rightExpr].Type

			var call ast.Expr
			if types.Implements(leftExprT, r.astExpr) && types.Implements(rightExprT, r.astExpr) {
				call = &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   callexpr.X,
						Sel: &ast.Ident{Name: "printExpr"},
					},
					Args: []ast.Expr{
						leftExpr,
						rightExpr,
						&ast.Ident{
							Name: strconv.FormatBool(token != 'r'),
						},
					},
				}
			} else {
				call = &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   rightExpr,
						Sel: &ast.Ident{Name: "formatFast"},
					},
					Args: []ast.Expr{callexpr.X},
				}
			}
			cursor.InsertBefore(&ast.ExprStmt{X: call})
		default:
			panic(fmt.Sprintf("unsupported escape %q", token))
		}
		fieldnum++
		i++
	}

	cursor.Delete()
	return true
}
