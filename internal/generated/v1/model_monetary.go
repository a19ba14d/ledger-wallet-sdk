/*
Formance Simple Wallets Service API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package walletsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Monetary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Monetary{}

// Monetary struct for Monetary
type Monetary struct {
	// The asset of the monetary value.
	Asset string `json:"asset"`
	// The amount of the monetary value.
	Amount int64 `json:"amount"`
}

type _Monetary Monetary

// NewMonetary instantiates a new Monetary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonetary(asset string, amount int64) *Monetary {
	this := Monetary{}
	this.Asset = asset
	this.Amount = amount
	return &this
}

// NewMonetaryWithDefaults instantiates a new Monetary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonetaryWithDefaults() *Monetary {
	this := Monetary{}
	return &this
}

// GetAsset returns the Asset field value
func (o *Monetary) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *Monetary) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *Monetary) SetAsset(v string) {
	o.Asset = v
}

// GetAmount returns the Amount field value
func (o *Monetary) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Monetary) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Monetary) SetAmount(v int64) {
	o.Amount = v
}

func (o Monetary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Monetary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset"] = o.Asset
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *Monetary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset",
		"amount",
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

	varMonetary := _Monetary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMonetary)

	if err != nil {
		return err
	}

	*o = Monetary(varMonetary)

	return err
}

type NullableMonetary struct {
	value *Monetary
	isSet bool
}

func (v NullableMonetary) Get() *Monetary {
	return v.value
}

func (v *NullableMonetary) Set(val *Monetary) {
	v.value = val
	v.isSet = true
}

func (v NullableMonetary) IsSet() bool {
	return v.isSet
}

func (v *NullableMonetary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonetary(val *Monetary) *NullableMonetary {
	return &NullableMonetary{value: val, isSet: true}
}

func (v NullableMonetary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonetary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


