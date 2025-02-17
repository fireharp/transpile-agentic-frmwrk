# Hello World Agent - LangChain Implementation

This example shows how to implement the Hello World agent using LangChain. It demonstrates:

1. Tool-based task definitions
2. LLM configuration
3. Task execution
4. Query handling

## Requirements

```bash
pip install langchain openai
```

You'll also need to set your OpenAI API key:

```bash
export OPENAI_API_KEY=your_api_key_here
```

## Usage

### 1. Run directly:

```bash
python agent.py
```

### 2. Generate from Go spec:

```bash
go run ../../main.go -mode=transpile -gofile=../hello-world-go/agent_spec.go -framework=langchain > generated_agent.py
```

## Code Structure

- `agent.py` - Main implementation showing:
  - Task function definitions (`greet`, `compute`, `sleep`)
  - LangChain Tool definitions
  - LLM configuration
  - Example usage

## Expected Output

```
Query Response: The term "Hello, World!" originated from early computer programming tutorials...
Task 1 result: Greet Task: Hello from the task runner!
Task 2 result: Compute Task: 5 * 7 = 35
Task 3 result: Sleep Task Complete
```

## Comparison with Other Versions

This implementation:

1. Uses LangChain's Tool system for task definitions
2. Provides more explicit LLM configuration
3. Shows how to adapt the task structure to LangChain's patterns
4. Maintains compatibility with the Go and PydanticAI versions
