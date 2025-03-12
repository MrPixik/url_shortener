package models

type UrlsObj struct {
	ID       int64
	Original string
	Short    string
	UserId   int
}
type URLMapping struct {
	OrigURL  string
	ShortURL string
}
