# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ServicePrincipal']


class ServicePrincipal(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_role_assignment_required: Optional[pulumi.Input[bool]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 oauth2_permissions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ServicePrincipalOauth2PermissionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Service Principal associated with an Application within Azure Active Directory.

        > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API. Please see The Granting a Service Principal permission to manage AAD for the required steps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication",
            homepage="http://homepage",
            identifier_uris=["http://uri"],
            reply_urls=["http://replyurl"],
            available_to_other_tenants=False,
            oauth2_allow_implicit_flow=True)
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal",
            application_id=example_application.application_id,
            app_role_assignment_required=False,
            tags=[
                "example",
                "tags",
                "here",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] app_role_assignment_required: Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
        :param pulumi.Input[str] application_id: The ID of the Azure AD Application for which to create a Service Principal.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ServicePrincipalOauth2PermissionArgs']]]] oauth2_permissions: A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2_permission` block as documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: A list of tags to apply to the Service Principal.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['app_role_assignment_required'] = app_role_assignment_required
            if application_id is None:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['oauth2_permissions'] = oauth2_permissions
            __props__['tags'] = tags
            __props__['display_name'] = None
            __props__['object_id'] = None
        super(ServicePrincipal, __self__).__init__(
            'azuread:index/servicePrincipal:ServicePrincipal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_role_assignment_required: Optional[pulumi.Input[bool]] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            oauth2_permissions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ServicePrincipalOauth2PermissionArgs']]]]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'ServicePrincipal':
        """
        Get an existing ServicePrincipal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] app_role_assignment_required: Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
        :param pulumi.Input[str] application_id: The ID of the Azure AD Application for which to create a Service Principal.
        :param pulumi.Input[str] display_name: The Display Name of the Azure Active Directory Application associated with this Service Principal.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ServicePrincipalOauth2PermissionArgs']]]] oauth2_permissions: A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2_permission` block as documented below.
        :param pulumi.Input[str] object_id: The Service Principal's Object ID.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: A list of tags to apply to the Service Principal.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_role_assignment_required"] = app_role_assignment_required
        __props__["application_id"] = application_id
        __props__["display_name"] = display_name
        __props__["oauth2_permissions"] = oauth2_permissions
        __props__["object_id"] = object_id
        __props__["tags"] = tags
        return ServicePrincipal(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appRoleAssignmentRequired")
    def app_role_assignment_required(self) -> Optional[bool]:
        """
        Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
        """
        return pulumi.get(self, "app_role_assignment_required")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> str:
        """
        The ID of the Azure AD Application for which to create a Service Principal.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The Display Name of the Azure Active Directory Application associated with this Service Principal.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="oauth2Permissions")
    def oauth2_permissions(self) -> List['outputs.ServicePrincipalOauth2Permission']:
        """
        A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2_permission` block as documented below.
        """
        return pulumi.get(self, "oauth2_permissions")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        """
        The Service Principal's Object ID.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[List[str]]:
        """
        A list of tags to apply to the Service Principal.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

