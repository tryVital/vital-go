// This file was auto-generated by Fern from our API Definition.

package api

type CreateRegistrableTestkitOrderRequest struct {
	UserId          string                         `json:"user_id" url:"-"`
	LabTestId       string                         `json:"lab_test_id" url:"-"`
	ShippingDetails *ShippingAddressWithValidation `json:"shipping_details,omitempty" url:"-"`
	Passthrough     *string                        `json:"passthrough,omitempty" url:"-"`
}

type RegisterTestkitRequest struct {
	// The user ID of the patient.
	UserId          *string                       `json:"user_id,omitempty" url:"-"`
	SampleId        string                        `json:"sample_id" url:"-"`
	PatientDetails  *PatientDetailsWithValidation `json:"patient_details,omitempty" url:"-"`
	PatientAddress  *PatientAddressWithValidation `json:"patient_address,omitempty" url:"-"`
	Physician       *PhysicianCreateRequestBase   `json:"physician,omitempty" url:"-"`
	HealthInsurance *HealthInsuranceCreateRequest `json:"health_insurance,omitempty" url:"-"`
	Consents        []*Consent                    `json:"consents,omitempty" url:"-"`
}
