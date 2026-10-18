# \CarrierAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CarrierPaidLookup**](CarrierAPI.md#CarrierPaidLookup) | **Post** /carrier/paid_lookup | Lookup



## CarrierPaidLookup

> interface{} CarrierPaidLookup(ctx).Phonenumber(phonenumber).Execute()

Lookup



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
	phonenumber := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CarrierAPI.CarrierPaidLookup(context.Background()).Phonenumber(phonenumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CarrierAPI.CarrierPaidLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CarrierPaidLookup`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CarrierAPI.CarrierPaidLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCarrierPaidLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phonenumber** | **int32** |  | 

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

