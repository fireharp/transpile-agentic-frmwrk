package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestAgentE2E(t *testing.T) {
	// Step 1: Build the sample agent.
	buildCmd := exec.Command("go", "build", "-o", "sample_agent_bin", "./examples/sample_agent.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build sample agent: %v", err)
	}
	defer func() {
		// Clean up the binary after the test.
		_ = exec.Command("rm", "sample_agent_bin").Run()
	}()

	// Step 2: Execute the binary.
	runCmd := exec.Command("./sample_agent_bin")
	var out bytes.Buffer
	runCmd.Stdout = &out
	runCmd.Stderr = &out
	if err := runCmd.Run(); err != nil {
		t.Fatalf("Failed to run sample agent: %v", err)
	}

	// Step 3: Verify the output.
	expectedSubstring := "All tasks completed"
	if !bytes.Contains(out.Bytes(), []byte(expectedSubstring)) {
		t.Errorf("Expected output to contain %q, got: %q", expectedSubstring, out.String())
	}
}
