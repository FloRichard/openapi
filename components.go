package openapi

// Components -.
type Components struct {
	Schemas       Schemas       `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	RequestBodies RequestBodies `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Responses     Responses     `json:"responses,omitempty" yaml:"responses,omitempty"`
}
