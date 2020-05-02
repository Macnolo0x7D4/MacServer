package database

import (
	"../models"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
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
	user := &User{Username: username, Password: password, Email: email, Role: role, Level: level}
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

func GetUserByEmail(email string) *User {
	user := &User{}
	Database.Where("email=?", email).First(user)
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

func (this *User) EncryptPassword(){
	hash, _ := bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)
	this.Password = string(hash)
}

func EncryptPassword(passwd string) string{
	hash, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(hash)
}

func Login(email, password string) (*User, error){
	user := GetUserByEmail(email)

	if(user.Id == 0){
		return nil, models.NotFoundAccountError
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if(err != nil){
		return nil, models.AccessDeniedError
	}

	return user, nil
}

func (this *User) Valid() error{
	if lenUsername := len(this.Username); lenUsername == 0 {
		return models.NoUsernameError
	} else if lenUsername < 4{
		return models.ShortUsernameError
	} else if lenUsername > 20{
		return models.LongUsernameError
	}

	if lenPassword := len(this.Password); lenPassword == 0{
		return models.NoPasswordError
	} else if lenPassword < 8{
		return models.WeakPasswordError
	} else if lenPassword > 31{
		return models.LongPasswordError
	}

	this.EncryptPassword()

	if len(this.Password) != 60 {
		return models.IsPasswordEncryptedError
	}

	if lenEmail := len(this.Email); lenEmail == 0{
		return models.NoEmailError
	} else if lenEmail < 12{
		return models.ShortEmailError
	} else if lenEmail > 40{
		return models.LongEmailError
	}

	if !emailRegexp.MatchString(this.Email) {
		return models.WrongEmailFormatError
	}

	if GetUserByUsername(this.Username).Id != 0{
		return models.DuplicatedUsernameError
	}

	if GetUserByEmail(this.Email).Id != 0{
		return models.DuplicatedEmailError
	}

	if this.Role == 0 {
		this.Role = 32
	}

	return nil
}
