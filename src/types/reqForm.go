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
	Title     string `binding:"required" json:"title"`
	Label     string `binding:"required" json:"label"`
	Detail    string `binding:"required" json:"detail"`
	Status    int    `binding:"required" json:"status"`
	AlarmTime uint64 `binding:"required" json:"alarmTime"`
}

type UpdateTaskForm struct {
	Title     string `json:"title"`
	Label     string `json:"label"`
	Detail    string `json:"detail"`
	Status    int    `json:"status"`
	AlarmTime uint64 `json:"alarmTime"`
}
