# Hello World Agent Example

This example demonstrates a simple agent that performs three basic tasks:

1. **Greet Task**: Prints a greeting message
2. **Compute Task**: Multiplies two numbers
3. **Sleep Task**: Waits for a specified duration

## Usage

You can use this example in different ways:

### 1. Run the agent directly:

```bash
go run ../../main.go -mode=run -gofile=agent_spec.go
```

### 2. Export to JSON:

```bash
go run ../../main.go -mode=export -gofile=agent_spec.go > agent_spec.json
```

### 3. Run from exported JSON:

```bash
go run ../../main.go -mode=run -json=agent_spec.json
```

### 4. Generate Python code:

```bash
# Using PydanticAI
go run ../../main.go -mode=transpile -gofile=agent_spec.go -framework=pydanticai

# Using LangChain
go run ../../main.go -mode=transpile -gofile=agent_spec.go -framework=langchain
```

## Expected Output

When running the agent, you should see output similar to:

```
Running Agent: HelloWorldAgent
Model: text-davinci-003
System Prompt: Be concise and answer in one sentence.
Query: Where does 'Hello World' come from?

All tasks completed. Results:
Task 1 result: Greet Task: Hello from the task runner!
Task 2 result: Compute Task: 5 * 7 = 35
Task 3 result: Sleep Task Complete
```
