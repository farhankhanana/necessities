package model

import (
	"encoding/json"
	"mime/multipart"
)

// Metadata defines model meta data for every protocol, e.g: HTTP, RPC
type Metadata struct {
	Body          interface{}           `json:"body,omitempty"`
	Header        map[string]string     `json:"header,omitempty"`
	Headers       map[string]string     `json:"headers,omitempty"`
	Params        map[string]string     `json:"params,omitempty"`
	File          *multipart.FileHeader `json:"file,omitempty"`
	RequestID     string                `json:"request_id,omitempty"`
	CorrelationID string                `json:"correlation_id,omitempty"`
	Path          string                `json:"path,omitempty"`
	BodyRaw       []byte                `json:"body_raw,omitempty"`
}

// HTTPQueryParameter defines model for query parameter used for filtering data
type HTTPQueryParameter struct {
	OrderBy string      `json:"order_by,omitempty" query:"order_by"`
	Desc    string      `json:"desc,omitempty" query:"desc"`
	Limit   json.Number `json:"limit,omitempty" query:"limit"`
	Offset  json.Number `json:"offset,omitempty" query:"offset"`
}
