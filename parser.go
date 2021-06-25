package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
)

type TsInterface struct {
	Name   string
	Fields map[string]string
}

func ParseComments(c string) string {
	// Checking if comment can be parsed
	str := SpaceStringsBuilder(c)
	reg := `^@tsInterface(\[context=(\"[A-Za-z]+\")\]|)+$`
	compReg := regexp.MustCompile(reg)
	isNotMatch := !compReg.MatchString(str)
	if isNotMatch {
		fmt.Println("error matching comment: ", str, "\n. Does not match format: ", "@tsInterface[context=\"internal\"] or @tsInterface")
		return ""
	}
	//Checking if there is a context
	if len(str) <= 14 {
		return ""
	} else {
		//Getting context
		contextReg := "\"[A-Za-z]+\""
		compContextReg := regexp.MustCompile(contextReg)
		context := compContextReg.FindString(str)
		context = context[1 : len(context)-1]
		return context
	}
}

//Parse a file to match context and structs. Return batches ( map key= context, v= interfaces)
func ParseGoFile(filename string) map[string][]TsInterface {
	fset := token.NewFileSet()
	coms, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error reaching file:", filename, err)
		return nil
	}
	comMap := make(map[int]string)
	batches := make(map[string][]TsInterface)
	contextSlice := make([]string, 0)

	for _, s := range coms.Comments {
		context := ParseComments(s.Text())
		comMap[fset.Position(s.Pos()).Line] = context
		contextSlice = append(contextSlice, context)
	}
	SetifyString(&contextSlice)

	for _, c := range contextSlice {
		batches[c] = make([]TsInterface, 0)
	}

	for _, node := range coms.Decls {
		switch node.(type) {
		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)
					line := fset.Position(typeSpec.Pos()).Line
					if c, ok := comMap[line-1]; ok {
						name := typeSpec.Name.Name
						tsinterface := TsInterface{
							Name:   name,
							Fields: make(map[string]string),
						}
						switch typeSpec.Type.(type) {
						case *ast.StructType:
							structType := typeSpec.Type.(*ast.StructType)
							for _, field := range structType.Fields.List {
								i := field.Type.(*ast.Ident)
								fieldType := i.Name
								for _, field_name := range field.Names {
									tsinterface.Fields[field_name.Name] = MapToTs(fieldType)
								}
							}

						}
						batches[c] = append(batches[c], tsinterface)
					}
				}
			}
		}
	}
	return batches

}
