from langchain_community.llms import OpenAI
from langchain.agents import Tool
import time


# Task functions
def greet(message: str = "Hello, World!") -> str:
    return f"Greet Task: {message}"


def compute(x: float, y: float) -> str:
    return f"Compute Task: {x} * {y} = {x * y}"


def sleep(duration_sec: float = 1.0) -> str:
    time.sleep(duration_sec)
    return "Sleep Task Complete"


# Create LangChain tools
tools = [
    Tool(
        name="greet",
        func=lambda x: greet(message=x),
        description="Send a greeting message",
    ),
    Tool(
        name="compute",
        func=lambda x: compute(x=float(x.split(",")[0]), y=float(x.split(",")[1])),
        description="Multiply two numbers",
    ),
    Tool(
        name="sleep",
        func=lambda x: sleep(duration_sec=float(x)),
        description="Sleep for a specified duration",
    ),
]

# Create the LLM
llm = OpenAI(model_name="gpt-3.5-turbo", temperature=0.9)

# Example usage
if __name__ == "__main__":
    # Run tasks manually since we're demonstrating task execution
    tasks = [
        {"task_name": "greet", "params": {"message": "Hello from the task runner!"}},
        {"task_name": "compute", "params": {"x": 5.0, "y": 7.0}},
        {"task_name": "sleep", "params": {"duration_sec": 2.0}},
    ]

    # Run the query
    response = llm("Where does 'Hello World' come from?")
    print(f"Query Response: {response}")

    # Run tasks
    results = []
    for i, task in enumerate(tasks, 1):
        task_name = task["task_name"]
        params = task["params"]

        if task_name == "greet":
            result = greet(**params)
        elif task_name == "compute":
            result = compute(**params)
        elif task_name == "sleep":
            result = sleep(**params)
        else:
            result = f"Unknown Task: {task_name}"

        results.append(result)
        print(f"Task {i} result: {result}")
