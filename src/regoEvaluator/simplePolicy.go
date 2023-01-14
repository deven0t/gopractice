package main

import (
	"context"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
)

func SimplePolicy() rego.ResultSet {
	ctx := context.Background()

	// Define a simple policy.
	module := `
		package example

		default allow = false

		allow {
			input.identity = "admin"
		}

		allow {
			input.method = "GET"
		}
	`

	// Create a new query that uses the compiled policy from above.
	rego := rego.New(
		rego.Query("data.example.allow"),
		rego.Compiler(getCompiler("example.com", module)),
		rego.Input(
			map[string]interface{}{
				"identity": "bob",
				"method":   "GET",
			},
		),
	)

	// Run evaluation.
	rs, err := rego.Eval(ctx)

	if err != nil {
		// Handle error.
	}
	return rs
	// Inspect results.
	//fmt.Println("len:", len(rs))
	//fmt.Println("value:", rs[0].Expressions[0].Value)
	//fmt.Println("allowed:", rs.Allowed()) // helper method
}

var compilerCache = make(map[string]*ast.Compiler)

func getCompiler(name, module string) *ast.Compiler {
	if value, ok := compilerCache[name]; ok {
		return value
	}
	// Compile the module. The keys are used as identifiers in error messages.
	compiler, _ := ast.CompileModules(map[string]string{
		"example.rego": module,
	})
	compilerCache[name] = compiler
	return compiler
}
