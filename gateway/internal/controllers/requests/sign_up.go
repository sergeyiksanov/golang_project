package requests

type SignUpRequest struct {
	Email    string `json:"email"`
	Passwotd string `json:"passwotd"`
}
