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

// CreateShorturlRequest ShortUrl payload
//
// Initial payload with shorten URL
// swagger:model create_shorturl_request
type CreateShorturlRequest struct {

	// Expiry time of a URL
	ExpiryDays int32 `json:"expiryDays,omitempty"`

	// original url.
	// Required: true
	OriginalURL string `json:"originalUrl"`
}

// Validate validates this create shorturl request
func (m *CreateShorturlRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOriginalURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShorturlRequest) validateOriginalURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("originalUrl", "body", string(m.OriginalURL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateShorturlRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShorturlRequest) UnmarshalBinary(b []byte) error {
	var res CreateShorturlRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
