# \InformationAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountryRetrieveAll**](InformationAPI.md#CountryRetrieveAll) | **Get** /country/retrieve_all | Country list
[**PoolRetrieveAll**](InformationAPI.md#PoolRetrieveAll) | **Post** /pool/retrieve_all | Pool list
[**PoolRetrieveValid**](InformationAPI.md#PoolRetrieveValid) | **Post** /pool/retrieve_valid | Suggested Pools
[**RequestBalance**](InformationAPI.md#RequestBalance) | **Post** /request/balance | Balance
[**RequestHistorySearch**](InformationAPI.md#RequestHistorySearch) | **Post** /request/history_search | History search
[**RequestSuccessRate**](InformationAPI.md#RequestSuccessRate) | **Post** /request/success_rate | Retrieve country success rates per service
[**RequestSuggestedCountries**](InformationAPI.md#RequestSuggestedCountries) | **Post** /request/suggested_countries | Suggested Countries
[**ServiceRetrieveAll**](InformationAPI.md#ServiceRetrieveAll) | **Get** /service/retrieve_all | Service list



## CountryRetrieveAll

> interface{} CountryRetrieveAll(ctx).ShowAll(showAll).Execute()

Country list



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
	showAll := "showAll_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InformationAPI.CountryRetrieveAll(context.Background()).ShowAll(showAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.CountryRetrieveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountryRetrieveAll`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.CountryRetrieveAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountryRetrieveAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showAll** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolRetrieveAll

> interface{} PoolRetrieveAll(ctx).Execute()

Pool list



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
	resp, r, err := apiClient.InformationAPI.PoolRetrieveAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.PoolRetrieveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoolRetrieveAll`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.PoolRetrieveAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoolRetrieveAllRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolRetrieveValid

> interface{} PoolRetrieveValid(ctx).Service(service).Country(country).Web(web).IncludeStock(includeStock).Execute()

Suggested Pools



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
	service := "service_example" // string | Service retrieved from ''Service list\\\" endpoint
	country := "country_example" // string | Country retrieved from ''Country list\\\" endpoint
	web := "web_example" // string | Return array format (optional)
	includeStock := "includeStock_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InformationAPI.PoolRetrieveValid(context.Background()).Service(service).Country(country).Web(web).IncludeStock(includeStock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.PoolRetrieveValid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoolRetrieveValid`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.PoolRetrieveValid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolRetrieveValidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Service retrieved from &#39;&#39;Service list\\\&quot; endpoint | 
 **country** | **string** | Country retrieved from &#39;&#39;Country list\\\&quot; endpoint | 
 **web** | **string** | Return array format | 
 **includeStock** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestBalance

> interface{} RequestBalance(ctx).Execute()

Balance



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
	resp, r, err := apiClient.InformationAPI.RequestBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.RequestBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestBalance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.RequestBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestBalanceRequest struct via the builder pattern


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


## RequestHistorySearch

> interface{} RequestHistorySearch(ctx).Search(search).Cursor(cursor).Limit(limit).Execute()

History search



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
	search := "search_example" // string | 
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InformationAPI.RequestHistorySearch(context.Background()).Search(search).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.RequestHistorySearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestHistorySearch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.RequestHistorySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestHistorySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

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


## RequestSuccessRate

> interface{} RequestSuccessRate(ctx).Service(service).Country(country).Execute()

Retrieve country success rates per service



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
	country := "country_example" // string | Country name or country ID which can be retrieved at the /country/retrieve_all endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InformationAPI.RequestSuccessRate(context.Background()).Service(service).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.RequestSuccessRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestSuccessRate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.RequestSuccessRate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSuccessRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint | 
 **country** | **string** | Country name or country ID which can be retrieved at the /country/retrieve_all endpoint | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSuggestedCountries

> interface{} RequestSuggestedCountries(ctx).Service(service).Execute()

Suggested Countries



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InformationAPI.RequestSuggestedCountries(context.Background()).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.RequestSuggestedCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestSuggestedCountries`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.RequestSuggestedCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSuggestedCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceRetrieveAll

> interface{} ServiceRetrieveAll(ctx).Country(country).Execute()

Service list



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
	country := "1" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InformationAPI.ServiceRetrieveAll(context.Background()).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InformationAPI.ServiceRetrieveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceRetrieveAll`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InformationAPI.ServiceRetrieveAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceRetrieveAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

