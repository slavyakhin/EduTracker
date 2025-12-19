package models

type App struct {
	ID      int
	Name    string
	Sercret string // Models can be logged, so it's better not to store secrets
}
