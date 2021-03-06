// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Vote vote
//
// swagger:model Vote
type Vote struct {

	// option
	Option string `json:"option,omitempty"`

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`

	// voter
	Voter string `json:"voter,omitempty"`
}

// Validate validates this vote
func (m *Vote) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Vote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vote) UnmarshalBinary(b []byte) error {
	var res Vote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
