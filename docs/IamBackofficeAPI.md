# GeminiCommerce\Iambackoffice\IamBackofficeAPI

All URIs are relative to *https://iambackoffice.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamBackofficeAssignRoles**](IamBackofficeAPI.md#IamBackofficeAssignRoles) | **Post** /iambackoffice.IamBackoffice/AssignRoles | AUTHZ
[**IamBackofficeAssignUserToGroup**](IamBackofficeAPI.md#IamBackofficeAssignUserToGroup) | **Post** /iambackoffice.IamBackoffice/AssignUserToGroup | 
[**IamBackofficeCreateGroup**](IamBackofficeAPI.md#IamBackofficeCreateGroup) | **Post** /iambackoffice.IamBackoffice/CreateGroup | GROUPS
[**IamBackofficeDisableUserMfa**](IamBackofficeAPI.md#IamBackofficeDisableUserMfa) | **Post** /iambackoffice.IamBackoffice/DisableUserMfa | 
[**IamBackofficeEnableUserMfa**](IamBackofficeAPI.md#IamBackofficeEnableUserMfa) | **Post** /iambackoffice.IamBackoffice/EnableUserMfa | 
[**IamBackofficeGenerateSecretForQR**](IamBackofficeAPI.md#IamBackofficeGenerateSecretForQR) | **Post** /iambackoffice.IamBackoffice/GenerateSecretForQR | MFA
[**IamBackofficeGetUser**](IamBackofficeAPI.md#IamBackofficeGetUser) | **Post** /iambackoffice.IamBackoffice/GetUser | USER
[**IamBackofficeLogin**](IamBackofficeAPI.md#IamBackofficeLogin) | **Post** /iambackoffice.IamBackoffice/Login | LOGIN
[**IamBackofficeLoginMfa**](IamBackofficeAPI.md#IamBackofficeLoginMfa) | **Post** /iambackoffice.IamBackoffice/LoginMfa | LOGIN MFA
[**IamBackofficeLogout**](IamBackofficeAPI.md#IamBackofficeLogout) | **Post** /iambackoffice.IamBackoffice/Logout | 
[**IamBackofficeRegister**](IamBackofficeAPI.md#IamBackofficeRegister) | **Post** /iambackoffice.IamBackoffice/Register | REGISTRATION
[**IamBackofficeSearchGroups**](IamBackofficeAPI.md#IamBackofficeSearchGroups) | **Post** /iambackoffice.IamBackoffice/SearchGroups | 
[**IamBackofficeUnassignRoles**](IamBackofficeAPI.md#IamBackofficeUnassignRoles) | **Post** /iambackoffice.IamBackoffice/UnassignRoles | 



## IamBackofficeAssignRoles

> IambackofficeAssignRolesResponse IamBackofficeAssignRoles(ctx).Body(body).Execute()

AUTHZ

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeAssignRolesRequest() // IambackofficeAssignRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeAssignRoles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeAssignRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeAssignRoles`: IambackofficeAssignRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeAssignRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeAssignRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeAssignRolesRequest**](IambackofficeAssignRolesRequest.md) |  | 

### Return type

[**IambackofficeAssignRolesResponse**](IambackofficeAssignRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeAssignUserToGroup

> map[string]interface{} IamBackofficeAssignUserToGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeAssignUserToGroupRequest() // IambackofficeAssignUserToGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeAssignUserToGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeAssignUserToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeAssignUserToGroup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeAssignUserToGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeAssignUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeAssignUserToGroupRequest**](IambackofficeAssignUserToGroupRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeCreateGroup

> IambackofficeCreateGroupResponse IamBackofficeCreateGroup(ctx).Body(body).Execute()

GROUPS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeCreateGroupRequest() // IambackofficeCreateGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeCreateGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeCreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeCreateGroup`: IambackofficeCreateGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeCreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeCreateGroupRequest**](IambackofficeCreateGroupRequest.md) |  | 

### Return type

[**IambackofficeCreateGroupResponse**](IambackofficeCreateGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeDisableUserMfa

> IambackofficeUserMfaResponse IamBackofficeDisableUserMfa(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeDisableUserMfaRequest() // IambackofficeDisableUserMfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeDisableUserMfa(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeDisableUserMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeDisableUserMfa`: IambackofficeUserMfaResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeDisableUserMfa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeDisableUserMfaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeDisableUserMfaRequest**](IambackofficeDisableUserMfaRequest.md) |  | 

### Return type

[**IambackofficeUserMfaResponse**](IambackofficeUserMfaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeEnableUserMfa

> IambackofficeUserMfaResponse IamBackofficeEnableUserMfa(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeEnableUserMfaRequest() // IambackofficeEnableUserMfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeEnableUserMfa(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeEnableUserMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeEnableUserMfa`: IambackofficeUserMfaResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeEnableUserMfa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeEnableUserMfaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeEnableUserMfaRequest**](IambackofficeEnableUserMfaRequest.md) |  | 

### Return type

[**IambackofficeUserMfaResponse**](IambackofficeUserMfaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeGenerateSecretForQR

> IambackofficeGenerateSecretForQRResponse IamBackofficeGenerateSecretForQR(ctx).Body(body).Execute()

MFA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeGenerateSecretForQR(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeGenerateSecretForQR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeGenerateSecretForQR`: IambackofficeGenerateSecretForQRResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeGenerateSecretForQR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeGenerateSecretForQRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**IambackofficeGenerateSecretForQRResponse**](IambackofficeGenerateSecretForQRResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeGetUser

> IambackofficeGetUserResponse IamBackofficeGetUser(ctx).Body(body).Execute()

USER

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeGetUserRequest() // IambackofficeGetUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeGetUser(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeGetUser`: IambackofficeGetUserResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeGetUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeGetUserRequest**](IambackofficeGetUserRequest.md) |  | 

### Return type

[**IambackofficeGetUserResponse**](IambackofficeGetUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeLogin

> IambackofficeLoginResponse IamBackofficeLogin(ctx).Body(body).Execute()

LOGIN

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeLoginRequest() // IambackofficeLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeLogin(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeLogin`: IambackofficeLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeLoginRequest**](IambackofficeLoginRequest.md) |  | 

### Return type

[**IambackofficeLoginResponse**](IambackofficeLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeLoginMfa

> IambackofficeLoginResponse IamBackofficeLoginMfa(ctx).Body(body).Execute()

LOGIN MFA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeLoginMfaRequest() // IambackofficeLoginMfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeLoginMfa(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeLoginMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeLoginMfa`: IambackofficeLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeLoginMfa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeLoginMfaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeLoginMfaRequest**](IambackofficeLoginMfaRequest.md) |  | 

### Return type

[**IambackofficeLoginResponse**](IambackofficeLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeLogout

> IambackofficeLogoutResponse IamBackofficeLogout(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeLogoutRequest() // IambackofficeLogoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeLogout(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeLogout`: IambackofficeLogoutResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeLogout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeLogoutRequest**](IambackofficeLogoutRequest.md) |  | 

### Return type

[**IambackofficeLogoutResponse**](IambackofficeLogoutResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeRegister

> IambackofficeRegistrationResponse IamBackofficeRegister(ctx).Body(body).Execute()

REGISTRATION

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeRegistrationRequest() // IambackofficeRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeRegister(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeRegister`: IambackofficeRegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeRegistrationRequest**](IambackofficeRegistrationRequest.md) |  | 

### Return type

[**IambackofficeRegistrationResponse**](IambackofficeRegistrationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeSearchGroups

> IambackofficeSearchGroupsResponse IamBackofficeSearchGroups(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeSearchGroupsRequest() // IambackofficeSearchGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeSearchGroups(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeSearchGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeSearchGroups`: IambackofficeSearchGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeSearchGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeSearchGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeSearchGroupsRequest**](IambackofficeSearchGroupsRequest.md) |  | 

### Return type

[**IambackofficeSearchGroupsResponse**](IambackofficeSearchGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamBackofficeUnassignRoles

> IambackofficeUnassignRolesResponse IamBackofficeUnassignRoles(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-iambackoffice"
)

func main() {
	body := *openapiclient.NewIambackofficeUnassignRolesRequest() // IambackofficeUnassignRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamBackofficeAPI.IamBackofficeUnassignRoles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamBackofficeAPI.IamBackofficeUnassignRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamBackofficeUnassignRoles`: IambackofficeUnassignRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `IamBackofficeAPI.IamBackofficeUnassignRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamBackofficeUnassignRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IambackofficeUnassignRolesRequest**](IambackofficeUnassignRolesRequest.md) |  | 

### Return type

[**IambackofficeUnassignRolesResponse**](IambackofficeUnassignRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

