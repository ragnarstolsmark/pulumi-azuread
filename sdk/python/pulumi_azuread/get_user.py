# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, account_enabled=None, display_name=None, id=None, immutable_id=None, mail=None, mail_nickname=None, object_id=None, onpremises_sam_account_name=None, onpremises_user_principal_name=None, usage_location=None, user_principal_name=None):
        if account_enabled and not isinstance(account_enabled, bool):
            raise TypeError("Expected argument 'account_enabled' to be a bool")
        __self__.account_enabled = account_enabled
        """
        `True` if the account is enabled; otherwise `False`.
        """
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        The Display Name of the Azure AD User.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if immutable_id and not isinstance(immutable_id, str):
            raise TypeError("Expected argument 'immutable_id' to be a str")
        __self__.immutable_id = immutable_id
        """
        The value used to associate an on-premises Active Directory user account with their Azure AD user object.
        """
        if mail and not isinstance(mail, str):
            raise TypeError("Expected argument 'mail' to be a str")
        __self__.mail = mail
        """
        The primary email address of the Azure AD User.
        """
        if mail_nickname and not isinstance(mail_nickname, str):
            raise TypeError("Expected argument 'mail_nickname' to be a str")
        __self__.mail_nickname = mail_nickname
        """
        The email alias of the Azure AD User.
        """
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        __self__.object_id = object_id
        if onpremises_sam_account_name and not isinstance(onpremises_sam_account_name, str):
            raise TypeError("Expected argument 'onpremises_sam_account_name' to be a str")
        __self__.onpremises_sam_account_name = onpremises_sam_account_name
        """
        The on premise sam account name of the Azure AD User.
        """
        if onpremises_user_principal_name and not isinstance(onpremises_user_principal_name, str):
            raise TypeError("Expected argument 'onpremises_user_principal_name' to be a str")
        __self__.onpremises_user_principal_name = onpremises_user_principal_name
        """
        The on premise user principal name of the Azure AD User.
        """
        if usage_location and not isinstance(usage_location, str):
            raise TypeError("Expected argument 'usage_location' to be a str")
        __self__.usage_location = usage_location
        """
        The usage location of the Azure AD User.
        """
        if user_principal_name and not isinstance(user_principal_name, str):
            raise TypeError("Expected argument 'user_principal_name' to be a str")
        __self__.user_principal_name = user_principal_name
        """
        The User Principal Name of the Azure AD User.
        """
class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            account_enabled=self.account_enabled,
            display_name=self.display_name,
            id=self.id,
            immutable_id=self.immutable_id,
            mail=self.mail,
            mail_nickname=self.mail_nickname,
            object_id=self.object_id,
            onpremises_sam_account_name=self.onpremises_sam_account_name,
            onpremises_user_principal_name=self.onpremises_user_principal_name,
            usage_location=self.usage_location,
            user_principal_name=self.user_principal_name)

def get_user(mail_nickname=None,object_id=None,user_principal_name=None,opts=None):
    """
    Gets information about an Azure Active Directory user.

    > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.




    :param str mail_nickname: The email alias of the Azure AD User.
    :param str object_id: Specifies the Object ID of the Application within Azure Active Directory.
    :param str user_principal_name: The User Principal Name of the Azure AD User.
    """
    __args__ = dict()


    __args__['mailNickname'] = mail_nickname
    __args__['objectId'] = object_id
    __args__['userPrincipalName'] = user_principal_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuread:index/getUser:getUser', __args__, opts=opts).value

    return AwaitableGetUserResult(
        account_enabled=__ret__.get('accountEnabled'),
        display_name=__ret__.get('displayName'),
        id=__ret__.get('id'),
        immutable_id=__ret__.get('immutableId'),
        mail=__ret__.get('mail'),
        mail_nickname=__ret__.get('mailNickname'),
        object_id=__ret__.get('objectId'),
        onpremises_sam_account_name=__ret__.get('onpremisesSamAccountName'),
        onpremises_user_principal_name=__ret__.get('onpremisesUserPrincipalName'),
        usage_location=__ret__.get('usageLocation'),
        user_principal_name=__ret__.get('userPrincipalName'))
