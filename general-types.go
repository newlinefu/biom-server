package main

type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Vector   []float64 `json:"vector"`
}

type RegistrationObject struct {
	Name     string                    `json:"name"`
	Email    string                    `json:"email"`
	Password string                    `json:"password"`
	Records  []KeyboardHandwritingData `json:"records"`
}

type AuthorizationRequestObject struct {
	Email    string                    `json:"email"`
	Password string                    `json:"password"`
	Records  []KeyboardHandwritingData `json:"records"`
}
type AuthorizationResponseObject struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}
