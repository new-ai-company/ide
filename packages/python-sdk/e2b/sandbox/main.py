import urllib.parse

from typing import Optional, Dict

from e2b.sandbox.filesystem.filesystem import Filesystem
from e2b.sandbox.process.main import Process
from e2b.sandbox.sandbox_api import SandboxApi
from e2b.connection_config import ConnectionConfig


class Sandbox(SandboxApi):
    _envd_port = 49982

    @property
    def files(self) -> Filesystem:
        return self._filesystem

    @property
    def commands(self) -> Process:
        return self._process

    def __init__(
        self,
        template: str = "base-v1",
        timeout: Optional[int] = None,
        metadata: Optional[Dict[str, str]] = None,
        api_key: Optional[str] = None,
        domain: Optional[str] = None,
        debug: Optional[bool] = None,
        sandbox_id: Optional[str] = None,
        request_timeout: Optional[float] = None,
    ):
        super().__init__()

        self._connection_config = ConnectionConfig(
            api_key=api_key,
            domain=domain,
            debug=debug,
            request_timeout=request_timeout,
        )

        if sandbox_id is None:
            self.sandbox_id = SandboxApi._create_sandbox(
                template=template,
                api_key=api_key,
                timeout=timeout,
                metadata=metadata,
                domain=domain,
                debug=debug,
                request_timeout=request_timeout,
            )
        else:
            self.sandbox_id = sandbox_id

        self._envd_api_url = (
            f"{'http' if debug else 'https'}://{self.get_host(self._envd_port)}"
        )

        self._filesystem = Filesystem(self._envd_api_url)
        self._process = Process(self._envd_api_url)

    def get_host(self, port: int) -> str:
        if self._connection_config.debug:
            return f"localhost:{port}"

        return f"{port}-{self.sandbox_id}.{self._connection_config.domain}"

    @classmethod
    def connect(
        cls,
        sandbox_id: str,
        api_key: Optional[str] = None,
        domain: Optional[str] = None,
        debug: Optional[bool] = None,
    ):
        return cls(
            sandbox_id=sandbox_id,
            api_key=api_key,
            domain=domain,
            debug=debug,
        )

    def set_timeout(
        self,
        timeout: int,
        request_timeout: Optional[float] = None,
    ) -> None:
        SandboxApi.set_timeout(
            sandbox_id=self.sandbox_id,
            timeout=timeout,
            **self._connection_config.__dict__,
            request_timeout=request_timeout,
        )

    def kill(self, request_timeout: Optional[float] = None) -> None:
        SandboxApi.kill(
            sandbox_id=self.sandbox_id,
            **self._connection_config.__dict__,
            request_timeout=request_timeout,
        )

    @property
    def upload_url(self) -> str:
        url = urllib.parse.urljoin(self._envd_api_url, "/files?")
        params = urllib.parse.urlencode(
            {"user": "user"},
        )
        url = urllib.parse.urljoin(url, params)

        return url
