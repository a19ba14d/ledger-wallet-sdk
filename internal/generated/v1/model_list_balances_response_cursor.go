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

// checks if the ListBalancesResponseCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBalancesResponseCursor{}

// ListBalancesResponseCursor struct for ListBalancesResponseCursor
type ListBalancesResponseCursor struct {
	PageSize int64 `json:"pageSize"`
	HasMore *bool `json:"hasMore,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
	Data []Balance `json:"data"`
}

type _ListBalancesResponseCursor ListBalancesResponseCursor

// NewListBalancesResponseCursor instantiates a new ListBalancesResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBalancesResponseCursor(pageSize int64, data []Balance) *ListBalancesResponseCursor {
	this := ListBalancesResponseCursor{}
	this.PageSize = pageSize
	this.Data = data
	return &this
}

// NewListBalancesResponseCursorWithDefaults instantiates a new ListBalancesResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBalancesResponseCursorWithDefaults() *ListBalancesResponseCursor {
	this := ListBalancesResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *ListBalancesResponseCursor) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListBalancesResponseCursor) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListBalancesResponseCursor) SetPageSize(v int64) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ListBalancesResponseCursor) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBalancesResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ListBalancesResponseCursor) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ListBalancesResponseCursor) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ListBalancesResponseCursor) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBalancesResponseCursor) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ListBalancesResponseCursor) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *ListBalancesResponseCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListBalancesResponseCursor) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBalancesResponseCursor) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListBalancesResponseCursor) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListBalancesResponseCursor) SetNext(v string) {
	o.Next = &v
}

// GetData returns the Data field value
func (o *ListBalancesResponseCursor) GetData() []Balance {
	if o == nil {
		var ret []Balance
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListBalancesResponseCursor) GetDataOk() ([]Balance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListBalancesResponseCursor) SetData(v []Balance) {
	o.Data = v
}

func (o ListBalancesResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBalancesResponseCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pageSize"] = o.PageSize
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ListBalancesResponseCursor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pageSize",
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

	varListBalancesResponseCursor := _ListBalancesResponseCursor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListBalancesResponseCursor)

	if err != nil {
		return err
	}

	*o = ListBalancesResponseCursor(varListBalancesResponseCursor)

	return err
}

type NullableListBalancesResponseCursor struct {
	value *ListBalancesResponseCursor
	isSet bool
}

func (v NullableListBalancesResponseCursor) Get() *ListBalancesResponseCursor {
	return v.value
}

func (v *NullableListBalancesResponseCursor) Set(val *ListBalancesResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableListBalancesResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableListBalancesResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBalancesResponseCursor(val *ListBalancesResponseCursor) *NullableListBalancesResponseCursor {
	return &NullableListBalancesResponseCursor{value: val, isSet: true}
}

func (v NullableListBalancesResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBalancesResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


