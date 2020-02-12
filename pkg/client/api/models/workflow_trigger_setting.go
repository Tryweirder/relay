// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkflowTriggerSetting A setting for a workflow trigger
// swagger:model WorkflowTriggerSetting
type WorkflowTriggerSetting struct {

	// The value for the setting
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this workflow trigger setting
func (m *WorkflowTriggerSetting) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowTriggerSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowTriggerSetting) UnmarshalBinary(b []byte) error {
	var res WorkflowTriggerSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}