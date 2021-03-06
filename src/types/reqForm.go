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
	Title     string `json:"title" binding:"required"`
	Label     string `json:"label"`
	Detail    string `json:"detail" binding:"required"`
	AlarmTime uint64 `json:"alarmTime" binding:"required"`
	Status    int    `json:"status"`
}

type UpdateTaskForm struct {
	Title     string `json:"title"`
	Label     string `json:"label"`
	Detail    string `json:"detail"`
	Status    int    `json:"status"`
	AlarmTime uint64 `json:"alarmTime"`
}

type SubscriptionForm struct {
	Endpoint             string `binding:"required" json:"endpoint"`
	ExpirationTime       uint   `json:"expirationTime"`
	ApplicationServerKey string `binding:"required" json:"p256dh"`
	Auth                 string `binding:"required" json:"auth"`
}

type PushMsgForm struct {
	Detail string `binding:"required" json:"detail"`
	UserId uint   `binding:"required" json:"userId"`
}

type CreateHabitForm struct {
	Title     string `json:"title" binding:"required"`
	Label     string `json:"label"`
	Detail    string `json:"detail" binding:"required"`
	Status    int    `json:"status"`
	AlarmTime string `json:"alarmTime" binding:"required"`
	AlarmDate string `json:"alarmDate" binding:"required"`
}
