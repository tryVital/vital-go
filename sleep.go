// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/tryVital/vital-go/core"
	time "time"
)

type SleepGetRequest struct {
	// Provider oura/strava etc
	Provider *string `json:"-" url:"provider,omitempty"`
	// Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
	StartDate string `json:"-" url:"start_date"`
	// Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
	EndDate *string `json:"-" url:"end_date,omitempty"`
}

type SleepGetRawRequest struct {
	// Provider oura/strava etc
	Provider *string `json:"-" url:"provider,omitempty"`
	// Date from in YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 00:00:00
	StartDate string `json:"-" url:"start_date"`
	// Date to YYYY-MM-DD or ISO formatted date time. If a date is provided without a time, the time will be set to 23:59:59
	EndDate *string `json:"-" url:"end_date,omitempty"`
}

type ClientFacingSleep struct {
	Id string `json:"id" url:"id"`
	// User id returned by vital create user request. This id should be stored in your database against the user and used for all interactions with the vital api.
	UserId string `json:"user_id" url:"user_id"`
	// Date of the specified record, formatted as ISO8601 datetime string in UTC 00:00. Deprecated in favour of calendar_date.
	Date time.Time `json:"date" url:"date"`
	// Date of the sleep summary in the YYYY-mm-dd format. This generally matches the sleep end date.
	CalendarDate string `json:"calendar_date" url:"calendar_date"`
	// UTC Time when the sleep period started
	BedtimeStart time.Time `json:"bedtime_start" url:"bedtime_start"`
	// UTC Time when the sleep period ended
	BedtimeStop time.Time `json:"bedtime_stop" url:"bedtime_stop"`
	// `long_sleep`: >=3 hours of sleep;
	// `short_sleep`: <3 hours of sleep;
	// `acknowledged_nap`: User-acknowledged naps, typically under 3 hours of sleep;
	// `unknown`: The sleep session recording is ongoing.
	Type SleepType `json:"type" url:"type"`
	// Timezone offset from UTC as seconds. For example, EEST (Eastern European Summer Time, +3h) is 10800. PST (Pacific Standard Time, -8h) is -28800::seconds
	TimezoneOffset *int `json:"timezone_offset,omitempty" url:"timezone_offset,omitempty"`
	// Total duration of the sleep period (sleep.duration = sleep.bedtime_end - sleep.bedtime_start)::seconds
	Duration int `json:"duration" url:"duration"`
	// Total amount of sleep registered during the sleep period (sleep.total = sleep.rem + sleep.light + sleep.deep)::seconds
	Total int `json:"total" url:"total"`
	// Total amount of awake time registered during the sleep period::seconds
	Awake int `json:"awake" url:"awake"`
	// Total amount of light sleep registered during the sleep period::seconds
	Light int `json:"light" url:"light"`
	// Total amount of REM sleep registered during the sleep period, minutes::seconds
	Rem int `json:"rem" url:"rem"`
	// Total amount of deep (N3) sleep registered during the sleep period::seconds
	Deep int `json:"deep" url:"deep"`
	// A value between 1 and 100 representing how well the user slept. Currently only available for Withings, Oura, Whoop and Garmin::scalar
	Score *int `json:"score,omitempty" url:"score,omitempty"`
	// The lowest heart rate (5 minutes sliding average) registered during the sleep period::beats per minute
	HrLowest *int `json:"hr_lowest,omitempty" url:"hr_lowest,omitempty"`
	// The average heart rate registered during the sleep period::beats per minute
	HrAverage *int `json:"hr_average,omitempty" url:"hr_average,omitempty"`
	// Resting heart rate recorded during a sleep session::bpm
	HrResting *int `json:"hr_resting,omitempty" url:"hr_resting,omitempty"`
	// Sleep efficiency is the percentage of the sleep period spent asleep (100% * sleep.total / sleep.duration)::perc
	Efficiency *float64 `json:"efficiency,omitempty" url:"efficiency,omitempty"`
	// Detected latency from bedtime_start to the beginning of the first five minutes of persistent sleep::seconds
	Latency *int `json:"latency,omitempty" url:"latency,omitempty"`
	// Skin temperature deviation from the long-term temperature average::celcius
	TemperatureDelta *float64 `json:"temperature_delta,omitempty" url:"temperature_delta,omitempty"`
	// The skin temperature::celcius
	SkinTemperature *float64 `json:"skin_temperature,omitempty" url:"skin_temperature,omitempty"`
	// Sleeping Heart Rate Dip is the percentage difference between your average waking heart rate and your average sleeping heart rate. In health studies, a greater "dip" is typically seen as a positive indicator of overall health. Currently only available for Garmin::perc
	HrDip *float64 `json:"hr_dip,omitempty" url:"hr_dip,omitempty"`
	// Some providers can provide updates to the sleep summary hours after the sleep period has ended. This field indicates the state of the sleep summary. For example, TENTATIVE means the summary is an intial prediction from the provider and can be subject to change. Currently only available for Garmin and EightSleep::str
	State *SleepSummaryState `json:"state,omitempty" url:"state,omitempty"`
	// The average heart rate variability registered during the sleep period::rmssd
	AverageHrv *float64 `json:"average_hrv,omitempty" url:"average_hrv,omitempty"`
	// Average respiratory rate::breaths per minute
	RespiratoryRate *float64 `json:"respiratory_rate,omitempty" url:"respiratory_rate,omitempty"`
	// Source the data has come from.
	Source      *ClientFacingSource      `json:"source,omitempty" url:"source,omitempty"`
	SleepStream *ClientFacingSleepStream `json:"sleep_stream,omitempty" url:"sleep_stream,omitempty"`
	CreatedAt   time.Time                `json:"created_at" url:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at" url:"updated_at"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientFacingSleep) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientFacingSleep) UnmarshalJSON(data []byte) error {
	type embed ClientFacingSleep
	var unmarshaler = struct {
		embed
		Date         *core.DateTime `json:"date"`
		BedtimeStart *core.DateTime `json:"bedtime_start"`
		BedtimeStop  *core.DateTime `json:"bedtime_stop"`
		CreatedAt    *core.DateTime `json:"created_at"`
		UpdatedAt    *core.DateTime `json:"updated_at"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = ClientFacingSleep(unmarshaler.embed)
	c.Date = unmarshaler.Date.Time()
	c.BedtimeStart = unmarshaler.BedtimeStart.Time()
	c.BedtimeStop = unmarshaler.BedtimeStop.Time()
	c.CreatedAt = unmarshaler.CreatedAt.Time()
	c.UpdatedAt = unmarshaler.UpdatedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientFacingSleep) MarshalJSON() ([]byte, error) {
	type embed ClientFacingSleep
	var marshaler = struct {
		embed
		Date         *core.DateTime `json:"date"`
		BedtimeStart *core.DateTime `json:"bedtime_start"`
		BedtimeStop  *core.DateTime `json:"bedtime_stop"`
		CreatedAt    *core.DateTime `json:"created_at"`
		UpdatedAt    *core.DateTime `json:"updated_at"`
	}{
		embed:        embed(*c),
		Date:         core.NewDateTime(c.Date),
		BedtimeStart: core.NewDateTime(c.BedtimeStart),
		BedtimeStop:  core.NewDateTime(c.BedtimeStop),
		CreatedAt:    core.NewDateTime(c.CreatedAt),
		UpdatedAt:    core.NewDateTime(c.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (c *ClientFacingSleep) String() string {
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

type ClientFacingSleepStream struct {
	Hrv             []*ClientFacingHrvTimeseries             `json:"hrv,omitempty" url:"hrv,omitempty"`
	Heartrate       []*ClientFacingHeartRateTimeseries       `json:"heartrate,omitempty" url:"heartrate,omitempty"`
	Hypnogram       []*ClientFacingHypnogramTimeseries       `json:"hypnogram,omitempty" url:"hypnogram,omitempty"`
	RespiratoryRate []*ClientFacingRespiratoryRateTimeseries `json:"respiratory_rate,omitempty" url:"respiratory_rate,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientFacingSleepStream) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientFacingSleepStream) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientFacingSleepStream
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientFacingSleepStream(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientFacingSleepStream) String() string {
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

type ClientSleepResponse struct {
	Sleep []*ClientFacingSleep `json:"sleep,omitempty" url:"sleep,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ClientSleepResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ClientSleepResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientSleepResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientSleepResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientSleepResponse) String() string {
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

type RawSleep struct {
	Sleep []*SleepV2InDb `json:"sleep,omitempty" url:"sleep,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RawSleep) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RawSleep) UnmarshalJSON(data []byte) error {
	type unmarshaler RawSleep
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RawSleep(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RawSleep) String() string {
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

// ℹ️ This enum is non-exhaustive.
type SleepSummaryState string

const (
	SleepSummaryStateTentative SleepSummaryState = "tentative"
	SleepSummaryStateConfirmed SleepSummaryState = "confirmed"
)

func NewSleepSummaryStateFromString(s string) (SleepSummaryState, error) {
	switch s {
	case "tentative":
		return SleepSummaryStateTentative, nil
	case "confirmed":
		return SleepSummaryStateConfirmed, nil
	}
	var t SleepSummaryState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SleepSummaryState) Ptr() *SleepSummaryState {
	return &s
}

// ℹ️ This enum is non-exhaustive.
type SleepType string

const (
	SleepTypeLongSleep       SleepType = "long_sleep"
	SleepTypeShortSleep      SleepType = "short_sleep"
	SleepTypeAcknowledgedNap SleepType = "acknowledged_nap"
	SleepTypeUnknown         SleepType = "unknown"
)

func NewSleepTypeFromString(s string) (SleepType, error) {
	switch s {
	case "long_sleep":
		return SleepTypeLongSleep, nil
	case "short_sleep":
		return SleepTypeShortSleep, nil
	case "acknowledged_nap":
		return SleepTypeAcknowledgedNap, nil
	case "unknown":
		return SleepTypeUnknown, nil
	}
	var t SleepType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SleepType) Ptr() *SleepType {
	return &s
}

type SleepV2InDb struct {
	Timestamp      time.Time              `json:"timestamp" url:"timestamp"`
	Data           map[string]interface{} `json:"data,omitempty" url:"data,omitempty"`
	ProviderId     string                 `json:"provider_id" url:"provider_id"`
	UserId         string                 `json:"user_id" url:"user_id"`
	SourceId       int                    `json:"source_id" url:"source_id"`
	PriorityId     *int                   `json:"priority_id,omitempty" url:"priority_id,omitempty"`
	Id             string                 `json:"id" url:"id"`
	Source         *ClientFacingProvider  `json:"source,omitempty" url:"source,omitempty"`
	Priority       *int                   `json:"priority,omitempty" url:"priority,omitempty"`
	SourceDeviceId *string                `json:"source_device_id,omitempty" url:"source_device_id,omitempty"`
	CreatedAt      *time.Time             `json:"created_at,omitempty" url:"created_at,omitempty"`
	UpdatedAt      *time.Time             `json:"updated_at,omitempty" url:"updated_at,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SleepV2InDb) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SleepV2InDb) UnmarshalJSON(data []byte) error {
	type embed SleepV2InDb
	var unmarshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
		UpdatedAt *core.DateTime `json:"updated_at,omitempty"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SleepV2InDb(unmarshaler.embed)
	s.Timestamp = unmarshaler.Timestamp.Time()
	s.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	s.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SleepV2InDb) MarshalJSON() ([]byte, error) {
	type embed SleepV2InDb
	var marshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
		UpdatedAt *core.DateTime `json:"updated_at,omitempty"`
	}{
		embed:     embed(*s),
		Timestamp: core.NewDateTime(s.Timestamp),
		CreatedAt: core.NewOptionalDateTime(s.CreatedAt),
		UpdatedAt: core.NewOptionalDateTime(s.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (s *SleepV2InDb) String() string {
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
