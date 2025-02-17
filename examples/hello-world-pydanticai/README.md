# Hello World Agent - PydanticAI Implementation

This example shows how to implement the Hello World agent using PydanticAI. It demonstrates:

1. Task function definitions with type hints
2. Agent task decorators
3. Concurrent task execution
4. Query handling

## Requirements

```bash
pip install pydantic-ai
```

## Usage

### 1. Run directly:

```bash
python agent.py
```

### 2. Generate from Go spec:

```bash
go run ../../main.go -mode=transpile -gofile=../hello-world-go/agent_spec.go -framework=pydanticai > generated_agent.py
```

## Code Structure

- `agent.py` - Main implementation showing:
  - Task function definitions (`greet`, `compute`, `sleep`)
  - Agent configuration
  - Task handler with decorator
  - Example usage

## Expected Output

```
Query Response: The term "Hello, World!" originated from early computer programming tutorials...
Task 1 result: Greet Task: Hello from the task runner!
Task 2 result: Compute Task: 5 * 7 = 35
Task 3 result: Sleep Task Complete
```

## Comparison with Go Version

This implementation:

1. Uses Python's type hints for better code clarity
2. Leverages PydanticAI's task decorators for simpler task handling
3. Maintains the same task structure and parameters
4. Provides similar output format
