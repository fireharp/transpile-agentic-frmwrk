package main

import (
	"encoding/json"
	"testing"

	"github.com/fireharp/transpile-agentic-frmwrk/parser"
)

func TestAgentSpec(t *testing.T) {
	// Test that the spec can be marshaled to JSON
	_, err := json.Marshal(AgentSpec)
	if err != nil {
		t.Errorf("Failed to marshal AgentSpec to JSON: %v", err)
	}

	// Test field values
	if AgentSpec.Name != "HelloWorldAgent" {
		t.Errorf("Expected Name to be 'HelloWorldAgent', got '%s'", AgentSpec.Name)
	}
	if AgentSpec.Model != "text-davinci-003" {
		t.Errorf("Expected Model to be 'text-davinci-003', got '%s'", AgentSpec.Model)
	}
	if AgentSpec.Temperature != 0.9 {
		t.Errorf("Expected Temperature to be 0.9, got %f", AgentSpec.Temperature)
	}

	// Test tasks
	if len(AgentSpec.Tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(AgentSpec.Tasks))
	}

	// Test individual tasks
	expectedTasks := []struct {
		name   string
		params map[string]interface{}
	}{
		{
			name: "greet",
			params: map[string]interface{}{
				"message": "Hello from the task runner!",
			},
		},
		{
			name: "compute",
			params: map[string]interface{}{
				"x": 5.0,
				"y": 7.0,
			},
		},
		{
			name: "sleep",
			params: map[string]interface{}{
				"duration_sec": 2.0,
			},
		},
	}

	for i, expected := range expectedTasks {
		task := AgentSpec.Tasks[i]
		if task.TaskName != expected.name {
			t.Errorf("Task %d: expected name '%s', got '%s'", i, expected.name, task.TaskName)
		}
		for k, v := range expected.params {
			if task.Params[k] != v {
				t.Errorf("Task %d: expected param %s=%v, got %v", i, k, v, task.Params[k])
			}
		}
	}
}

func TestParseAgentSpec(t *testing.T) {
	// Test that the spec can be parsed from the file
	spec, err := parser.ParseGoFile("agent_spec.go")
	if err != nil {
		t.Errorf("Failed to parse agent_spec.go: %v", err)
	}

	// Compare parsed spec with original
	if spec.Name != AgentSpec.Name {
		t.Errorf("Parsed Name '%s' doesn't match original '%s'", spec.Name, AgentSpec.Name)
	}
	if spec.Model != AgentSpec.Model {
		t.Errorf("Parsed Model '%s' doesn't match original '%s'", spec.Model, AgentSpec.Model)
	}
	if spec.Temperature != AgentSpec.Temperature {
		t.Errorf("Parsed Temperature %f doesn't match original %f", spec.Temperature, AgentSpec.Temperature)
	}
	if len(spec.Tasks) != len(AgentSpec.Tasks) {
		t.Errorf("Parsed Tasks length %d doesn't match original %d", len(spec.Tasks), len(AgentSpec.Tasks))
	}
}
