package easyjson

// easyjson:json
type URLDB struct {
	ID       int64  `json:"id"`
	Original string `json:"original"`
	Short    string `json:"short"`
}

// easyjson:json
type URLFileRecord struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

// easyjson:json
type URLRequest struct {
	URL string `json:"url"`
}

// easyjson:json
type URLResponse struct {
	URL string `json:"short-url"`
}
