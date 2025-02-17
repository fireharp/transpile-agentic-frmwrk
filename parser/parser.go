package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

// ParseGoFile parses a Go source file and extracts a UniversalAgentSpec
// by looking for a composite literal with type "UniversalAgentSpec".
func ParseGoFile(filename string) (UniversalAgentSpec, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return UniversalAgentSpec{}, err
	}
	var spec UniversalAgentSpec
	ast.Inspect(node, func(n ast.Node) bool {
		cl, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}
		// Check if the composite literal's type is "UniversalAgentSpec" or "parser.UniversalAgentSpec"
		if ident, ok := cl.Type.(*ast.Ident); ok && ident.Name == "UniversalAgentSpec" {
			// Handle unqualified type name
			parseFields(cl, &spec)
			return false
		} else if sel, ok := cl.Type.(*ast.SelectorExpr); ok {
			// Handle qualified type name (e.g., parser.UniversalAgentSpec)
			if pkg, ok := sel.X.(*ast.Ident); ok && pkg.Name == "parser" && sel.Sel.Name == "UniversalAgentSpec" {
				parseFields(cl, &spec)
				return false
			}
		}
		return true
	})
	return spec, nil
}

// parseFields extracts field values from a composite literal into a UniversalAgentSpec
func parseFields(cl *ast.CompositeLit, spec *UniversalAgentSpec) {
	for _, elt := range cl.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}
		switch key.Name {
		case "Name":
			if lit, ok := kv.Value.(*ast.BasicLit); ok {
				spec.Name, _ = strconv.Unquote(lit.Value)
			}
		case "Model":
			if lit, ok := kv.Value.(*ast.BasicLit); ok {
				spec.Model, _ = strconv.Unquote(lit.Value)
			}
		case "SystemPrompt":
			if lit, ok := kv.Value.(*ast.BasicLit); ok {
				spec.SystemPrompt, _ = strconv.Unquote(lit.Value)
			}
		case "Query":
			if lit, ok := kv.Value.(*ast.BasicLit); ok {
				spec.Query, _ = strconv.Unquote(lit.Value)
			}
		case "Temperature":
			if lit, ok := kv.Value.(*ast.BasicLit); ok {
				if val, err := strconv.ParseFloat(lit.Value, 64); err == nil {
					spec.Temperature = val
				}
			}
		case "Tasks":
			if taskList, ok := kv.Value.(*ast.CompositeLit); ok {
				for _, taskElt := range taskList.Elts {
					if taskComp, ok := taskElt.(*ast.CompositeLit); ok {
						var task TaskSpec
						for _, taskField := range taskComp.Elts {
							if taskKV, ok := taskField.(*ast.KeyValueExpr); ok {
								if taskKey, ok := taskKV.Key.(*ast.Ident); ok {
									switch taskKey.Name {
									case "TaskName":
										if lit, ok := taskKV.Value.(*ast.BasicLit); ok {
											task.TaskName, _ = strconv.Unquote(lit.Value)
										}
									case "Params":
										if paramComp, ok := taskKV.Value.(*ast.CompositeLit); ok {
											params := make(map[string]interface{})
											for _, paramElt := range paramComp.Elts {
												if paramKV, ok := paramElt.(*ast.KeyValueExpr); ok {
													if paramKey, ok := paramKV.Key.(*ast.BasicLit); ok {
														key, _ := strconv.Unquote(paramKey.Value)
														if paramVal, ok := paramKV.Value.(*ast.BasicLit); ok {
															switch paramVal.Kind {
															case token.FLOAT:
																params[key], _ = strconv.ParseFloat(paramVal.Value, 64)
															case token.STRING:
																params[key], _ = strconv.Unquote(paramVal.Value)
															}
														}
													}
												}
											}
											task.Params = params
										}
									}
								}
							}
						}
						spec.Tasks = append(spec.Tasks, task)
					}
				}
			}
		}
	}
}
