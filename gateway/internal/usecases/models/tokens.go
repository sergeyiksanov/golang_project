package models

type RefreshToken string

type Tokens struct {
	Access  string
	Refresh string
}
