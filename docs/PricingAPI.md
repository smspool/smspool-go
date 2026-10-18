# \PricingAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestPrice**](PricingAPI.md#RequestPrice) | **Post** /request/price | Retrieve specific price &amp; Success Rate
[**RequestPricing**](PricingAPI.md#RequestPricing) | **Post** /request/pricing | Retrieve all pricing



## RequestPrice

> interface{} RequestPrice(ctx).Country(country).Service(service).Pool(pool).SkipSuccessRate(skipSuccessRate).Execute()

Retrieve specific price & Success Rate

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
	country := "country_example" // string | Country retrieved from ''Country list\\\" endpoint
	service := "service_example" // string | Service retrieved from ''Service list\\\" endpoint
	pool := "pool_example" // string | Pool retrieved from ''Pool list\\\" endpoint (optional)
	skipSuccessRate := "skipSuccessRate_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.RequestPrice(context.Background()).Country(country).Service(service).Pool(pool).SkipSuccessRate(skipSuccessRate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.RequestPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestPrice`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.RequestPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Country retrieved from &#39;&#39;Country list\\\&quot; endpoint | 
 **service** | **string** | Service retrieved from &#39;&#39;Service list\\\&quot; endpoint | 
 **pool** | **string** | Pool retrieved from &#39;&#39;Pool list\\\&quot; endpoint | 
 **skipSuccessRate** | **string** |  | 

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


## RequestPricing

> interface{} RequestPricing(ctx).Country(country).Service(service).Pool(pool).MaxPrice(maxPrice).Execute()

Retrieve all pricing

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
	country := "country_example" // string | Optional parameter to request services per country (optional)
	service := "service_example" // string | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint (optional)
	pool := "pool_example" // string |  (optional)
	maxPrice := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.RequestPricing(context.Background()).Country(country).Service(service).Pool(pool).MaxPrice(maxPrice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.RequestPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestPricing`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.RequestPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Optional parameter to request services per country | 
 **service** | **string** | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint | 
 **pool** | **string** |  | 
 **maxPrice** | **float32** |  | 

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

