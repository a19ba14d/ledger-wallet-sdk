# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | Pointer to **string** |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Postings** | [**[]Posting**](Posting.md) |  | 
**Reference** | Pointer to **string** |  | [optional] 
**Metadata** | **map[string]string** | Metadata associated with the wallet. | 
**Id** | **int64** |  | 
**PreCommitVolumes** | Pointer to [**map[string]map[string]Volume**](map.md) |  | [optional] 
**PostCommitVolumes** | Pointer to [**map[string]map[string]Volume**](map.md) |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction(timestamp time.Time, postings []Posting, metadata map[string]string, id int64, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedger

`func (o *Transaction) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *Transaction) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *Transaction) SetLedger(v string)`

SetLedger sets Ledger field to given value.

### HasLedger

`func (o *Transaction) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### GetTimestamp

`func (o *Transaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Transaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Transaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetPostings

`func (o *Transaction) GetPostings() []Posting`

GetPostings returns the Postings field if non-nil, zero value otherwise.

### GetPostingsOk

`func (o *Transaction) GetPostingsOk() (*[]Posting, bool)`

GetPostingsOk returns a tuple with the Postings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostings

`func (o *Transaction) SetPostings(v []Posting)`

SetPostings sets Postings field to given value.


### GetReference

`func (o *Transaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Transaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Transaction) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Transaction) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetMetadata

`func (o *Transaction) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transaction) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transaction) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetId

`func (o *Transaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v int64)`

SetId sets Id field to given value.


### GetPreCommitVolumes

`func (o *Transaction) GetPreCommitVolumes() map[string]map[string]Volume`

GetPreCommitVolumes returns the PreCommitVolumes field if non-nil, zero value otherwise.

### GetPreCommitVolumesOk

`func (o *Transaction) GetPreCommitVolumesOk() (*map[string]map[string]Volume, bool)`

GetPreCommitVolumesOk returns a tuple with the PreCommitVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCommitVolumes

`func (o *Transaction) SetPreCommitVolumes(v map[string]map[string]Volume)`

SetPreCommitVolumes sets PreCommitVolumes field to given value.

### HasPreCommitVolumes

`func (o *Transaction) HasPreCommitVolumes() bool`

HasPreCommitVolumes returns a boolean if a field has been set.

### GetPostCommitVolumes

`func (o *Transaction) GetPostCommitVolumes() map[string]map[string]Volume`

GetPostCommitVolumes returns the PostCommitVolumes field if non-nil, zero value otherwise.

### GetPostCommitVolumesOk

`func (o *Transaction) GetPostCommitVolumesOk() (*map[string]map[string]Volume, bool)`

GetPostCommitVolumesOk returns a tuple with the PostCommitVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCommitVolumes

`func (o *Transaction) SetPostCommitVolumes(v map[string]map[string]Volume)`

SetPostCommitVolumes sets PostCommitVolumes field to given value.

### HasPostCommitVolumes

`func (o *Transaction) HasPostCommitVolumes() bool`

HasPostCommitVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


