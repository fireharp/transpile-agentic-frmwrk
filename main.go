package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
	"text/template"
	"time"

	"github.com/fireharp/transpile-agentic-frmwrk/parser"
)

// Templates for generating Python code for PydanticAI and LangChain.
const pydanticAITemplate = `
from pydantic_ai import Agent

agent = Agent(
    "{{.Model}}",
    system_prompt="{{.SystemPrompt}}"
)

result = agent.run_sync("{{.Query}}")
print(result.data)
`

const langChainTemplate = `
from langchain.llms import OpenAI

llm = OpenAI(model_name="{{.Model}}", temperature={{.Temperature}})
response = llm("{{.Query}}")
print(response)
`

// generateCode applies a template to a UniversalAgentSpec.
func generateCode(tmplStr string, spec parser.UniversalAgentSpec) (string, error) {
	tmpl, err := template.New("code").Parse(tmplStr)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, spec); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// --- Runtime execution (runner) functions ---

// runAgent executes the tasks defined in the UniversalAgentSpec concurrently.
func runAgent(spec parser.UniversalAgentSpec) {
	fmt.Printf("Running Agent: %s\nModel: %s\nSystem Prompt: %s\nQuery: %s\n\n",
		spec.Name, spec.Model, spec.SystemPrompt, spec.Query)

	// If there are no tasks, simply print the query result simulation.
	if len(spec.Tasks) == 0 {
		fmt.Println("No tasks defined. Simulating query execution...")
		time.Sleep(1 * time.Second)
		fmt.Println("Query executed.")
		return
	}

	var wg sync.WaitGroup
	results := make([]string, len(spec.Tasks))

	// Process each task concurrently.
	for i, task := range spec.Tasks {
		wg.Add(1)
		go func(index int, t parser.TaskSpec) {
			defer wg.Done()
			switch t.TaskName {
			case "greet":
				msg, ok := t.Params["message"].(string)
				if !ok {
					msg = "Hello, World!"
				}
				results[index] = fmt.Sprintf("Greet Task: %s", msg)
			case "sleep":
				dur, ok := t.Params["duration_sec"].(float64)
				if !ok {
					dur = 1
				}
				time.Sleep(time.Duration(dur) * time.Second)
				results[index] = "Sleep Task Complete"
			case "compute":
				x, okX := t.Params["x"].(float64)
				y, okY := t.Params["y"].(float64)
				if !okX || !okY {
					results[index] = "Compute Error: invalid params"
					return
				}
				results[index] = fmt.Sprintf("Compute Task: %v * %v = %v", x, y, x*y)
			default:
				results[index] = fmt.Sprintf("Unknown Task: %s", t.TaskName)
			}
		}(i, task)
	}
	wg.Wait()

	fmt.Println("All tasks completed. Results:")
	for i, res := range results {
		fmt.Printf("Task %d result: %s\n", i+1, res)
	}
}

// --- Main entry point ---

func main() {
	// Define a mode flag: "transpile", "run", or "export"
	mode := flag.String("mode", "transpile", "Mode: transpile, run, or export")
	// For transpile mode: expect a Go file and target framework.
	goFile := flag.String("gofile", "", "Path to Go source file containing UniversalAgentSpec")
	framework := flag.String("framework", "pydanticai", "Target framework: pydanticai or langchain")
	// For run mode: expect a JSON file containing the spec.
	jsonFile := flag.String("json", "", "Path to JSON file containing agent spec")
	// For export mode: output format
	outFormat := flag.String("format", "json", "Output format for export mode: json or yaml")
	flag.Parse()

	switch *mode {
	case "transpile":
		if *goFile == "" {
			fmt.Println("For transpile mode, please provide a Go file using -gofile")
			return
		}
		// Parse the Go file to extract the agent spec.
		spec, err := parser.ParseGoFile(*goFile)
		if err != nil {
			fmt.Println("Error parsing agent spec:", err)
			return
		}
		fmt.Printf("Parsed Agent Spec: %+v\n", spec)

		// Generate Python code based on the target framework.
		var code string
		if *framework == "pydanticai" {
			code, err = generateCode(pydanticAITemplate, spec)
		} else if *framework == "langchain" {
			code, err = generateCode(langChainTemplate, spec)
		} else {
			fmt.Println("Unsupported framework:", *framework)
			return
		}
		if err != nil {
			fmt.Println("Error generating code:", err)
			return
		}
		fmt.Println("Generated Code:")
		fmt.Println(code)
	case "run":
		if *jsonFile == "" {
			fmt.Println("For run mode, please provide a JSON file using -json")
			return
		}
		// Read and parse the JSON file.
		data, err := os.ReadFile(*jsonFile)
		if err != nil {
			fmt.Printf("Error reading JSON file: %v\n", err)
			return
		}
		var spec parser.UniversalAgentSpec
		if err := json.Unmarshal(data, &spec); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}
		// Run the agent (execute its tasks concurrently).
		runAgent(spec)
	case "export":
		if *goFile == "" {
			fmt.Println("For export mode, please provide a Go file using -gofile")
			return
		}
		// Parse the Go file to extract the agent spec
		spec, err := parser.ParseGoFile(*goFile)
		if err != nil {
			fmt.Println("Error parsing agent spec:", err)
			return
		}

		// Export based on format
		switch *outFormat {
		case "json":
			jsonData, err := json.MarshalIndent(spec, "", "    ")
			if err != nil {
				fmt.Printf("Error marshaling to JSON: %v\n", err)
				return
			}
			fmt.Println(string(jsonData))
		case "yaml":
			fmt.Println("YAML export not implemented yet")
			return
		default:
			fmt.Printf("Unsupported output format: %s\n", *outFormat)
		}
	default:
		fmt.Println("Unknown mode. Use -mode=transpile, -mode=run, or -mode=export")
	}
}
