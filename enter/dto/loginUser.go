package dto

type LoginUser struct {
	Phone    string `json:"phone,omitempty"`
	Code     string `json:"code,omitempty"`
	Password string `json:"password,omitempty"`
}
