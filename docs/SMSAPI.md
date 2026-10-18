# \SMSAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurchaseSms**](SMSAPI.md#PurchaseSms) | **Post** /purchase/sms | Order SMS
[**RequestActive**](SMSAPI.md#RequestActive) | **Post** /request/active | Active orders
[**RequestArchive**](SMSAPI.md#RequestArchive) | **Post** /request/archive | Archive all orders
[**RequestAreacodes**](SMSAPI.md#RequestAreacodes) | **Post** /request/areacodes | Request available areacodes
[**RequestHistory**](SMSAPI.md#RequestHistory) | **Post** /request/history | Order History
[**SmsActivate**](SMSAPI.md#SmsActivate) | **Post** /sms/activate | Activate SMS
[**SmsAllStock**](SMSAPI.md#SmsAllStock) | **Post** /sms/all_stock | One-time stock for all services
[**SmsCancel**](SMSAPI.md#SmsCancel) | **Post** /sms/cancel | Cancel SMS
[**SmsCancelAll**](SMSAPI.md#SmsCancelAll) | **Post** /sms/cancel_all | Cancel All SMS
[**SmsCheck**](SMSAPI.md#SmsCheck) | **Post** /sms/check | Check SMS
[**SmsCheckResend**](SMSAPI.md#SmsCheckResend) | **Post** /sms/check_resend | Check Resend
[**SmsClearCache**](SMSAPI.md#SmsClearCache) | **Post** /sms/clear_cache | Clear SMS Cache
[**SmsMissed**](SMSAPI.md#SmsMissed) | **Post** /sms/missed | Missed SMS
[**SmsReactivate**](SMSAPI.md#SmsReactivate) | **Post** /sms/reactivate | Reactivate SMS
[**SmsResend**](SMSAPI.md#SmsResend) | **Post** /sms/resend | Resend
[**SmsStock**](SMSAPI.md#SmsStock) | **Post** /sms/stock | One-time SMS stock
[**SmsUnarchive**](SMSAPI.md#SmsUnarchive) | **Post** /sms/unarchive | Unarchive order



## PurchaseSms

> interface{} PurchaseSms(ctx).Service(service).Country(country).Pool(pool).MaxPrice(maxPrice).PricingOption(pricingOption).Quantity(quantity).Areacode(areacode).Exclude(exclude).CreateToken(createToken).ActivationType(activationType).Carrier(carrier).Phonenumber(phonenumber).Software(software).Execute()

Order SMS

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
	service := "service_example" // string | The service or number retrieved from the `Service List` endpoint
	country := "country_example" // string | The country in ISO format or number retrieved from the `Country List` endpoint (optional)
	pool := "pool_example" // string | The pool or number retrieved from the `Pool List` endpoint (optional) (optional)
	maxPrice := float32(8.14) // float32 | The max price you are willing to pay per number (optional)
	pricingOption := int32(56) // int32 | Set to 0 if you'd like the cheapest numbers, set to 1 for highest success rate (optional)
	quantity := int32(56) // int32 | Quantity of numbers ordered (optional)
	areacode := "areacode_example" // string | Areacodes you would like to include or exclude in JSON format (optional)
	exclude := int32(56) // int32 | Set exclude to 1 if you would like to exclude all listed area codes. (optional)
	createToken := int32(56) // int32 | Optional param; set to 1 if you'd like to create a token link that anyone can access. (optional)
	activationType := "activationType_example" // string | Options: SMS, VOICE, FLASH. SMS will receive regular SMS, VOICE will receive phone calls that mentions the SMS and FLASH is for FLASH calls for verification. (optional)
	carrier := int32(56) // int32 | Options: 1 or 2. Pick a carrier (only works for pool Foxtrot US) (optional)
	phonenumber := int32(56) // int32 | In case you want to select a specific phone number to order (optional)
	software := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.PurchaseSms(context.Background()).Service(service).Country(country).Pool(pool).MaxPrice(maxPrice).PricingOption(pricingOption).Quantity(quantity).Areacode(areacode).Exclude(exclude).CreateToken(createToken).ActivationType(activationType).Carrier(carrier).Phonenumber(phonenumber).Software(software).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.PurchaseSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseSms`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.PurchaseSms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | The service or number retrieved from the &#x60;Service List&#x60; endpoint | 
 **country** | **string** | The country in ISO format or number retrieved from the &#x60;Country List&#x60; endpoint | 
 **pool** | **string** | The pool or number retrieved from the &#x60;Pool List&#x60; endpoint (optional) | 
 **maxPrice** | **float32** | The max price you are willing to pay per number | 
 **pricingOption** | **int32** | Set to 0 if you&#39;d like the cheapest numbers, set to 1 for highest success rate | 
 **quantity** | **int32** | Quantity of numbers ordered | 
 **areacode** | **string** | Areacodes you would like to include or exclude in JSON format | 
 **exclude** | **int32** | Set exclude to 1 if you would like to exclude all listed area codes. | 
 **createToken** | **int32** | Optional param; set to 1 if you&#39;d like to create a token link that anyone can access. | 
 **activationType** | **string** | Options: SMS, VOICE, FLASH. SMS will receive regular SMS, VOICE will receive phone calls that mentions the SMS and FLASH is for FLASH calls for verification. | 
 **carrier** | **int32** | Options: 1 or 2. Pick a carrier (only works for pool Foxtrot US) | 
 **phonenumber** | **int32** | In case you want to select a specific phone number to order | 
 **software** | **int32** |  | 

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


## RequestActive

> interface{} RequestActive(ctx).Execute()

Active orders



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
	resp, r, err := apiClient.SMSAPI.RequestActive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.RequestActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestActive`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.RequestActive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestActiveRequest struct via the builder pattern


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


## RequestArchive

> interface{} RequestArchive(ctx).Execute()

Archive all orders



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
	resp, r, err := apiClient.SMSAPI.RequestArchive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.RequestArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestArchive`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.RequestArchive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestArchiveRequest struct via the builder pattern


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


## RequestAreacodes

> interface{} RequestAreacodes(ctx).Service(service).Country(country).Pool(pool).Execute()

Request available areacodes

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
	service := "service_example" // string | The service or number retrieved from the `Service List` endpoint
	country := "country_example" // string | The country in ISO format or number retrieved from the `Country List` endpoint
	pool := "pool_example" // string | The pool or number retrieved from the `Pool List` endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.RequestAreacodes(context.Background()).Service(service).Country(country).Pool(pool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.RequestAreacodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestAreacodes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.RequestAreacodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestAreacodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | The service or number retrieved from the &#x60;Service List&#x60; endpoint | 
 **country** | **string** | The country in ISO format or number retrieved from the &#x60;Country List&#x60; endpoint | 
 **pool** | **string** | The pool or number retrieved from the &#x60;Pool List&#x60; endpoint (optional) | 

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


## RequestHistory

> interface{} RequestHistory(ctx).Start(start).Length(length).Search(search).Execute()

Order History



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
	start := int32(56) // int32 | Which row you'd like to start from (optional)
	length := int32(56) // int32 | Max length of all rows (optional)
	search := "search_example" // string | Search query for phone number, order_code, service or country (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.RequestHistory(context.Background()).Start(start).Length(length).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.RequestHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestHistory`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.RequestHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Which row you&#39;d like to start from | 
 **length** | **int32** | Max length of all rows | 
 **search** | **string** | Search query for phone number, order_code, service or country | 

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


## SmsActivate

> interface{} SmsActivate(ctx).Orderid(orderid).Execute()

Activate SMS



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
	resp, r, err := apiClient.SMSAPI.SmsActivate(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsActivate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsActivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsActivateRequest struct via the builder pattern


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


## SmsAllStock

> interface{} SmsAllStock(ctx).Country(country).Service(service).Pool(pool).Execute()

One-time stock for all services



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
	country := "country_example" // string | Country retrieved from ''Country list\\\" endpoint (optional) (optional)
	service := "service_example" // string | Service retrieved from ''Service list\\\" endpoint (optional) (optional)
	pool := "pool_example" // string | Pool retrieved from ''Pool list\\\" endpoint (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.SmsAllStock(context.Background()).Country(country).Service(service).Pool(pool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsAllStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsAllStock`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsAllStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsAllStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Country retrieved from &#39;&#39;Country list\\\&quot; endpoint (optional) | 
 **service** | **string** | Service retrieved from &#39;&#39;Service list\\\&quot; endpoint (optional) | 
 **pool** | **string** | Pool retrieved from &#39;&#39;Pool list\\\&quot; endpoint (optional) | 

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


## SmsCancel

> interface{} SmsCancel(ctx).Orderid(orderid).Orderids(orderids).Expiry(expiry).Ban(ban).Execute()

Cancel SMS



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
	orderid := "orderid_example" // string |  (optional)
	orderids := "orderids_example" // string |  (optional)
	expiry := int32(56) // int32 |  (optional)
	ban := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.SmsCancel(context.Background()).Orderid(orderid).Orderids(orderids).Expiry(expiry).Ban(ban).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsCancel`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderid** | **string** |  | 
 **orderids** | **string** |  | 
 **expiry** | **int32** |  | 
 **ban** | **int32** |  | 

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


## SmsCancelAll

> interface{} SmsCancelAll(ctx).Execute()

Cancel All SMS



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
	resp, r, err := apiClient.SMSAPI.SmsCancelAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsCancelAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsCancelAll`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsCancelAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSmsCancelAllRequest struct via the builder pattern


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


## SmsCheck

> interface{} SmsCheck(ctx).Orderid(orderid).Execute()

Check SMS



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
	resp, r, err := apiClient.SMSAPI.SmsCheck(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsCheck`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsCheckRequest struct via the builder pattern


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


## SmsCheckResend

> interface{} SmsCheckResend(ctx).Orderid(orderid).Execute()

Check Resend



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
	resp, r, err := apiClient.SMSAPI.SmsCheckResend(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsCheckResend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsCheckResend`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsCheckResend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsCheckResendRequest struct via the builder pattern


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


## SmsClearCache

> interface{} SmsClearCache(ctx).Execute()

Clear SMS Cache



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
	resp, r, err := apiClient.SMSAPI.SmsClearCache(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsClearCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsClearCache`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsClearCache`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSmsClearCacheRequest struct via the builder pattern


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


## SmsMissed

> interface{} SmsMissed(ctx).Phonenumber(phonenumber).Limit(limit).Execute()

Missed SMS



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
	phonenumber := int32(56) // int32 | Only return messages for this number, including the country code. Leave it out for all your numbers. (optional)
	limit := int32(56) // int32 | Maximum number of messages. Defaults to 100, at most 500. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.SmsMissed(context.Background()).Phonenumber(phonenumber).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsMissed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsMissed`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsMissed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsMissedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phonenumber** | **int32** | Only return messages for this number, including the country code. Leave it out for all your numbers. | 
 **limit** | **int32** | Maximum number of messages. Defaults to 100, at most 500. | 

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


## SmsReactivate

> interface{} SmsReactivate(ctx).Orderid(orderid).Execute()

Reactivate SMS



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
	resp, r, err := apiClient.SMSAPI.SmsReactivate(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsReactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsReactivate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsReactivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsReactivateRequest struct via the builder pattern


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


## SmsResend

> interface{} SmsResend(ctx).Orderid(orderid).Execute()

Resend



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
	resp, r, err := apiClient.SMSAPI.SmsResend(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsResend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsResend`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsResend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsResendRequest struct via the builder pattern


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


## SmsStock

> interface{} SmsStock(ctx).Country(country).Service(service).Pool(pool).IncludePrice(includePrice).Execute()

One-time SMS stock



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
	includePrice := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.SmsStock(context.Background()).Country(country).Service(service).Pool(pool).IncludePrice(includePrice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsStock`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Country retrieved from &#39;&#39;Country list\\\&quot; endpoint | 
 **service** | **string** | Service retrieved from &#39;&#39;Service list\\\&quot; endpoint | 
 **pool** | **string** | Pool retrieved from &#39;&#39;Pool list\\\&quot; endpoint | 
 **includePrice** | **int32** |  | 

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


## SmsUnarchive

> interface{} SmsUnarchive(ctx).Orderid(orderid).Execute()

Unarchive order



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
	orderid := "orderid_example" // string | The order ID retrieved from /purchase/sms

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMSAPI.SmsUnarchive(context.Background()).Orderid(orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMSAPI.SmsUnarchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmsUnarchive`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SMSAPI.SmsUnarchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmsUnarchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderid** | **string** | The order ID retrieved from /purchase/sms | 

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

