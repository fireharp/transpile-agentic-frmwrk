package main

import "github.com/fireharp/transpile-agentic-frmwrk/parser"

// HelloWorldAgent demonstrates a simple agent that runs three tasks:
// 1. Greet - prints a greeting message
// 2. Compute - multiplies two numbers
// 3. Sleep - waits for a specified duration
var AgentSpec = parser.UniversalAgentSpec{
	Name:         "HelloWorldAgent",
	Model:        "text-davinci-003",
	SystemPrompt: "Be concise and answer in one sentence.",
	Query:        "Where does 'Hello World' come from?",
	Temperature:  0.9,
	Tasks: []parser.TaskSpec{
		{
			TaskName: "greet",
			Params: map[string]interface{}{
				"message": "Hello from the task runner!",
			},
		},
		{
			TaskName: "compute",
			Params: map[string]interface{}{
				"x": 5.0,
				"y": 7.0,
			},
		},
		{
			TaskName: "sleep",
			Params: map[string]interface{}{
				"duration_sec": 2.0,
			},
		},
	},
}
