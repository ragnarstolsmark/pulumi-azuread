// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package azuread

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Application within Azure Active Directory.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write owned by applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/r/application.html.markdown.
type Application struct {
	pulumi.CustomResourceState

	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles ApplicationAppRoleArrayOutput `pulumi:"appRoles"`
	// The Application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrOutput `pulumi:"availableToOtherTenants"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
	GroupMembershipClaims pulumi.StringPtrOutput `pulumi:"groupMembershipClaims"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringOutput `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayOutput `pulumi:"identifierUris"`
	// The URL of the logout page.
	LogoutUrl pulumi.StringPtrOutput `pulumi:"logoutUrl"`
	// The display name for the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrOutput `pulumi:"oauth2AllowImplicitFlow"`
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions ApplicationOauth2PermissionArrayOutput `pulumi:"oauth2Permissions"`
	// The Application's Object ID.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient pulumi.BoolOutput `pulumi:"publicClient"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayOutput `pulumi:"replyUrls"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayOutput `pulumi:"requiredResourceAccesses"`
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}
	var resource Application
	err := ctx.RegisterResource("azuread:index/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azuread:index/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles []ApplicationAppRole `pulumi:"appRoles"`
	// The Application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants *bool `pulumi:"availableToOtherTenants"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
	GroupMembershipClaims *string `pulumi:"groupMembershipClaims"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage *string `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// The URL of the logout page.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// The display name for the application.
	Name *string `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow *bool `pulumi:"oauth2AllowImplicitFlow"`
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions []ApplicationOauth2Permission `pulumi:"oauth2Permissions"`
	// The Application's Object ID.
	ObjectId *string `pulumi:"objectId"`
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners []string `pulumi:"owners"`
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient *bool `pulumi:"publicClient"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls []string `pulumi:"replyUrls"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses []ApplicationRequiredResourceAccess `pulumi:"requiredResourceAccesses"`
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type *string `pulumi:"type"`
}

type ApplicationState struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles ApplicationAppRoleArrayInput
	// The Application ID.
	ApplicationId pulumi.StringPtrInput
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrInput
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
	GroupMembershipClaims pulumi.StringPtrInput
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringPtrInput
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// The URL of the logout page.
	LogoutUrl pulumi.StringPtrInput
	// The display name for the application.
	Name pulumi.StringPtrInput
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrInput
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions ApplicationOauth2PermissionArrayInput
	// The Application's Object ID.
	ObjectId pulumi.StringPtrInput
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners pulumi.StringArrayInput
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient pulumi.BoolPtrInput
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayInput
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles []ApplicationAppRole `pulumi:"appRoles"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants *bool `pulumi:"availableToOtherTenants"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
	GroupMembershipClaims *string `pulumi:"groupMembershipClaims"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage *string `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// The URL of the logout page.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// The display name for the application.
	Name *string `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow *bool `pulumi:"oauth2AllowImplicitFlow"`
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions []ApplicationOauth2Permission `pulumi:"oauth2Permissions"`
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners []string `pulumi:"owners"`
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient *bool `pulumi:"publicClient"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls []string `pulumi:"replyUrls"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses []ApplicationRequiredResourceAccess `pulumi:"requiredResourceAccesses"`
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles ApplicationAppRoleArrayInput
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrInput
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.
	GroupMembershipClaims pulumi.StringPtrInput
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringPtrInput
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// The URL of the logout page.
	LogoutUrl pulumi.StringPtrInput
	// The display name for the application.
	Name pulumi.StringPtrInput
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrInput
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions ApplicationOauth2PermissionArrayInput
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners pulumi.StringArrayInput
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient pulumi.BoolPtrInput
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayInput
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
