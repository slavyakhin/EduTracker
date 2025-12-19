package models

type User struct {
	ID       int64
	Name     string
	PassHash []byte
}
