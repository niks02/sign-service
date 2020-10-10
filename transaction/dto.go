package transaction

// Request struct represents transaction request body
type Request struct {
	Transaction string `json:"txn"`
}

// Response struct represents transaction response body
type Response struct {
	ID uint64 `json:"id"`
}
