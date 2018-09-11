package model

// User struct
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"varchar(64)"`
	Email        string `gorm:"varchar(12)"`
	PasswordHash string `gorm:"varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}
