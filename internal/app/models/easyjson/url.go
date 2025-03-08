package easyjson

//go:generate easyjson -all url.go

// URLRequest struct for processing requests with single JSON URL.
// easyjson:json
type URLRequest struct {
	OrigURL string `json:"url"`
}

// easyjson:json
type URLRequestArrELem struct {
	Id      string `json:"correlation_id"`
	OrigURL string `json:"original_url"`
}

// URLRequestArr type for processing requests with JSON-array URLs.
// easyjson:json
type URLRequestArr []URLRequestArrELem

// URLResponse struct for writing responses with single JSON URL.
// easyjson:json
type URLResponse struct {
	ShortURL string `json:"short-url"`
}

// easyjson:json
type URLResponseArrElem struct {
	Id       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

// URLResponseArr type for writing responses with JSON-array URLs.
//
//easyjson:json
type URLResponseArr []URLResponseArrElem
