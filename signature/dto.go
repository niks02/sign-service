package signature

// Request struct represents signature request body
type Request struct {
	IDs []uint64 `json:"ids"`
}

// Response struct represents signature response body
type Response struct {
	Message   []string `json:"message"`
	Signature string   `json:"signature"`
}
