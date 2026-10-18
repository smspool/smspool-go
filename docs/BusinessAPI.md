# \BusinessAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessCreate**](BusinessAPI.md#BusinessCreate) | **Post** /business/create | Register user
[**BusinessUserHistory**](BusinessAPI.md#BusinessUserHistory) | **Post** /business/user/history | History
[**BusinessUserUpdate**](BusinessAPI.md#BusinessUserUpdate) | **Post** /business/user/update | Update
[**BusinessUsers**](BusinessAPI.md#BusinessUsers) | **Get** /business/users | Retrieve users



## BusinessCreate

> interface{} BusinessCreate(ctx).Username(username).Password(password).DailySpendLimit(dailySpendLimit).Execute()

Register user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smspool/smspool-go"
)

func main() {
	username := "username_example" // string | 
	password := "password_example" // string | 
	dailySpendLimit := float32(8.14) // float32 | Optional daily spend cap in USD. Omit for no limit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessAPI.BusinessCreate(context.Background()).Username(username).Password(password).DailySpendLimit(dailySpendLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessAPI.BusinessCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessCreate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BusinessAPI.BusinessCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBusinessCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **password** | **string** |  | 
 **dailySpendLimit** | **float32** | Optional daily spend cap in USD. Omit for no limit. | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessUserHistory

> interface{} BusinessUserHistory(ctx).Id(id).Execute()

History

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smspool/smspool-go"
)

func main() {
	id := int32(56) // int32 | User ID retrievd from `Retrieve users` endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessAPI.BusinessUserHistory(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessAPI.BusinessUserHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUserHistory`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BusinessAPI.BusinessUserHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUserHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | User ID retrievd from &#x60;Retrieve users&#x60; endpoint | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessUserUpdate

> interface{} BusinessUserUpdate(ctx).Id(id).Password(password).Balance(balance).Active(active).DailySpendLimit(dailySpendLimit).Execute()

Update

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smspool/smspool-go"
)

func main() {
	id := int32(56) // int32 | User ID retrievd from `Retrieve users` endpoint
	password := "password_example" // string | Leave empty if you do not wish to update the password (optional)
	balance := float32(8.14) // float32 | Leave empty if you do not wish to update the balance, sets the balance by the requested value (optional)
	active := int32(56) // int32 | Disable/enable user account (optional) (optional)
	dailySpendLimit := float32(8.14) // float32 | Daily spend cap in USD. Leave empty to leave unchanged, send -1 to remove the cap, 0 to block all spending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessAPI.BusinessUserUpdate(context.Background()).Id(id).Password(password).Balance(balance).Active(active).DailySpendLimit(dailySpendLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessAPI.BusinessUserUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUserUpdate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BusinessAPI.BusinessUserUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | User ID retrievd from &#x60;Retrieve users&#x60; endpoint | 
 **password** | **string** | Leave empty if you do not wish to update the password | 
 **balance** | **float32** | Leave empty if you do not wish to update the balance, sets the balance by the requested value | 
 **active** | **int32** | Disable/enable user account (optional) | 
 **dailySpendLimit** | **float32** | Daily spend cap in USD. Leave empty to leave unchanged, send -1 to remove the cap, 0 to block all spending. | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessUsers

> interface{} BusinessUsers(ctx).Execute()

Retrieve users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/smspool/smspool-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessAPI.BusinessUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessAPI.BusinessUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BusinessUsers`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BusinessAPI.BusinessUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessUsersRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

