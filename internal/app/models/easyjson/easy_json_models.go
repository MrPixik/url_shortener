package easyjson

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
	URL string `json:"result"`
}
