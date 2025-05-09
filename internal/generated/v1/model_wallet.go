/*
Formance Simple Wallets Service API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package walletsclient

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Wallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Wallet{}

// Wallet struct for Wallet
type Wallet struct {
	// The unique ID of the wallet.
	Id string `json:"id"`
	// Metadata associated with the wallet.
	Metadata map[string]string `json:"metadata"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	Ledger string `json:"ledger"`
	Balances *WalletBalances `json:"balances,omitempty"`
}

type _Wallet Wallet

// NewWallet instantiates a new Wallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWallet(id string, metadata map[string]string, name string, createdAt time.Time, ledger string) *Wallet {
	this := Wallet{}
	this.Id = id
	this.Metadata = metadata
	this.Name = name
	this.CreatedAt = createdAt
	this.Ledger = ledger
	return &this
}

// NewWalletWithDefaults instantiates a new Wallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletWithDefaults() *Wallet {
	this := Wallet{}
	return &this
}

// GetId returns the Id field value
func (o *Wallet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Wallet) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value
func (o *Wallet) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Wallet) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *Wallet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Wallet) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Wallet) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Wallet) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetLedger returns the Ledger field value
func (o *Wallet) GetLedger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ledger
}

// GetLedgerOk returns a tuple with the Ledger field value
// and a boolean to check if the value has been set.
func (o *Wallet) GetLedgerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ledger, true
}

// SetLedger sets field value
func (o *Wallet) SetLedger(v string) {
	o.Ledger = v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *Wallet) GetBalances() WalletBalances {
	if o == nil || IsNil(o.Balances) {
		var ret WalletBalances
		return ret
	}
	return *o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wallet) GetBalancesOk() (*WalletBalances, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *Wallet) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given WalletBalances and assigns it to the Balances field.
func (o *Wallet) SetBalances(v WalletBalances) {
	o.Balances = &v
}

func (o Wallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Wallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["ledger"] = o.Ledger
	if !IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	return toSerialize, nil
}

func (o *Wallet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"metadata",
		"name",
		"createdAt",
		"ledger",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWallet := _Wallet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWallet)

	if err != nil {
		return err
	}

	*o = Wallet(varWallet)

	return err
}

type NullableWallet struct {
	value *Wallet
	isSet bool
}

func (v NullableWallet) Get() *Wallet {
	return v.value
}

func (v *NullableWallet) Set(val *Wallet) {
	v.value = val
	v.isSet = true
}

func (v NullableWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWallet(val *Wallet) *NullableWallet {
	return &NullableWallet{value: val, isSet: true}
}

func (v NullableWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


