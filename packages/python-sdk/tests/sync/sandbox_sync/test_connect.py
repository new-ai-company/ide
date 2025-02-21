from e2b import Sandbox


def test_connect(template):
    sbx = Sandbox(template=template, timeout=10, auto_pause=True)
    try:
        assert sbx.is_running()

        sbx_connection = Sandbox.connect(sbx.sandbox_id, auto_pause=True)
        assert sbx_connection.is_running()
    finally:
        sbx.kill()
