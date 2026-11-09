# \RentalAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurchaseRental**](RentalAPI.md#PurchaseRental) | **Post** /purchase/rental | Order Rental
[**RentalAutoExtend**](RentalAPI.md#RentalAutoExtend) | **Post** /rental/auto_extend | Enable Auto Extend Rental
[**RentalExtend**](RentalAPI.md#RentalExtend) | **Post** /rental/extend | Extend Rental
[**RentalHistory**](RentalAPI.md#RentalHistory) | **Post** /rental/history | Rental History
[**RentalInfo**](RentalAPI.md#RentalInfo) | **Post** /rental/info | Retrieve Rental Info
[**RentalRefund**](RentalAPI.md#RentalRefund) | **Post** /rental/refund | Refund Rental
[**RentalReset**](RentalAPI.md#RentalReset) | **Post** /rental/reset | Reset Rental
[**RentalRetrieve**](RentalAPI.md#RentalRetrieve) | **Post** /rental/retrieve | Retrieve Active Rentals
[**RentalRetrieveAll**](RentalAPI.md#RentalRetrieveAll) | **Post** /rental/retrieve_all | Retrieve Rental IDs
[**RentalRetrieveHistory**](RentalAPI.md#RentalRetrieveHistory) | **Post** /rental/retrieve_history | Retrieve your rentals
[**RentalRetrieveMessages**](RentalAPI.md#RentalRetrieveMessages) | **Post** /rental/retrieve_messages | Retrieve Rental Messages
[**RentalRetrievePricing**](RentalAPI.md#RentalRetrievePricing) | **Post** /rental/retrieve_pricing | Retrieve Rental Pricing
[**RentalRetrieveServices**](RentalAPI.md#RentalRetrieveServices) | **Post** /rental/retrieve_services | Retrieve Rental Services
[**RentalRetrieveStatus**](RentalAPI.md#RentalRetrieveStatus) | **Post** /rental/retrieve_status | Retrieve Rental Status
[**RentalServicePrice**](RentalAPI.md#RentalServicePrice) | **Post** /rental/service_price | Retrieve rental service price
[**RentalStock**](RentalAPI.md#RentalStock) | **Post** /rental/stock | Retrieve Rental Stock



## PurchaseRental

> interface{} PurchaseRental(ctx).Id(id).Days(days).ServiceId(serviceId).CreateToken(createToken).City(city).State(state).Zipcode(zipcode).Execute()

Order Rental

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
	id := int32(56) // int32 | Rental ID retrieved from Retrieve Rental IDs
	days := int32(56) // int32 | 
	serviceId := int32(56) // int32 | Optional value on which service you'd purchase the rental for (optional)
	createToken := int32(56) // int32 | Create a rental token which allows you to share this specific rental (optional)
	city := "city_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	zipcode := "zipcode_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.PurchaseRental(context.Background()).Id(id).Days(days).ServiceId(serviceId).CreateToken(createToken).City(city).State(state).Zipcode(zipcode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.PurchaseRental``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseRental`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.PurchaseRental`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseRentalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | Rental ID retrieved from Retrieve Rental IDs | 
 **days** | **int32** |  | 
 **serviceId** | **int32** | Optional value on which service you&#39;d purchase the rental for | 
 **createToken** | **int32** | Create a rental token which allows you to share this specific rental | 
 **city** | **string** |  | 
 **state** | **string** |  | 
 **zipcode** | **string** |  | 

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


## RentalAutoExtend

> interface{} RentalAutoExtend(ctx).RentalCode(rentalCode).Execute()

Enable Auto Extend Rental

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
	rentalCode := "rentalCode_example" // string | The retrieved rental code from the `Order rental` endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalAutoExtend(context.Background()).RentalCode(rentalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalAutoExtend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalAutoExtend`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalAutoExtend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalAutoExtendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rentalCode** | **string** | The retrieved rental code from the &#x60;Order rental&#x60; endpoint. | 

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


## RentalExtend

> interface{} RentalExtend(ctx).Days(days).RentalCode(rentalCode).Execute()

Extend Rental

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
	days := int32(56) // int32 | 
	rentalCode := "rentalCode_example" // string | The retrieved rental code from the `Order rental` endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalExtend(context.Background()).Days(days).RentalCode(rentalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalExtend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalExtend`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalExtend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalExtendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** |  | 
 **rentalCode** | **string** | The retrieved rental code from the &#x60;Order rental&#x60; endpoint. | 

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


## RentalHistory

> interface{} RentalHistory(ctx).Execute()

Rental History



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
	resp, r, err := apiClient.RentalAPI.RentalHistory(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalHistory`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRentalHistoryRequest struct via the builder pattern


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


## RentalInfo

> interface{} RentalInfo(ctx).RentalCode(rentalCode).Execute()

Retrieve Rental Info



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
	rentalCode := "rentalCode_example" // string | The retrieved rental code from the `Order rental` endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalInfo(context.Background()).RentalCode(rentalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalInfo`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rentalCode** | **string** | The retrieved rental code from the &#x60;Order rental&#x60; endpoint. | 

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


## RentalRefund

> interface{} RentalRefund(ctx).RentalCode(rentalCode).Execute()

Refund Rental

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
	rentalCode := "rentalCode_example" // string | The retrieved rental code from the `Order rental` endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalRefund(context.Background()).RentalCode(rentalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRefund`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rentalCode** | **string** | The retrieved rental code from the &#x60;Order rental&#x60; endpoint. | 

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


## RentalReset

> interface{} RentalReset(ctx).RentalCode(rentalCode).Execute()

Reset Rental

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
	rentalCode := "rentalCode_example" // string | The retrieved rental code from the `Order rental` endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalReset(context.Background()).RentalCode(rentalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalReset`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalReset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rentalCode** | **string** | The retrieved rental code from the &#x60;Order rental&#x60; endpoint. | 

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


## RentalRetrieve

> interface{} RentalRetrieve(ctx).Execute()

Retrieve Active Rentals



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
	resp, r, err := apiClient.RentalAPI.RentalRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRetrieve`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRentalRetrieveRequest struct via the builder pattern


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


## RentalRetrieveAll

> interface{} RentalRetrieveAll(ctx).Type_(type_).Execute()

Retrieve Rental IDs



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
	type_ := int32(56) // int32 | Choose whether the rental is extendable or not with a 0 or 1

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalRetrieveAll(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRetrieveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRetrieveAll`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRetrieveAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalRetrieveAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **int32** | Choose whether the rental is extendable or not with a 0 or 1 | 

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


## RentalRetrieveHistory

> interface{} RentalRetrieveHistory(ctx).Execute()

Retrieve your rentals



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
	resp, r, err := apiClient.RentalAPI.RentalRetrieveHistory(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRetrieveHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRetrieveHistory`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRetrieveHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRentalRetrieveHistoryRequest struct via the builder pattern


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


## RentalRetrieveMessages

> interface{} RentalRetrieveMessages(ctx).RentalCode(rentalCode).Execute()

Retrieve Rental Messages

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
	rentalCode := "rentalCode_example" // string | The retrieved rental code from the `Order rental` endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalRetrieveMessages(context.Background()).RentalCode(rentalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRetrieveMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRetrieveMessages`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRetrieveMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalRetrieveMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rentalCode** | **string** | The retrieved rental code from the &#x60;Order rental&#x60; endpoint. | 

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


## RentalRetrievePricing

> interface{} RentalRetrievePricing(ctx).Id(id).Service(service).Execute()

Retrieve Rental Pricing

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
	id := int32(56) // int32 | The rental ID retrieved from `Retrieve Rental IDs`
	service := "service_example" // string | Service name or service ID which can be retrieved at the /service/retrieve_all endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalRetrievePricing(context.Background()).Id(id).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRetrievePricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRetrievePricing`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRetrievePricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalRetrievePricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | The rental ID retrieved from &#x60;Retrieve Rental IDs&#x60; | 
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


## RentalRetrieveServices

> interface{} RentalRetrieveServices(ctx).Rental(rental).Execute()

Retrieve Rental Services

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
	rental := int32(56) // int32 | The rental ID retrieved from `Retrieve Rental IDs` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalRetrieveServices(context.Background()).Rental(rental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRetrieveServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRetrieveServices`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRetrieveServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalRetrieveServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rental** | **int32** | The rental ID retrieved from &#x60;Retrieve Rental IDs&#x60; | 

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


## RentalRetrieveStatus

> interface{} RentalRetrieveStatus(ctx).RentalCode(rentalCode).Execute()

Retrieve Rental Status

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
	rentalCode := "rentalCode_example" // string | The retrieved rental code from the `Order rental` endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalRetrieveStatus(context.Background()).RentalCode(rentalCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalRetrieveStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalRetrieveStatus`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalRetrieveStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalRetrieveStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rentalCode** | **string** | The retrieved rental code from the &#x60;Order rental&#x60; endpoint. | 

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


## RentalServicePrice

> interface{} RentalServicePrice(ctx).Id(id).ServiceId(serviceId).Execute()

Retrieve rental service price



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
	id := int32(56) // int32 | The rental ID retrieved from `Retrieve Rental IDs`
	serviceId := int32(56) // int32 | The service ID retrieved from `Retrieve Rental Services`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalServicePrice(context.Background()).Id(id).ServiceId(serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalServicePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalServicePrice`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalServicePrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalServicePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | The rental ID retrieved from &#x60;Retrieve Rental IDs&#x60; | 
 **serviceId** | **int32** | The service ID retrieved from &#x60;Retrieve Rental Services&#x60; | 

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


## RentalStock

> interface{} RentalStock(ctx).Id(id).Days(days).ServiceId(serviceId).Execute()

Retrieve Rental Stock

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
	id := int32(56) // int32 | The rental ID retrieved from `Retrieve Rental IDs`
	days := int32(56) // int32 | 
	serviceId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RentalAPI.RentalStock(context.Background()).Id(id).Days(days).ServiceId(serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RentalAPI.RentalStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RentalStock`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RentalAPI.RentalStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRentalStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | The rental ID retrieved from &#x60;Retrieve Rental IDs&#x60; | 
 **days** | **int32** |  | 
 **serviceId** | **int32** |  | 

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

