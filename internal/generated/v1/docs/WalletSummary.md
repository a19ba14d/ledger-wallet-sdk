# WalletSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | [**[]BalanceWithAssets**](BalanceWithAssets.md) |  | 
**AvailableFunds** | **map[string]int64** |  | 
**ExpiredFunds** | **map[string]int64** |  | 
**ExpirableFunds** | **map[string]int64** |  | 
**HoldFunds** | **map[string]int64** |  | 

## Methods

### NewWalletSummary

`func NewWalletSummary(balances []BalanceWithAssets, availableFunds map[string]int64, expiredFunds map[string]int64, expirableFunds map[string]int64, holdFunds map[string]int64, ) *WalletSummary`

NewWalletSummary instantiates a new WalletSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletSummaryWithDefaults

`func NewWalletSummaryWithDefaults() *WalletSummary`

NewWalletSummaryWithDefaults instantiates a new WalletSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *WalletSummary) GetBalances() []BalanceWithAssets`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *WalletSummary) GetBalancesOk() (*[]BalanceWithAssets, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *WalletSummary) SetBalances(v []BalanceWithAssets)`

SetBalances sets Balances field to given value.


### GetAvailableFunds

`func (o *WalletSummary) GetAvailableFunds() map[string]int64`

GetAvailableFunds returns the AvailableFunds field if non-nil, zero value otherwise.

### GetAvailableFundsOk

`func (o *WalletSummary) GetAvailableFundsOk() (*map[string]int64, bool)`

GetAvailableFundsOk returns a tuple with the AvailableFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFunds

`func (o *WalletSummary) SetAvailableFunds(v map[string]int64)`

SetAvailableFunds sets AvailableFunds field to given value.


### GetExpiredFunds

`func (o *WalletSummary) GetExpiredFunds() map[string]int64`

GetExpiredFunds returns the ExpiredFunds field if non-nil, zero value otherwise.

### GetExpiredFundsOk

`func (o *WalletSummary) GetExpiredFundsOk() (*map[string]int64, bool)`

GetExpiredFundsOk returns a tuple with the ExpiredFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredFunds

`func (o *WalletSummary) SetExpiredFunds(v map[string]int64)`

SetExpiredFunds sets ExpiredFunds field to given value.


### GetExpirableFunds

`func (o *WalletSummary) GetExpirableFunds() map[string]int64`

GetExpirableFunds returns the ExpirableFunds field if non-nil, zero value otherwise.

### GetExpirableFundsOk

`func (o *WalletSummary) GetExpirableFundsOk() (*map[string]int64, bool)`

GetExpirableFundsOk returns a tuple with the ExpirableFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirableFunds

`func (o *WalletSummary) SetExpirableFunds(v map[string]int64)`

SetExpirableFunds sets ExpirableFunds field to given value.


### GetHoldFunds

`func (o *WalletSummary) GetHoldFunds() map[string]int64`

GetHoldFunds returns the HoldFunds field if non-nil, zero value otherwise.

### GetHoldFundsOk

`func (o *WalletSummary) GetHoldFundsOk() (*map[string]int64, bool)`

GetHoldFundsOk returns a tuple with the HoldFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldFunds

`func (o *WalletSummary) SetHoldFunds(v map[string]int64)`

SetHoldFunds sets HoldFunds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


