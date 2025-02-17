import pytest
from agent import greet, compute, sleep, tools


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


def test_tools():
    # Test greet tool
    greet_tool = next(tool for tool in tools if tool.name == "greet")
    assert greet_tool.func("Test message") == "Greet Task: Test message"

    # Test compute tool
    compute_tool = next(tool for tool in tools if tool.name == "compute")
    assert compute_tool.func("2.0,3.0") == "Compute Task: 2.0 * 3.0 = 6.0"

    # Test sleep tool
    sleep_tool = next(tool for tool in tools if tool.name == "sleep")
    assert sleep_tool.func("0.1") == "Sleep Task Complete"


def test_tool_descriptions():
    # Verify all tools have descriptions
    for tool in tools:
        assert tool.description, f"Tool {tool.name} missing description"


def test_tool_error_handling():
    compute_tool = next(tool for tool in tools if tool.name == "compute")

    # Test invalid input format
    with pytest.raises(Exception):
        compute_tool.func("invalid")

    # Test invalid number format
    with pytest.raises(ValueError):
        compute_tool.func("invalid,3.0")


if __name__ == "__main__":
    pytest.main([__file__])
