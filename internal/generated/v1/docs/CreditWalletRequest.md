# CreditWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Monetary**](Monetary.md) |  | 
**Metadata** | Pointer to **map[string]string** | Metadata associated with the wallet. | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]Subject**](Subject.md) |  | [optional] 
**Balance** | Pointer to **string** | The balance to credit | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreditWalletRequest

`func NewCreditWalletRequest(amount Monetary, ) *CreditWalletRequest`

NewCreditWalletRequest instantiates a new CreditWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditWalletRequestWithDefaults

`func NewCreditWalletRequestWithDefaults() *CreditWalletRequest`

NewCreditWalletRequestWithDefaults instantiates a new CreditWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreditWalletRequest) GetAmount() Monetary`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditWalletRequest) GetAmountOk() (*Monetary, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditWalletRequest) SetAmount(v Monetary)`

SetAmount sets Amount field to given value.


### GetMetadata

`func (o *CreditWalletRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreditWalletRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreditWalletRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreditWalletRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreditWalletRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreditWalletRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetReference

`func (o *CreditWalletRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreditWalletRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreditWalletRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreditWalletRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSources

`func (o *CreditWalletRequest) GetSources() []Subject`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreditWalletRequest) GetSourcesOk() (*[]Subject, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreditWalletRequest) SetSources(v []Subject)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CreditWalletRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *CreditWalletRequest) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *CreditWalletRequest) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetBalance

`func (o *CreditWalletRequest) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CreditWalletRequest) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CreditWalletRequest) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CreditWalletRequest) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetTimestamp

`func (o *CreditWalletRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CreditWalletRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CreditWalletRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CreditWalletRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


