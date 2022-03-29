package internal

type User struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"password"`
}
