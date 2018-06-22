// Code generated by go generate
// This file was generated by robots at 2018-03-27 01:14:08.744890705 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uint32 is an optional uint32
type Uint32 struct {
	value *uint32
}

// NewUint32 creates a optional.Uint32 from a uint32
func NewUint32(v uint32) Uint32 {
	return Uint32{&v}
}

// Set sets the uint32 value
func (u Uint32) Set(v uint32) {
	u.value = &v
}

// Get returns the uint32 value or an error if not present
func (u Uint32) Get() (uint32, error) {
	if !u.Present() {
		return *u.value, errors.New("value not present")
	}
	return *u.value, nil
}

// Present returns whether or not the value is present
func (u Uint32) Present() bool {
	return u.value != nil
}

// OrElse returns the uint32 value or a default value if the value is not present
func (u Uint32) OrElse(v uint32) uint32 {
	if u.Present() {
		return *u.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (u Uint32) If(fn func(uint32)) {
	if u.Present() {
		fn(*u.value)
	}
}

func (u Uint32) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.value)
	}
	var zero uint32
	return json.Marshal(zero)
}

func (u *Uint32) UnmarshalJSON(data []byte) error {
	var value uint32

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.value = &value
	return nil
}
