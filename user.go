// This file was auto-generated by Fern from our API Definition.

package api

type UserCreateBody struct {
	// A unique ID representing the end user. Typically this will be a user ID from your application. Personally identifiable information, such as an email address or phone number, should not be used in the client_user_id.
	ClientUserId string `json:"client_user_id"`
	// Fallback time zone of the user, in the form of a valid IANA tzdatabase identifier (e.g., `Europe/London` or `America/Los_Angeles`).
	// Used when pulling data from sources that are completely time zone agnostic (e.g., all time is relative to UTC clock, without any time zone attributions on data points).
	FallbackTimeZone *string `json:"fallback_time_zone,omitempty"`
}

type UserGetAllRequest struct {
	Offset *int `json:"-"`
	Limit  *int `json:"-"`
}

type UserPatchBody struct {
	// Fallback time zone of the user, in the form of a valid IANA tzdatabase identifier (e.g., `Europe/London` or `America/Los_Angeles`).
	// Used when pulling data from sources that are completely time zone agnostic (e.g., all time is relative to UTC clock, without any time zone attributions on data points).
	FallbackTimeZone *string `json:"fallback_time_zone,omitempty"`
}

type UserRefreshRequest struct {
	Timeout *float64 `json:"-"`
}
