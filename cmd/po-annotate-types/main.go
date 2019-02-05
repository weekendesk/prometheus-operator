package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"go/ast"
	"io/ioutil"

	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

type typeCommentAppender struct {
	commentAllTypes bool
	targetTypes     []string
	comment         string
}

func (a typeCommentAppender) Apply(pkg *ast.Package, file *ast.File) *ast.File {
	comments := []*ast.CommentGroup{}
	ast.Inspect(file, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.CommentGroup:
			c, _ := node.(*ast.CommentGroup)
			comments = append(comments, c)
		case *ast.GenDecl:
			if n.Tok != token.TYPE {
				break
			}

			gd, _ := node.(*ast.GenDecl)
			if ts, ok := gd.Specs[0].(*ast.TypeSpec); ok {
				if ts.Name.IsExported() && a.isTargetType(ts) && !strings.Contains(gd.Doc.Text(), a.comment) {
					var commentPosition token.Pos

					if gd.Doc == nil || len(gd.Doc.List) == 0 {
						commentPosition = gd.Pos() - 1
						gd.Doc = &ast.CommentGroup{
							List: []*ast.Comment{},
						}
					} else {
						commentPosition = gd.Doc.End() + 1
					}

					gd.Doc = &ast.CommentGroup{
						List: append(gd.Doc.List, &ast.Comment{
							Text:  "// " + a.comment,
							Slash: commentPosition,
						}),
					}
				}
			}
		}
		return true
	})
	// rewrite comments
	file.Comments = comments
	return file
}

func (a typeCommentAppender) isTargetType(ts *ast.TypeSpec) bool {
	return a.commentAllTypes || find(a.targetTypes, ts.Name.Name)
}
func find(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	var path string
	var targetTypes []string
	var comment string

	pflag.StringVar(&path, "path-to-package", "vendor/k8s.io/client-go/pkg/api/v1", "path in which to find types to annotate")
	pflag.StringArrayVar(&targetTypes, "type", []string{}, "name of a type to annotate (type must be exported)")
	pflag.StringVar(&comment, "annotation", "+k8s:deepcopy-gen=true", "the annotation to apply to the target types")
	pflag.Parse()

	fmt.Printf("adding annotation \"%v\" to the following types of %v: %v\n", comment, path, targetTypes)

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	if len(packages) != 1 {
		var pkgs []string

		for k := range packages {
			pkgs = append(pkgs, k)
		}

		panic(fmt.Errorf(fmt.Sprintf("expected 1 package in %v, found the %v followings: %v", path, len(pkgs), pkgs)))
	}

	appendComment := typeCommentAppender{
		commentAllTypes: len(targetTypes) == 0,
		targetTypes:     targetTypes,
		comment:         comment,
	}

	for _, pkg := range packages {
		for fileName, node := range pkg.Files {
			rewriteSourceFile(
				fileName,
				fset,
				appendComment.Apply(pkg, node),
			)
		}
	}
}

func rewriteSourceFile(path string, fset *token.FileSet, node ast.Node) {
	code, _ := asGoCode(fset, node)
	ioutil.WriteFile(path, []byte(code), 0644)
}

func asGoCode(fset *token.FileSet, node ast.Node) (string, error) {
	buf := new(bytes.Buffer)
	if err := printer.Fprint(buf, fset, node); err != nil {
		return "", err
	}

	str := fmt.Sprint(buf)
	return str, nil
}
