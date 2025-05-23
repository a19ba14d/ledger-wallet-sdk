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

// checks if the GetWalletSummaryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWalletSummaryResponse{}

// GetWalletSummaryResponse struct for GetWalletSummaryResponse
type GetWalletSummaryResponse struct {
	Data WalletSummary `json:"data"`
}

type _GetWalletSummaryResponse GetWalletSummaryResponse

// NewGetWalletSummaryResponse instantiates a new GetWalletSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletSummaryResponse(data WalletSummary) *GetWalletSummaryResponse {
	this := GetWalletSummaryResponse{}
	this.Data = data
	return &this
}

// NewGetWalletSummaryResponseWithDefaults instantiates a new GetWalletSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletSummaryResponseWithDefaults() *GetWalletSummaryResponse {
	this := GetWalletSummaryResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetWalletSummaryResponse) GetData() WalletSummary {
	if o == nil {
		var ret WalletSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetWalletSummaryResponse) GetDataOk() (*WalletSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetWalletSummaryResponse) SetData(v WalletSummary) {
	o.Data = v
}

func (o GetWalletSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWalletSummaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetWalletSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetWalletSummaryResponse := _GetWalletSummaryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetWalletSummaryResponse)

	if err != nil {
		return err
	}

	*o = GetWalletSummaryResponse(varGetWalletSummaryResponse)

	return err
}

type NullableGetWalletSummaryResponse struct {
	value *GetWalletSummaryResponse
	isSet bool
}

func (v NullableGetWalletSummaryResponse) Get() *GetWalletSummaryResponse {
	return v.value
}

func (v *NullableGetWalletSummaryResponse) Set(val *GetWalletSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletSummaryResponse(val *GetWalletSummaryResponse) *NullableGetWalletSummaryResponse {
	return &NullableGetWalletSummaryResponse{value: val, isSet: true}
}

func (v NullableGetWalletSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


