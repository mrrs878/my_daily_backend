package controller

type EmailAdd struct {
	Email string `form:"email" json:"email" validate:"required,EmailValid"`
}
