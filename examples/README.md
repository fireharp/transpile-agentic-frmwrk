# Agent Examples

This directory contains various examples of agent specifications and how to use them with the transpiler.

## Examples

### Hello World Implementations

Three equivalent implementations of a simple agent that performs basic tasks:

- [Go Version](./hello-world-go) - The reference implementation using our Go spec format
- [PydanticAI Version](./hello-world-pydanticai) - Implementation using PydanticAI framework
- [LangChain Version](./hello-world-langchain) - Implementation using LangChain framework

Each Hello World example demonstrates:

- Greeting
- Basic computation
- Sleep/delay
- Query handling

## Adding New Examples

Each example should:

1. Be in its own directory
2. Have a descriptive README.md explaining what it does and how to use it
3. Include the agent specification in Go format
4. (Optional) Include framework-specific implementations
5. (Optional) Include pre-exported JSON/YAML versions
6. (Optional) Include example outputs from different frameworks

## Running Examples

All examples can be run using the main transpiler binary. From the example's directory:

```bash
# Run directly from Go spec
go run ../../main.go -mode=run -gofile=agent_spec.go

# Export to JSON
go run ../../main.go -mode=export -gofile=agent_spec.go > agent_spec.json

# Run from JSON
go run ../../main.go -mode=run -json=agent_spec.json

# Generate Python code
go run ../../main.go -mode=transpile -gofile=agent_spec.go -framework=pydanticai
go run ../../main.go -mode=transpile -gofile=agent_spec.go -framework=langchain
```

## Framework Support

Currently supported frameworks:

1. PydanticAI - Modern Python framework with strong typing
2. LangChain - Popular framework with extensive tool support

Each framework has its own strengths:

- **PydanticAI**: Better for type safety and simpler task definitions
- **LangChain**: Better for complex workflows and tool integration
