package runner

import (
	"fmt"
	"sync"
	"time"
)

// UniversalAgentSpec defines the common specification for an agent.
type UniversalAgentSpec struct {
	Name         string
	Model        string
	SystemPrompt string
	Query        string
	Temperature  float64
}

// RunAgent simulates running an agent by processing its query asynchronously.
func RunAgent(spec UniversalAgentSpec) error {
	fmt.Printf("Running agent: %s with model: %s\n", spec.Name, spec.Model)
	
	// Example: use a goroutine to simulate asynchronous work.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Simulate task processing based on the agent spec.
		fmt.Printf("Agent %s processing query: %s\n", spec.Name, spec.Query)
		time.Sleep(2 * time.Second)
		fmt.Printf("Agent %s finished processing\n", spec.Name)
	}()
	wg.Wait()
	return nil
}
