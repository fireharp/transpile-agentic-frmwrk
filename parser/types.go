package parser

// UniversalAgentSpec is our framework-agnostic agent specification.
type UniversalAgentSpec struct {
	Name         string     `json:"name"`
	Model        string     `json:"model"`
	SystemPrompt string     `json:"system_prompt"`
	Query        string     `json:"query"`
	Temperature  float64    `json:"temperature"`
	Tasks        []TaskSpec `json:"tasks,omitempty"`
}

// TaskSpec represents an individual task with a name and arbitrary parameters.
type TaskSpec struct {
	TaskName string                 `json:"task_name"`
	Params   map[string]interface{} `json:"params"`
} 