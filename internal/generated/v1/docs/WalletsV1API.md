# \WalletsV1API

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmHold**](WalletsV1API.md#ConfirmHold) | **Post** /holds/{hold_id}/confirm | Confirm a hold
[**CreateBalance**](WalletsV1API.md#CreateBalance) | **Post** /wallets/{id}/balances | Create a balance
[**CreateWallet**](WalletsV1API.md#CreateWallet) | **Post** /wallets | Create a new wallet
[**CreditWallet**](WalletsV1API.md#CreditWallet) | **Post** /wallets/{id}/credit | Credit a wallet
[**DebitWallet**](WalletsV1API.md#DebitWallet) | **Post** /wallets/{id}/debit | Debit a wallet
[**GetBalance**](WalletsV1API.md#GetBalance) | **Get** /wallets/{id}/balances/{balanceName} | Get detailed balance
[**GetHold**](WalletsV1API.md#GetHold) | **Get** /holds/{holdID} | Get a hold
[**GetHolds**](WalletsV1API.md#GetHolds) | **Get** /holds | Get all holds for a wallet
[**GetServerInfo**](WalletsV1API.md#GetServerInfo) | **Get** /_info | Get server info
[**GetTransactions**](WalletsV1API.md#GetTransactions) | **Get** /transactions | 
[**GetWallet**](WalletsV1API.md#GetWallet) | **Get** /wallets/{id} | Get a wallet
[**GetWalletSummary**](WalletsV1API.md#GetWalletSummary) | **Get** /wallets/{id}/summary | Get wallet summary
[**ListBalances**](WalletsV1API.md#ListBalances) | **Get** /wallets/{id}/balances | List balances of a wallet
[**ListWallets**](WalletsV1API.md#ListWallets) | **Get** /wallets | List all wallets
[**UpdateWallet**](WalletsV1API.md#UpdateWallet) | **Patch** /wallets/{id} | Update a wallet
[**VoidHold**](WalletsV1API.md#VoidHold) | **Post** /holds/{hold_id}/void | Cancel a hold



## ConfirmHold

> ConfirmHold(ctx, holdId).IdempotencyKey(idempotencyKey).ConfirmHoldRequest(confirmHoldRequest).Execute()

Confirm a hold

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	holdId := "holdId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Use an idempotency key (optional)
	confirmHoldRequest := *openapiclient.NewConfirmHoldRequest() // ConfirmHoldRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsV1API.ConfirmHold(context.Background(), holdId).IdempotencyKey(idempotencyKey).ConfirmHoldRequest(confirmHoldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.ConfirmHold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Use an idempotency key | 
 **confirmHoldRequest** | [**ConfirmHoldRequest**](ConfirmHoldRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBalance

> CreateBalanceResponse CreateBalance(ctx, id).IdempotencyKey(idempotencyKey).Body(body).Execute()

Create a balance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Use an idempotency key (optional)
	body := Balance(987) // Balance |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.CreateBalance(context.Background(), id).IdempotencyKey(idempotencyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.CreateBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBalance`: CreateBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.CreateBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Use an idempotency key | 
 **body** | **Balance** |  | 

### Return type

[**CreateBalanceResponse**](CreateBalanceResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWallet

> CreateWalletResponse CreateWallet(ctx).IdempotencyKey(idempotencyKey).CreateWalletRequest(createWalletRequest).Execute()

Create a new wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	idempotencyKey := "idempotencyKey_example" // string | Use an idempotency key (optional)
	createWalletRequest := *openapiclient.NewCreateWalletRequest(map[string]string{"key": "Inner_example"}, "Name_example") // CreateWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.CreateWallet(context.Background()).IdempotencyKey(idempotencyKey).CreateWalletRequest(createWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.CreateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWallet`: CreateWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.CreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Use an idempotency key | 
 **createWalletRequest** | [**CreateWalletRequest**](CreateWalletRequest.md) |  | 

### Return type

[**CreateWalletResponse**](CreateWalletResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditWallet

> CreditWallet(ctx, id).IdempotencyKey(idempotencyKey).CreditWalletRequest(creditWalletRequest).Execute()

Credit a wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Use an idempotency key (optional)
	creditWalletRequest := *openapiclient.NewCreditWalletRequest(*openapiclient.NewMonetary("Asset_example", int64(123))) // CreditWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsV1API.CreditWallet(context.Background(), id).IdempotencyKey(idempotencyKey).CreditWalletRequest(creditWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.CreditWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Use an idempotency key | 
 **creditWalletRequest** | [**CreditWalletRequest**](CreditWalletRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DebitWallet

> DebitWalletResponse DebitWallet(ctx, id).IdempotencyKey(idempotencyKey).DebitWalletRequest(debitWalletRequest).Execute()

Debit a wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Use an idempotency key (optional)
	debitWalletRequest := *openapiclient.NewDebitWalletRequest(*openapiclient.NewMonetary("Asset_example", int64(123)), map[string]string{"key": "Inner_example"}) // DebitWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.DebitWallet(context.Background(), id).IdempotencyKey(idempotencyKey).DebitWalletRequest(debitWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.DebitWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DebitWallet`: DebitWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.DebitWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDebitWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Use an idempotency key | 
 **debitWalletRequest** | [**DebitWalletRequest**](DebitWalletRequest.md) |  | 

### Return type

[**DebitWalletResponse**](DebitWalletResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalance

> GetBalanceResponse GetBalance(ctx, id, balanceName).Execute()

Get detailed balance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	balanceName := "balanceName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.GetBalance(context.Background(), id, balanceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.GetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalance`: GetBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.GetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**balanceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBalanceResponse**](GetBalanceResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHold

> GetHoldResponse GetHold(ctx, holdID).Execute()

Get a hold

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	holdID := "holdID_example" // string | The hold ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.GetHold(context.Background(), holdID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.GetHold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHold`: GetHoldResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.GetHold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holdID** | **string** | The hold ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHoldResponse**](GetHoldResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHolds

> GetHoldsResponse GetHolds(ctx).PageSize(pageSize).WalletID(walletID).Metadata(metadata).Cursor(cursor).Execute()

Get all holds for a wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageSize := int64(100) // int64 | The maximum number of results to return per page (optional) (default to 15)
	walletID := "wallet1" // string | The wallet to filter on (optional)
	metadata := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | Filter holds by metadata key value pairs. Nested objects can be used as seen in the example below. (optional)
	cursor := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.GetHolds(context.Background()).PageSize(pageSize).WalletID(walletID).Metadata(metadata).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.GetHolds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHolds`: GetHoldsResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.GetHolds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The maximum number of results to return per page | [default to 15]
 **walletID** | **string** | The wallet to filter on | 
 **metadata** | **map[string]map[string]string** | Filter holds by metadata key value pairs. Nested objects can be used as seen in the example below. | 
 **cursor** | **string** | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set.  | 

### Return type

[**GetHoldsResponse**](GetHoldsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInfo

> ServerInfo GetServerInfo(ctx).Execute()

Get server info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.GetServerInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.GetServerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInfo`: ServerInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.GetServerInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInfoRequest struct via the builder pattern


### Return type

[**ServerInfo**](ServerInfo.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> GetTransactionsResponse GetTransactions(ctx).PageSize(pageSize).WalletID(walletID).Cursor(cursor).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageSize := int64(100) // int64 | The maximum number of results to return per page (optional) (default to 15)
	walletID := "wallet1" // string | A wallet ID to filter on (optional)
	cursor := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the cursor is set.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.GetTransactions(context.Background()).PageSize(pageSize).WalletID(walletID).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.GetTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactions`: GetTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.GetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The maximum number of results to return per page | [default to 15]
 **walletID** | **string** | A wallet ID to filter on | 
 **cursor** | **string** | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the cursor is set.  | 

### Return type

[**GetTransactionsResponse**](GetTransactionsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallet

> GetWalletResponse GetWallet(ctx, id).Execute()

Get a wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.GetWallet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.GetWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallet`: GetWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.GetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWalletResponse**](GetWalletResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletSummary

> GetWalletSummaryResponse GetWalletSummary(ctx, id).Execute()

Get wallet summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.GetWalletSummary(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.GetWalletSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletSummary`: GetWalletSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.GetWalletSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWalletSummaryResponse**](GetWalletSummaryResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBalances

> ListBalancesResponse ListBalances(ctx, id).Execute()

List balances of a wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.ListBalances(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.ListBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBalances`: ListBalancesResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.ListBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBalancesResponse**](ListBalancesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWallets

> ListWalletsResponse ListWallets(ctx).Name(name).Metadata(metadata).PageSize(pageSize).Cursor(cursor).Expand(expand).Execute()

List all wallets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "wallet1" // string | Filter on wallet name (optional)
	metadata := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | Filter wallets by metadata key value pairs. Nested objects can be used as seen in the example below. (optional)
	pageSize := int64(100) // int64 | The maximum number of results to return per page (optional) (default to 15)
	cursor := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set.  (optional)
	expand := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsV1API.ListWallets(context.Background()).Name(name).Metadata(metadata).PageSize(pageSize).Cursor(cursor).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.ListWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWallets`: ListWalletsResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsV1API.ListWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter on wallet name | 
 **metadata** | **map[string]map[string]string** | Filter wallets by metadata key value pairs. Nested objects can be used as seen in the example below. | 
 **pageSize** | **int64** | The maximum number of results to return per page | [default to 15]
 **cursor** | **string** | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set.  | 
 **expand** | **[]string** |  | 

### Return type

[**ListWalletsResponse**](ListWalletsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWallet

> UpdateWallet(ctx, id).IdempotencyKey(idempotencyKey).UpdateWalletRequest(updateWalletRequest).Execute()

Update a wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Use an idempotency key (optional)
	updateWalletRequest := *openapiclient.NewUpdateWalletRequest(map[string]string{"key": "Inner_example"}) // UpdateWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsV1API.UpdateWallet(context.Background(), id).IdempotencyKey(idempotencyKey).UpdateWalletRequest(updateWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.UpdateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Use an idempotency key | 
 **updateWalletRequest** | [**UpdateWalletRequest**](UpdateWalletRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidHold

> VoidHold(ctx, holdId).IdempotencyKey(idempotencyKey).Execute()

Cancel a hold

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	holdId := "holdId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Use an idempotency key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsV1API.VoidHold(context.Background(), holdId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsV1API.VoidHold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Use an idempotency key | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

