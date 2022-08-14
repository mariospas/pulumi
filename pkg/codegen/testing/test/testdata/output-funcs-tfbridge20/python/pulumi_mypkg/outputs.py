# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'StorageAccountKeyResponseResult',
    'GetAmiIdsFilterResult',
]

@pulumi.output_type
class StorageAccountKeyResponseResult(dict):
    """
    An access key for the storage account.
    """
    def __init__(__self__, *,
                 creation_time: str,
                 key_name: str,
                 permissions: str,
                 value: str):
        """
        An access key for the storage account.
        :param str creation_time: Creation time of the key, in round trip date format.
        :param str key_name: Name of the key.
        :param str permissions: Permissions for the key -- read-only or full permissions.
        :param str value: Base 64-encoded value of the key.
        """
        pulumi.set(__self__, "creation_time", creation_time)
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Creation time of the key, in round trip date format.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        Name of the key.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def permissions(self) -> str:
        """
        Permissions for the key -- read-only or full permissions.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Base 64-encoded value of the key.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetAmiIdsFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")


