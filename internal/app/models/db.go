package models

type URLDB struct {
	ID       int64
	Original string
	Short    string
}
type URLMapping struct {
	OrigURL  string
	ShortURL string
}
