// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/tryVital/vital-go/core"
)

type ProvidersGetAllRequest struct {
	SourceType *string `json:"-" url:"source_type,omitempty"`
}

type ClientFacingProviderDetailed struct {
	// Name of source of information
	Name string `json:"name" url:"name"`
	// Slug for designated source
	Slug string `json:"slug" url:"slug"`
	// Description of source of information
	Description string `json:"description" url:"description"`
	// URL for source logo
	Logo               *string                `json:"logo,omitempty" url:"logo,omitempty"`
	AuthType           *SourceAuthType        `json:"auth_type,omitempty" url:"auth_type,omitempty"`
	SupportedResources []ClientFacingResource `json:"supported_resources,omitempty" url:"supported_resources,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientFacingProviderDetailed) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientFacingProviderDetailed) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientFacingProviderDetailed
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientFacingProviderDetailed(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientFacingProviderDetailed) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}
