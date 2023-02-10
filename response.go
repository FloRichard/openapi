package openapi

// Response -.
type Response struct {
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

func (r *Response) IsRef() bool {
	return r.Ref != ""
}

// Responses -.
type Responses map[string]*Response
