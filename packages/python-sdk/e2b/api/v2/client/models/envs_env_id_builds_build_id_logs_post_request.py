# coding: utf-8

"""
    E2B API

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: 0.1.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import List
from pydantic import BaseModel, Field, StrictStr, conlist


class EnvsEnvIDBuildsBuildIDLogsPostRequest(BaseModel):
    """
    EnvsEnvIDBuildsBuildIDLogsPostRequest
    """

    api_secret: StrictStr = Field(..., alias="apiSecret", description="API secret")
    logs: conlist(StrictStr) = Field(...)

    """Pydantic configuration"""
    model_config = {
        "populate_by_name": True,
        "validate_assignment": True,
    }

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> EnvsEnvIDBuildsBuildIDLogsPostRequest:
        """Create an instance of EnvsEnvIDBuildsBuildIDLogsPostRequest from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.model_dump(by_alias=True, exclude={}, exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> EnvsEnvIDBuildsBuildIDLogsPostRequest:
        """Create an instance of EnvsEnvIDBuildsBuildIDLogsPostRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return EnvsEnvIDBuildsBuildIDLogsPostRequest.model_validate(obj)

        # raise errors for additional fields in the input
        for _key in obj.keys():
            if _key not in ["apiSecret", "logs"]:
                raise ValueError(
                    "Error due to additional fields (not defined in EnvsEnvIDBuildsBuildIDLogsPostRequest) in the input: "
                    + obj
                )

        _obj = EnvsEnvIDBuildsBuildIDLogsPostRequest.model_validate(
            {"api_secret": obj.get("apiSecret"), "logs": obj.get("logs")}
        )
        return _obj
