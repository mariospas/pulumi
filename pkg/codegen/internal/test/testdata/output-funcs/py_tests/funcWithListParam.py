# coding=utf-8
# *** WARNING: this file was generated by . ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'FuncWithListParamResult',
    'AwaitableFuncWithListParamResult',
    'func_with_list_param',
    'func_with_list_param_output',
]

@pulumi.output_type
class FuncWithListParamResult:
    def __init__(__self__, r=None):
        if r and not isinstance(r, str):
            raise TypeError("Expected argument 'r' to be a str")
        pulumi.set(__self__, "r", r)

    @property
    @pulumi.getter
    def r(self) -> str:
        return pulumi.get(self, "r")


class AwaitableFuncWithListParamResult(FuncWithListParamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return FuncWithListParamResult(
            r=self.r)


def func_with_list_param(a: Optional[Sequence[str]] = None,
                         b: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableFuncWithListParamResult:
    """
    Check codegen of functions with a List parameter.
    """
    __args__ = dict()
    __args__['a'] = a
    __args__['b'] = b
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('madeup-package:codegentest:funcWithListParam', __args__, opts=opts, typ=FuncWithListParamResult).value

    return AwaitableFuncWithListParamResult(
        r=__ret__.r)


@_utilities.lift_output_func(func_with_list_param)
def func_with_list_param_output(a: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                b: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[FuncWithListParamResult]:
    ...
