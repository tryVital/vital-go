// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/tryVital/vital-go/core"
	time "time"
)

type UserCreateBody struct {
	// A unique ID representing the end user. Typically this will be a user ID from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
	ClientUserId string `json:"client_user_id" url:"-"`
	// Fallback time zone of the user, in the form of a valid IANA tzdatabase identifier (e.g., `Europe/London` or `America/Los_Angeles`).
	// Used when pulling data from sources that are completely time zone agnostic (e.g., all time is relative to UTC clock, without any time zone attributions on data points).
	FallbackTimeZone *string `json:"fallback_time_zone,omitempty" url:"-"`
	// Fallback date of birth of the user, in YYYY-mm-dd format. Used for calculating max heartrate for providers that don not provide users' age.
	FallbackBirthDate *string `json:"fallback_birth_date,omitempty" url:"-"`
	// Starting bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
	IngestionStart *string `json:"ingestion_start,omitempty" url:"-"`
	// Ending bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
	IngestionEnd *string `json:"ingestion_end,omitempty" url:"-"`
}

type CreateInsuranceRequest struct {
	PayorCode    string                                                  `json:"payor_code" url:"-"`
	MemberId     string                                                  `json:"member_id" url:"-"`
	GroupId      *string                                                 `json:"group_id,omitempty" url:"-"`
	Relationship ResponsibleRelationship                                 `json:"relationship" url:"-"`
	Insured      *VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails `json:"insured,omitempty" url:"-"`
	Guarantor    *GuarantorDetails                                       `json:"guarantor,omitempty" url:"-"`
}

type UserGetAllRequest struct {
	Offset *int `json:"-" url:"offset,omitempty"`
	Limit  *int `json:"-" url:"limit,omitempty"`
}

type UserPatchBody struct {
	// Fallback time zone of the user, in the form of a valid IANA tzdatabase identifier (e.g., `Europe/London` or `America/Los_Angeles`).
	// Used when pulling data from sources that are completely time zone agnostic (e.g., all time is relative to UTC clock, without any time zone attributions on data points).
	FallbackTimeZone *string `json:"fallback_time_zone,omitempty" url:"-"`
	// Fallback date of birth of the user, in YYYY-mm-dd format. Used for calculating max heartrate for providers that don not provide users' age.
	FallbackBirthDate *string `json:"fallback_birth_date,omitempty" url:"-"`
	// Starting bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
	IngestionStart *string `json:"ingestion_start,omitempty" url:"-"`
	// Ending bound for user [data ingestion bounds](https://docs.tryvital.io/wearables/providers/data-ingestion-bounds).
	IngestionEnd *string `json:"ingestion_end,omitempty" url:"-"`
	// A unique ID representing the end user. Typically this will be a user ID from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
	ClientUserId *string `json:"client_user_id,omitempty" url:"-"`
}

type UserRefreshRequest struct {
	Timeout *float64 `json:"-" url:"timeout,omitempty"`
}

// ℹ️ This enum is non-exhaustive.
type Availability string

const (
	AvailabilityAvailable   Availability = "available"
	AvailabilityUnavailable Availability = "unavailable"
)

func NewAvailabilityFromString(s string) (Availability, error) {
	switch s {
	case "available":
		return AvailabilityAvailable, nil
	case "unavailable":
		return AvailabilityUnavailable, nil
	}
	var t Availability
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a Availability) Ptr() *Availability {
	return &a
}

type ClientFacingConnectionErrorDetails struct {
	// ℹ️ This enum is non-exhaustive.
	ErrorType    ClientFacingConnectionErrorDetailsErrorType `json:"error_type" url:"error_type"`
	ErrorMessage string                                      `json:"error_message" url:"error_message"`
	ErroredAt    time.Time                                   `json:"errored_at" url:"errored_at"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientFacingConnectionErrorDetails) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientFacingConnectionErrorDetails) UnmarshalJSON(data []byte) error {
	type embed ClientFacingConnectionErrorDetails
	var unmarshaler = struct {
		embed
		ErroredAt *core.DateTime `json:"errored_at"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = ClientFacingConnectionErrorDetails(unmarshaler.embed)
	c.ErroredAt = unmarshaler.ErroredAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientFacingConnectionErrorDetails) MarshalJSON() ([]byte, error) {
	type embed ClientFacingConnectionErrorDetails
	var marshaler = struct {
		embed
		ErroredAt *core.DateTime `json:"errored_at"`
	}{
		embed:     embed(*c),
		ErroredAt: core.NewDateTime(c.ErroredAt),
	}
	return json.Marshal(marshaler)
}

func (c *ClientFacingConnectionErrorDetails) String() string {
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

// ℹ️ This enum is non-exhaustive.
type ClientFacingConnectionErrorDetailsErrorType string

const (
	ClientFacingConnectionErrorDetailsErrorTypeTokenRefreshFailed        ClientFacingConnectionErrorDetailsErrorType = "token_refresh_failed"
	ClientFacingConnectionErrorDetailsErrorTypeWebhookRegistrationFailed ClientFacingConnectionErrorDetailsErrorType = "webhook_registration_failed"
	ClientFacingConnectionErrorDetailsErrorTypeUserNotFound              ClientFacingConnectionErrorDetailsErrorType = "user_not_found"
	ClientFacingConnectionErrorDetailsErrorTypeDeregisteredPerProvider   ClientFacingConnectionErrorDetailsErrorType = "deregistered_per_provider"
	ClientFacingConnectionErrorDetailsErrorTypeRequiredScopesNotGranted  ClientFacingConnectionErrorDetailsErrorType = "required_scopes_not_granted"
	ClientFacingConnectionErrorDetailsErrorTypeProviderCredentialError   ClientFacingConnectionErrorDetailsErrorType = "provider_credential_error"
	ClientFacingConnectionErrorDetailsErrorTypeProviderPasswordExpired   ClientFacingConnectionErrorDetailsErrorType = "provider_password_expired"
	ClientFacingConnectionErrorDetailsErrorTypeUnknown                   ClientFacingConnectionErrorDetailsErrorType = "unknown"
)

func NewClientFacingConnectionErrorDetailsErrorTypeFromString(s string) (ClientFacingConnectionErrorDetailsErrorType, error) {
	switch s {
	case "token_refresh_failed":
		return ClientFacingConnectionErrorDetailsErrorTypeTokenRefreshFailed, nil
	case "webhook_registration_failed":
		return ClientFacingConnectionErrorDetailsErrorTypeWebhookRegistrationFailed, nil
	case "user_not_found":
		return ClientFacingConnectionErrorDetailsErrorTypeUserNotFound, nil
	case "deregistered_per_provider":
		return ClientFacingConnectionErrorDetailsErrorTypeDeregisteredPerProvider, nil
	case "required_scopes_not_granted":
		return ClientFacingConnectionErrorDetailsErrorTypeRequiredScopesNotGranted, nil
	case "provider_credential_error":
		return ClientFacingConnectionErrorDetailsErrorTypeProviderCredentialError, nil
	case "provider_password_expired":
		return ClientFacingConnectionErrorDetailsErrorTypeProviderPasswordExpired, nil
	case "unknown":
		return ClientFacingConnectionErrorDetailsErrorTypeUnknown, nil
	}
	var t ClientFacingConnectionErrorDetailsErrorType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ClientFacingConnectionErrorDetailsErrorType) Ptr() *ClientFacingConnectionErrorDetailsErrorType {
	return &c
}

type ClientFacingInsurance struct {
	MemberId     string                                                  `json:"member_id" url:"member_id"`
	PayorCode    string                                                  `json:"payor_code" url:"payor_code"`
	Relationship ResponsibleRelationship                                 `json:"relationship" url:"relationship"`
	Insured      *VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails `json:"insured,omitempty" url:"insured,omitempty"`
	Company      *CompanyDetails                                         `json:"company,omitempty" url:"company,omitempty"`
	GroupId      *string                                                 `json:"group_id,omitempty" url:"group_id,omitempty"`
	Guarantor    *GuarantorDetails                                       `json:"guarantor,omitempty" url:"guarantor,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientFacingInsurance) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientFacingInsurance) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientFacingInsurance
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientFacingInsurance(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientFacingInsurance) String() string {
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

type ClientFacingProviderWithStatus struct {
	// Name of source of information
	Name string `json:"name" url:"name"`
	// Slug for designated source
	Slug string `json:"slug" url:"slug"`
	// URL for source logo
	Logo      string    `json:"logo" url:"logo"`
	CreatedOn time.Time `json:"created_on" url:"created_on"`
	// Status of source, either error or connected
	Status string `json:"status" url:"status"`
	// The unique identifier of the associated external data provider user.
	//
	// * OAuth Providers: User unique identifier; provider-specific formats
	// * Password Providers: Username
	// * Email Providers: Email
	// * Junction Mobile SDK Providers: `null` (not available)
	ExternalUserId *string `json:"external_user_id,omitempty" url:"external_user_id,omitempty"`
	// Details of the terminal connection error — populated only when the status is `error`.
	ErrorDetails         *ClientFacingConnectionErrorDetails `json:"error_details,omitempty" url:"error_details,omitempty"`
	ResourceAvailability map[string]*ResourceAvailability    `json:"resource_availability,omitempty" url:"resource_availability,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientFacingProviderWithStatus) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientFacingProviderWithStatus) UnmarshalJSON(data []byte) error {
	type embed ClientFacingProviderWithStatus
	var unmarshaler = struct {
		embed
		CreatedOn *core.DateTime `json:"created_on"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = ClientFacingProviderWithStatus(unmarshaler.embed)
	c.CreatedOn = unmarshaler.CreatedOn.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientFacingProviderWithStatus) MarshalJSON() ([]byte, error) {
	type embed ClientFacingProviderWithStatus
	var marshaler = struct {
		embed
		CreatedOn *core.DateTime `json:"created_on"`
	}{
		embed:     embed(*c),
		CreatedOn: core.NewDateTime(c.CreatedOn),
	}
	return json.Marshal(marshaler)
}

func (c *ClientFacingProviderWithStatus) String() string {
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

type ClientFacingUserKey struct {
	// User id returned by vital create user request. This id should be stored in your database against the user and used for all interactions with the vital api.
	UserId string `json:"user_id" url:"user_id"`
	// A unique ID representing the end user. Typically this will be a user ID from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
	ClientUserId string `json:"client_user_id" url:"client_user_id"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientFacingUserKey) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientFacingUserKey) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientFacingUserKey
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientFacingUserKey(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientFacingUserKey) String() string {
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

type CompanyDetails struct {
	Name    string   `json:"name" url:"name"`
	Address *Address `json:"address,omitempty" url:"address,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CompanyDetails) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CompanyDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler CompanyDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CompanyDetails(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CompanyDetails) String() string {
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

type MetricsResult struct {
	TeamId                          string                   `json:"team_id" url:"team_id"`
	NumberOfConnectedSources        *int                     `json:"number_of_connected_sources,omitempty" url:"number_of_connected_sources,omitempty"`
	NumberOfUsers                   *int                     `json:"number_of_users,omitempty" url:"number_of_users,omitempty"`
	NumberOfErroredConnectedSources *int                     `json:"number_of_errored_connected_sources,omitempty" url:"number_of_errored_connected_sources,omitempty"`
	NumberOfConnectedSourcesByWeek  []*TimeseriesMetricPoint `json:"number_of_connected_sources_by_week,omitempty" url:"number_of_connected_sources_by_week,omitempty"`
	NumberOfOrderedTests            *int                     `json:"number_of_ordered_tests,omitempty" url:"number_of_ordered_tests,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (m *MetricsResult) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MetricsResult) UnmarshalJSON(data []byte) error {
	type unmarshaler MetricsResult
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MetricsResult(value)

	extraProperties, err := core.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties

	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MetricsResult) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type PaginatedUsersResponse struct {
	Users  []*ClientFacingUser `json:"users,omitempty" url:"users,omitempty"`
	Total  int                 `json:"total" url:"total"`
	Offset int                 `json:"offset" url:"offset"`
	Limit  int                 `json:"limit" url:"limit"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PaginatedUsersResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedUsersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedUsersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedUsersResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedUsersResponse) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type ResourceAvailability struct {
	Status            Availability             `json:"status" url:"status"`
	ScopeRequirements *ScopeRequirementsGrants `json:"scope_requirements,omitempty" url:"scope_requirements,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResourceAvailability) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResourceAvailability) UnmarshalJSON(data []byte) error {
	type unmarshaler ResourceAvailability
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResourceAvailability(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResourceAvailability) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type ScopeRequirementsGrants struct {
	UserGranted *ScopeRequirementsStr `json:"user_granted,omitempty" url:"user_granted,omitempty"`
	UserDenied  *ScopeRequirementsStr `json:"user_denied,omitempty" url:"user_denied,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *ScopeRequirementsGrants) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *ScopeRequirementsGrants) UnmarshalJSON(data []byte) error {
	type unmarshaler ScopeRequirementsGrants
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ScopeRequirementsGrants(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ScopeRequirementsGrants) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type ScopeRequirementsStr struct {
	Required []string `json:"required,omitempty" url:"required,omitempty"`
	Optional []string `json:"optional,omitempty" url:"optional,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *ScopeRequirementsStr) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *ScopeRequirementsStr) UnmarshalJSON(data []byte) error {
	type unmarshaler ScopeRequirementsStr
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ScopeRequirementsStr(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ScopeRequirementsStr) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type TimeseriesMetricPoint struct {
	Date  time.Time `json:"date" url:"date"`
	Value float64   `json:"value" url:"value"`
	All   float64   `json:"all" url:"all"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (t *TimeseriesMetricPoint) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TimeseriesMetricPoint) UnmarshalJSON(data []byte) error {
	type embed TimeseriesMetricPoint
	var unmarshaler = struct {
		embed
		Date *core.DateTime `json:"date"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TimeseriesMetricPoint(unmarshaler.embed)
	t.Date = unmarshaler.Date.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties

	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TimeseriesMetricPoint) MarshalJSON() ([]byte, error) {
	type embed TimeseriesMetricPoint
	var marshaler = struct {
		embed
		Date *core.DateTime `json:"date"`
	}{
		embed: embed(*t),
		Date:  core.NewDateTime(t.Date),
	}
	return json.Marshal(marshaler)
}

func (t *TimeseriesMetricPoint) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type UserInfo struct {
	FirstName         string             `json:"first_name" url:"first_name"`
	LastName          string             `json:"last_name" url:"last_name"`
	Email             string             `json:"email" url:"email"`
	PhoneNumber       string             `json:"phone_number" url:"phone_number"`
	Gender            string             `json:"gender" url:"gender"`
	Dob               string             `json:"dob" url:"dob"`
	Address           *Address           `json:"address,omitempty" url:"address,omitempty"`
	MedicalProxy      *GuarantorDetails  `json:"medical_proxy,omitempty" url:"medical_proxy,omitempty"`
	Race              *Race              `json:"race,omitempty" url:"race,omitempty"`
	Ethnicity         *Ethnicity         `json:"ethnicity,omitempty" url:"ethnicity,omitempty"`
	SexualOrientation *SexualOrientation `json:"sexual_orientation,omitempty" url:"sexual_orientation,omitempty"`
	GenderIdentity    *GenderIdentity    `json:"gender_identity,omitempty" url:"gender_identity,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UserInfo) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler UserInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserInfo(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserInfo) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserRefreshSuccessResponse struct {
	// Whether operation was successful or not
	Success bool `json:"success" url:"success"`
	// A unique ID representing the end user. Typically this will be a user ID from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
	UserId            string   `json:"user_id" url:"user_id"`
	RefreshedSources  []string `json:"refreshed_sources,omitempty" url:"refreshed_sources,omitempty"`
	InProgressSources []string `json:"in_progress_sources,omitempty" url:"in_progress_sources,omitempty"`
	FailedSources     []string `json:"failed_sources,omitempty" url:"failed_sources,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UserRefreshSuccessResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserRefreshSuccessResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserRefreshSuccessResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserRefreshSuccessResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserRefreshSuccessResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserSignInTokenResponse struct {
	UserId      string `json:"user_id" url:"user_id"`
	SignInToken string `json:"sign_in_token" url:"sign_in_token"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UserSignInTokenResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserSignInTokenResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserSignInTokenResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserSignInTokenResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserSignInTokenResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserSuccessResponse struct {
	// Whether operation was successful or not
	Success bool `json:"success" url:"success"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UserSuccessResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserSuccessResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserSuccessResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserSuccessResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserSuccessResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails struct {
	FirstName   string   `json:"first_name" url:"first_name"`
	LastName    string   `json:"last_name" url:"last_name"`
	Gender      Gender   `json:"gender" url:"gender"`
	Address     *Address `json:"address,omitempty" url:"address,omitempty"`
	Dob         string   `json:"dob" url:"dob"`
	Email       string   `json:"email" url:"email"`
	PhoneNumber string   `json:"phone_number" url:"phone_number"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VitalCoreSchemasDbSchemasLabTestInsurancePersonDetails) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type UserUndoDeleteRequest struct {
	// User ID to undo deletion. Mutually exclusive with `client_user_id`.
	UserId *string `json:"-" url:"user_id,omitempty"`
	// Client User ID to undo deletion. Mutually exclusive with `user_id`.
	ClientUserId *string `json:"-" url:"client_user_id,omitempty"`
}

type UserInfoCreateRequest struct {
	FirstName         string             `json:"first_name" url:"-"`
	LastName          string             `json:"last_name" url:"-"`
	Email             string             `json:"email" url:"-"`
	PhoneNumber       string             `json:"phone_number" url:"-"`
	Gender            string             `json:"gender" url:"-"`
	Dob               string             `json:"dob" url:"-"`
	Address           *Address           `json:"address,omitempty" url:"-"`
	MedicalProxy      *GuarantorDetails  `json:"medical_proxy,omitempty" url:"-"`
	Race              *Race              `json:"race,omitempty" url:"-"`
	Ethnicity         *Ethnicity         `json:"ethnicity,omitempty" url:"-"`
	SexualOrientation *SexualOrientation `json:"sexual_orientation,omitempty" url:"-"`
	GenderIdentity    *GenderIdentity    `json:"gender_identity,omitempty" url:"-"`
}
