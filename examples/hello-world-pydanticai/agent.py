from pydantic_ai import Agent
from typing import Dict, Any
import time


def greet(message: str = "Hello, World!") -> str:
    return f"Greet Task: {message}"


def compute(x: float, y: float) -> str:
    return f"Compute Task: {x} * {y} = {x * y}"


def sleep(duration_sec: float = 1.0) -> str:
    time.sleep(duration_sec)
    return "Sleep Task Complete"


# Create the agent
agent = Agent("gpt-3.5-turbo", system_prompt="Be concise and answer in one sentence.")


# Define task handler
def run_tasks(tasks: list[Dict[str, Any]]) -> list[str]:
    results = []
    for task in tasks:
        task_name = task["task_name"]
        params = task["params"]

        if task_name == "greet":
            results.append(greet(**params))
        elif task_name == "compute":
            results.append(compute(**params))
        elif task_name == "sleep":
            results.append(sleep(**params))
        else:
            results.append(f"Unknown Task: {task_name}")

    return results


# Register task handler with agent
run_tasks = agent.task(run_tasks)


# Example usage
if __name__ == "__main__":
    tasks = [
        {"task_name": "greet", "params": {"message": "Hello from the task runner!"}},
        {"task_name": "compute", "params": {"x": 5.0, "y": 7.0}},
        {"task_name": "sleep", "params": {"duration_sec": 2.0}},
    ]

    # Run the tasks
    result = agent.run_sync("Where does 'Hello World' come from?", tasks=tasks)
    print(f"Query Response: {result.data}")

    # Print task results
    for i, res in enumerate(result.tasks, 1):
        print(f"Task {i} result: {res}")
