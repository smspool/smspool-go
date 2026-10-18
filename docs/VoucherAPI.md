# \VoucherAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VoucherBulkGenerate**](VoucherAPI.md#VoucherBulkGenerate) | **Post** /voucher/bulk_generate | Bulk generate vouchers
[**VoucherDelete**](VoucherAPI.md#VoucherDelete) | **Post** /voucher/delete | Delete voucher
[**VoucherGenerate**](VoucherAPI.md#VoucherGenerate) | **Post** /voucher/generate | Generate voucher
[**VoucherRetrieve**](VoucherAPI.md#VoucherRetrieve) | **Post** /voucher/retrieve | Retrieve vouchers



## VoucherBulkGenerate

> interface{} VoucherBulkGenerate(ctx).Amount(amount).Quantity(quantity).Execute()

Bulk generate vouchers

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
	amount := float32(8.14) // float32 | 
	quantity := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.VoucherBulkGenerate(context.Background()).Amount(amount).Quantity(quantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.VoucherBulkGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoucherBulkGenerate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.VoucherBulkGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoucherBulkGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** |  | 
 **quantity** | **int32** |  | 

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


## VoucherDelete

> interface{} VoucherDelete(ctx).Voucher(voucher).Execute()

Delete voucher

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
	voucher := "voucher_example" // string | Your voucher code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.VoucherDelete(context.Background()).Voucher(voucher).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.VoucherDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoucherDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.VoucherDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoucherDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voucher** | **string** | Your voucher code | 

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


## VoucherGenerate

> interface{} VoucherGenerate(ctx).Amount(amount).Execute()

Generate voucher

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
	amount := float32(8.14) // float32 | The amount per promocode

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.VoucherGenerate(context.Background()).Amount(amount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.VoucherGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoucherGenerate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.VoucherGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoucherGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** | The amount per promocode | 

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


## VoucherRetrieve

> interface{} VoucherRetrieve(ctx).Execute()

Retrieve vouchers

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
	resp, r, err := apiClient.VoucherAPI.VoucherRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.VoucherRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoucherRetrieve`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.VoucherRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVoucherRetrieveRequest struct via the builder pattern


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

