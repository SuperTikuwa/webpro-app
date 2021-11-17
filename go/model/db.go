package model

type User struct {
	ID       int64  `gorm:"id"`
	Name     string `gorm:"name"`
	Password string `gorm:"password"`
}

type Schedule struct {
	ID          int64  `gorm:"id"`
	UserID      int64  `gorm:"user_id"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Place       string `gorm:"place"`
	StartTime   string `gorm:"start_time"`
	EndTime     string `gorm:"end_time"`
}
