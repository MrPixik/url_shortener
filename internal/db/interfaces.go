package db

type DatabaseService interface {
	Ping() error
}
