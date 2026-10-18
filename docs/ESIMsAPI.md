# \ESIMsAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsimDelete**](ESIMsAPI.md#EsimDelete) | **Post** /esim/delete | Delete eSIM
[**EsimHistory**](ESIMsAPI.md#EsimHistory) | **Post** /esim/history | View eSIM history
[**EsimLabelUpdate**](ESIMsAPI.md#EsimLabelUpdate) | **Post** /esim/label/update | Update eSIM label
[**EsimPlans**](ESIMsAPI.md#EsimPlans) | **Post** /esim/plans | View eSIM plans per country
[**EsimPricing**](ESIMsAPI.md#EsimPricing) | **Post** /esim/pricing | View eSIM countries
[**EsimProfile**](ESIMsAPI.md#EsimProfile) | **Post** /esim/profile | View eSIM profile
[**EsimPurchase**](ESIMsAPI.md#EsimPurchase) | **Post** /esim/purchase | Purchase eSIM
[**EsimTopup**](ESIMsAPI.md#EsimTopup) | **Post** /esim/topup | Top up eSIM
[**EsimTopupPlans**](ESIMsAPI.md#EsimTopupPlans) | **Post** /esim/topup_plans | Top Up Plans



## EsimDelete

> interface{} EsimDelete(ctx).TransactionId(transactionId).Execute()

Delete eSIM

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
	transactionId := "transactionId_example" // string | transactionId retrieved from /esim/purchase

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimDelete(context.Background()).TransactionId(transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** | transactionId retrieved from /esim/purchase | 

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


## EsimHistory

> interface{} EsimHistory(ctx).Start(start).Length(length).Search(search).Execute()

View eSIM history

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
	start := int32(56) // int32 | Start page (optional)
	length := int32(56) // int32 | Amount of rows per page (optional)
	search := "search_example" // string | Alphanumeric search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimHistory(context.Background()).Start(start).Length(length).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimHistory`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start page | 
 **length** | **int32** | Amount of rows per page | 
 **search** | **string** | Alphanumeric search | 

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


## EsimLabelUpdate

> interface{} EsimLabelUpdate(ctx).TransactionId(transactionId).Label(label).Execute()

Update eSIM label



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
	transactionId := "transactionId_example" // string | transactionId retrieved from /esim/purchase
	label := "label_example" // string | The new label: letters, numbers and spaces, at most 30 characters. Empty or left out removes the label. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimLabelUpdate(context.Background()).TransactionId(transactionId).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimLabelUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimLabelUpdate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimLabelUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimLabelUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** | transactionId retrieved from /esim/purchase | 
 **label** | **string** | The new label: letters, numbers and spaces, at most 30 characters. Empty or left out removes the label. | 

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


## EsimPlans

> interface{} EsimPlans(ctx).Country(country).Execute()

View eSIM plans per country

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
	country := "country_example" // string | ISO 3166 country code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimPlans(context.Background()).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimPlans`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | ISO 3166 country code | 

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


## EsimPricing

> interface{} EsimPricing(ctx).Start(start).Length(length).Search(search).Execute()

View eSIM countries

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
	start := int32(56) // int32 | Start page (optional)
	length := int32(56) // int32 | Amount of rows per page (optional)
	search := "search_example" // string | Search query for country (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimPricing(context.Background()).Start(start).Length(length).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimPricing`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start page | 
 **length** | **int32** | Amount of rows per page | 
 **search** | **string** | Search query for country | 

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


## EsimProfile

> interface{} EsimProfile(ctx).TransactionId(transactionId).Execute()

View eSIM profile

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
	transactionId := "transactionId_example" // string | transactionId retrieved from /esim/purchase (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimProfile(context.Background()).TransactionId(transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimProfile`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** | transactionId retrieved from /esim/purchase | 

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


## EsimPurchase

> interface{} EsimPurchase(ctx).Plan(plan).Execute()

Purchase eSIM

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
	plan := int32(56) // int32 | The plan ID which is obtained from /esim/pricing

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimPurchase(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimPurchase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimPurchase`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimPurchase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **int32** | The plan ID which is obtained from /esim/pricing | 

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


## EsimTopup

> interface{} EsimTopup(ctx).TransactionId(transactionId).Plan(plan).Execute()

Top up eSIM

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
	transactionId := "transactionId_example" // string | Transaction ID retrieved from eSIM history
	plan := int32(56) // int32 | Topup plan retrieved from \\\"Top Up Plans\\\"

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimTopup(context.Background()).TransactionId(transactionId).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimTopup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimTopup`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimTopup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimTopupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** | Transaction ID retrieved from eSIM history | 
 **plan** | **int32** | Topup plan retrieved from \\\&quot;Top Up Plans\\\&quot; | 

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


## EsimTopupPlans

> interface{} EsimTopupPlans(ctx).Plan(plan).Execute()

Top Up Plans

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
	plan := int32(56) // int32 | `plan` retrieved from eSIM history

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ESIMsAPI.EsimTopupPlans(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ESIMsAPI.EsimTopupPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EsimTopupPlans`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ESIMsAPI.EsimTopupPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsimTopupPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **int32** | &#x60;plan&#x60; retrieved from eSIM history | 

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

