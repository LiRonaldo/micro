package DBModels

import "time"

type Users struct {
	UserId   int       `gorm:"cloumn:user_id;AUTO_INCREMENT;PRIMARY_KEY"`
	UserName string    `gorm:"cloumn:user_name;type:varchar(50);unique_index"`
	UserPwd  string    `gorm:"cloumn:user_pwd;type:varchar(50)"`
	UserDate time.Time `gorm:"cloumn:user_date;"`
}
