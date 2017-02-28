package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// VersionInfo version info
// swagger:model VersionInfo
type VersionInfo struct {

	// id
	ID string `json:"id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// update at
	UpdateAt strfmt.DateTime `json:"updateAt,omitempty"`
}

// Validate validates this version info
func (m *VersionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
