// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IpamAddArgs IPAM request args
//
// swagger:model IpamAddArgs
type IpamAddArgs struct {

	// clean gateway
	CleanGateway bool `json:"cleanGateway,omitempty"`

	// container ID
	// Required: true
	ContainerID *string `json:"containerID"`

	// default IPv4 IP pool
	DefaultIPV4IPPool []string `json:"defaultIPv4IPPool"`

	// default IPv6 IP pool
	DefaultIPV6IPPool []string `json:"defaultIPv6IPPool"`

	// if name
	// Required: true
	IfName *string `json:"ifName"`

	// net namespace
	// Required: true
	NetNamespace *string `json:"netNamespace"`

	// pod name
	// Required: true
	PodName *string `json:"podName"`

	// pod namespace
	// Required: true
	PodNamespace *string `json:"podNamespace"`
}

// Validate validates this ipam add args
func (m *IpamAddArgs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIfName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamAddArgs) validateContainerID(formats strfmt.Registry) error {

	if err := validate.Required("containerID", "body", m.ContainerID); err != nil {
		return err
	}

	return nil
}

func (m *IpamAddArgs) validateIfName(formats strfmt.Registry) error {

	if err := validate.Required("ifName", "body", m.IfName); err != nil {
		return err
	}

	return nil
}

func (m *IpamAddArgs) validateNetNamespace(formats strfmt.Registry) error {

	if err := validate.Required("netNamespace", "body", m.NetNamespace); err != nil {
		return err
	}

	return nil
}

func (m *IpamAddArgs) validatePodName(formats strfmt.Registry) error {

	if err := validate.Required("podName", "body", m.PodName); err != nil {
		return err
	}

	return nil
}

func (m *IpamAddArgs) validatePodNamespace(formats strfmt.Registry) error {

	if err := validate.Required("podNamespace", "body", m.PodNamespace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ipam add args based on context it is used
func (m *IpamAddArgs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpamAddArgs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamAddArgs) UnmarshalBinary(b []byte) error {
	var res IpamAddArgs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
