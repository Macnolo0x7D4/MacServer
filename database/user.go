package database

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Role      int
	Level     int
	CreatedAt time.Time
}

var emailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type Users []User

func NewUser(username, password, email string, role, level int) *User {
	passwordEncrypted := SetPassword(password)
	user := &User{Username: username, Password: passwordEncrypted, Email: email, Role: role, Level: level}
	return user
}

func CreateDefaultUser() *User {
	user := NewUser("Macnolo0x7D4", "root", "yosoymacnolo@gmail.com", 0, 255)
	user.Save()
	return user
}

func GetUsers() Users {
	users := Users{}
	Database.Find(&users)
	return users
}

func GetUserById(id int) *User {
	user := &User{}
	Database.Where("id=?", id).First(user)
	return user
}

func GetUserByUsername(username string) *User {
	user := &User{}
	Database.Where("username=?", username).First(user)
	return user
}

func (this *User) Save() (bool, error) {
	if this.Id == 0 {
		Database.Create(&this)
	} else {
		user := &User{Username: this.Username, Password: this.Password, Email: this.Email}
		Database.Model(&this).UpdateColumns(user)
	}

	return false, nil
}

func (this *User) Delete() {
	Database.Delete(&this)
}

func SetPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func Login(username, password string) bool {
	user := GetUserByUsername(username)
	if result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); result != nil {
		return false
	} else {
		return true
	}
}

func (this *User) Valid() (bool, error) {
	if len(this.Username) > 20 {
		err := errors.New("The username is too large")
		return false, err
	}

	if !emailRegexp.MatchString(this.Email) {
		err := errors.New("The email format is not vaild")
		return false, err
	}

	if len(this.Password) != 60 {
		this.Password = SetPassword(this.Password)
	}

	if this.Role == 0 {
		this.Role = 32
	}

	return false, nil
}
