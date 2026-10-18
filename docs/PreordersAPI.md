# \PreordersAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreorderCancel**](PreordersAPI.md#PreorderCancel) | **Post** /preorder/cancel | Cancel
[**PreorderCheck**](PreordersAPI.md#PreorderCheck) | **Post** /preorder/check | Check
[**PreorderCreate**](PreordersAPI.md#PreorderCreate) | **Post** /preorder/create | Create
[**PreorderPrice**](PreordersAPI.md#PreorderPrice) | **Post** /preorder/price | Price
[**PreorderRetrieve**](PreordersAPI.md#PreorderRetrieve) | **Post** /preorder/retrieve | Retrieve



## PreorderCancel

> interface{} PreorderCancel(ctx).Orderid(orderid).Execute()

Cancel

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
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreordersAPI.PreorderCancel(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreordersAPI.PreorderCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreorderCancel`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PreordersAPI.PreorderCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreorderCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderid** | **string** |  | 

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


## PreorderCheck

> interface{} PreorderCheck(ctx).Orderid(orderid).Execute()

Check

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
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreordersAPI.PreorderCheck(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreordersAPI.PreorderCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreorderCheck`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PreordersAPI.PreorderCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreorderCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderid** | **string** |  | 

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


## PreorderCreate

> interface{} PreorderCreate(ctx).Service(service).Country(country).Pool(pool).MaxPrice(maxPrice).Areacode(areacode).Execute()

Create

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
	service := "service_example" // string | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint
	country := "country_example" // string | Country name or country ID which can be retrieved at the /country/retrieve_all endpoint
	pool := "pool_example" // string | 
	maxPrice := float32(8.14) // float32 | 
	areacode := "areacode_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreordersAPI.PreorderCreate(context.Background()).Service(service).Country(country).Pool(pool).MaxPrice(maxPrice).Areacode(areacode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreordersAPI.PreorderCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreorderCreate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PreordersAPI.PreorderCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreorderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint | 
 **country** | **string** | Country name or country ID which can be retrieved at the /country/retrieve_all endpoint | 
 **pool** | **string** |  | 
 **maxPrice** | **float32** |  | 
 **areacode** | **string** |  | 

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


## PreorderPrice

> interface{} PreorderPrice(ctx).Service(service).Country(country).Pool(pool).Execute()

Price

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
	service := "service_example" // string | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint
	country := "country_example" // string | Country name or country ID which can be retrieved at the /country/retrieve_all endpoint
	pool := "pool_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreordersAPI.PreorderPrice(context.Background()).Service(service).Country(country).Pool(pool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreordersAPI.PreorderPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreorderPrice`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PreordersAPI.PreorderPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreorderPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint | 
 **country** | **string** | Country name or country ID which can be retrieved at the /country/retrieve_all endpoint | 
 **pool** | **string** |  | 

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


## PreorderRetrieve

> interface{} PreorderRetrieve(ctx).Execute()

Retrieve



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
	resp, r, err := apiClient.PreordersAPI.PreorderRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreordersAPI.PreorderRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreorderRetrieve`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PreordersAPI.PreorderRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPreorderRetrieveRequest struct via the builder pattern


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

