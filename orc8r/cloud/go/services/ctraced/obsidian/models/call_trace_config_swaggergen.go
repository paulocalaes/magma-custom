// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CallTraceConfig Call Trace spec
// swagger:model call_trace_config
type CallTraceConfig struct {

	// Capture filter options for TShark as it would be typed out in shell. Only applies if trace_type is GATEWAY_CUSTOM.
	//
	CaptureFilters string `json:"capture_filters,omitempty"`

	// Display filter options for TShark as it would be typed out in shell. Only applies if trace_type is GATEWAY_CUSTOM.
	//
	DisplayFilters string `json:"display_filters,omitempty"`

	// ID of gateway to run call tracing on
	GatewayID string `json:"gateway_id,omitempty"`

	// Timeout of call trace in seconds
	Timeout uint32 `json:"timeout,omitempty"`

	// trace id
	// Required: true
	TraceID string `json:"trace_id"`

	// Trace Type:
	//  * GATEWAY - Call trace for a gateway
	//
	// Required: true
	// Enum: [GATEWAY]
	TraceType string `json:"trace_type"`
}

// Validate validates this call trace config
func (m *CallTraceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallTraceConfig) validateTraceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("trace_id", "body", string(m.TraceID)); err != nil {
		return err
	}

	return nil
}

var callTraceConfigTypeTraceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GATEWAY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callTraceConfigTypeTraceTypePropEnum = append(callTraceConfigTypeTraceTypePropEnum, v)
	}
}

const (

	// CallTraceConfigTraceTypeGATEWAY captures enum value "GATEWAY"
	CallTraceConfigTraceTypeGATEWAY string = "GATEWAY"
)

// prop value enum
func (m *CallTraceConfig) validateTraceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callTraceConfigTypeTraceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallTraceConfig) validateTraceType(formats strfmt.Registry) error {

	if err := validate.RequiredString("trace_type", "body", string(m.TraceType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTraceTypeEnum("trace_type", "body", m.TraceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallTraceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallTraceConfig) UnmarshalBinary(b []byte) error {
	var res CallTraceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
