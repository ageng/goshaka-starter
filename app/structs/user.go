package structs

type UserCreate struct {
	Username             string `json:"username" validate:"required"`
	Email                string `json:"email" validate:"required"`
	FirstName            string `json:"first_name" validate:"required"`
	LastName             string `json:"last_name"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}

type UserUpdate struct {
	Username             string `json:"username" validate:"required"`
	Email                string `json:"email" validate:"required"`
	FirstName            string `json:"first_name" validate:"required"`
	LastName             string `json:"last_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation" validate:"eqfield=Password"`
}