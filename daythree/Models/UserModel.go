package Models

import (
	"bootcamp/daythree/Config"
	"fmt"
	"time"
)

//	type Person struct {
//		gorm.Model
//		FirstName string `json:"firstname" binding:"required"`
//		LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(256)"`
//		Age       int8   `json:"age" binding:"gte=1,lte=130" gorm:"type:int"`
//		Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
//	}
//
//	type Video struct {
//		ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
//		Title       string    `json:"title" binding:"min=2,max=10" validate:"is-cool" gorm:"type:varchar(100)"`
//		Description string    `json:"description" binding:"max=20" gorm:"type:varchar(100)"`
//		URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
//		Author      Person    `json:"author" binding:"required" gorm:"foreignKey:PersonID;references:ID"`
//		PersonId    uint64    `json:"-"`
//		CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
//		UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
//	}
type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Phone     string    `json:"phone" binding:"required,lte=10"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP(3)" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP(3)" json:"updated_at"`
}

func (b *User) TableName() string {
	return "user"
}
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	fmt.Println("users doc users", user)
	return Config.DB.Create(user).Error
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Where("id = ?", id).Updates(user)
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
