// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NewIntegration new integration
// swagger:model NewIntegration
type NewIntegration struct {

	// name
	Name string `json:"name,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this new integration
func (m *NewIntegration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewIntegration) UnmarshalBinary(b []byte) error {
	var res NewIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}