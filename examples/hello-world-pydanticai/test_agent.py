import pytest
from unittest.mock import patch, MagicMock

# Configure mock to return the original function when task decorator is used
mock_agent = MagicMock()
mock_agent.task = lambda f: f

with patch("pydantic_ai.Agent", return_value=mock_agent):
    from agent import greet, compute, sleep, run_tasks


def test_greet():
    # Test default message
    assert greet() == "Greet Task: Hello, World!"
    # Test custom message
    assert greet("Custom message") == "Greet Task: Custom message"


def test_compute():
    # Test integer multiplication
    assert compute(5, 7) == "Compute Task: 5 * 7 = 35"
    # Test float multiplication
    assert compute(3.5, 2) == "Compute Task: 3.5 * 2 = 7.0"


def test_sleep():
    import time

    start = time.time()
    result = sleep(0.1)  # Use small duration for tests
    duration = time.time() - start
    assert result == "Sleep Task Complete"
    assert duration >= 0.1  # Should have slept at least 0.1 seconds


def test_run_tasks():
    tasks = [
        {"task_name": "greet", "params": {"message": "Test message"}},
        {"task_name": "compute", "params": {"x": 2.0, "y": 3.0}},
        {"task_name": "sleep", "params": {"duration_sec": 0.1}},
    ]

    results = run_tasks(tasks)
    assert len(results) == 3
    assert results[0] == "Greet Task: Test message"
    assert results[1] == "Compute Task: 2.0 * 3.0 = 6.0"
    assert results[2] == "Sleep Task Complete"


def test_unknown_task():
    tasks = [{"task_name": "unknown", "params": {}}]

    results = run_tasks(tasks)
    assert len(results) == 1
    assert results[0] == "Unknown Task: unknown"


if __name__ == "__main__":
    pytest.main([__file__])
