package openapi

// Operation -.
type Operation struct {
	OperationId string 	`json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters  Parameters  `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody RequestBody `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   Responses   `json:"responses" yaml:"responses"`
}
