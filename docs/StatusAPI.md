# \StatusAPI

All URIs are relative to *https://api.smspool.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Up**](StatusAPI.md#Up) | **Get** /up | Health check



## Up

> interface{} Up(ctx).Execute()

Health check



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
	resp, r, err := apiClient.StatusAPI.Up(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.Up``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Up`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.Up`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpRequest struct via the builder pattern


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

