package model

type ReqUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=10"`
	Email    string `json:"email" validate:"required,email"`
	IsAdmin  *bool  `json:"isAdmin"`
}

type ReqLogin struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
