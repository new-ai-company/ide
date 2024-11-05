import pytest

from e2b import Sandbox


@pytest.mark.skip("Old envd")
def test_command_envs(sandbox: Sandbox):
    cmd = sandbox.commands.run("echo $FOO", envs={"FOO": "bar"})
    assert cmd.stdout.strip() == "bar"


@pytest.mark.skip_debug()
@pytest.mark.skip("Old envd")
def test_sandbox_envs(template: str):
    sandbox = Sandbox(template, envs={"FOO": "bar"})
    try:
        cmd = sandbox.commands.run("echo $FOO")
        assert cmd.stdout.strip() == "bar"
    finally:
        sandbox.kill()
