package types

type LoginForm struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}

type RegisterForm struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}
