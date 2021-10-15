// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyQosProfile policy qos profile
// swagger:model policy_qos_profile
type PolicyQosProfile struct {

	// arp
	Arp *Arp `json:"arp,omitempty"`

	// class id
	// Required: true
	ClassID QosClassID `json:"class_id"`

	// gbr
	Gbr *Gbr `json:"gbr,omitempty"`

	// id
	// Required: true
	// Min Length: 1
	ID string `json:"id"`

	// max req bw dl
	// Required: true
	MaxReqBwDl *uint32 `json:"max_req_bw_dl"`

	// max req bw ul
	// Required: true
	MaxReqBwUl *uint32 `json:"max_req_bw_ul"`
}

// Validate validates this policy qos profile
func (m *PolicyQosProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxReqBwDl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxReqBwUl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyQosProfile) validateArp(formats strfmt.Registry) error {

	if swag.IsZero(m.Arp) { // not required
		return nil
	}

	if m.Arp != nil {
		if err := m.Arp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arp")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyQosProfile) validateClassID(formats strfmt.Registry) error {

	if err := m.ClassID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("class_id")
		}
		return err
	}

	return nil
}

func (m *PolicyQosProfile) validateGbr(formats strfmt.Registry) error {

	if swag.IsZero(m.Gbr) { // not required
		return nil
	}

	if m.Gbr != nil {
		if err := m.Gbr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gbr")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyQosProfile) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *PolicyQosProfile) validateMaxReqBwDl(formats strfmt.Registry) error {

	if err := validate.Required("max_req_bw_dl", "body", m.MaxReqBwDl); err != nil {
		return err
	}

	return nil
}

func (m *PolicyQosProfile) validateMaxReqBwUl(formats strfmt.Registry) error {

	if err := validate.Required("max_req_bw_ul", "body", m.MaxReqBwUl); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyQosProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyQosProfile) UnmarshalBinary(b []byte) error {
	var res PolicyQosProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
