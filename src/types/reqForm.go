package types

type LoginForm struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}

type RegisterForm struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}

type CreateDataDictForm struct {
	GroupName string `binding:"required"`
	Label     string `binding:"required"`
	Value     string `binding:"required"`
}

type UpdateDataDictForm struct {
	CreateDataDictForm
	Id uint `binding:"required"`
}

type CreateGoodsForm struct {
	Name        string `binding:"required"`
	Service     string
	Class       uint   `binding:"required"`
	Description string `binding:"required"`
}

type UpdateGoodForm struct {
	CreateGoodsForm
	Id uint `binding:"required"`
}

type CreateTaskForm struct {
	Title  string `binding:"required"`
	Label  string `binding:"required"`
	Detail string `binding:"required"`
	Status uint   `binding:"required"`
	UserId uint   `binding:"required"`
}
