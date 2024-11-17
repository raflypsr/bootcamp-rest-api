package models

import (
	"html"
	"strings"
	"time"

	"fp-super-bootcamp-go/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Users struct {
		ID        uint       `gorm:"primary_key" json:"id"`
		Email     string     `gorm:"type:varchar(255); unique; not null;" json:"email"`
		Username  string     `gorm:"type:varchar(255); not null;" json:"username"`
		Password  string     `gorm:"not null;" json:"password"`
		Role      string     `json:"role"`
		CreatedAt time.Time  `gorm:"not null;" json:"created_at"`
		UpdatedAt time.Time  `gorm:"not null;" json:"updated_at"`
		Profile   Profile    `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
		Reviews   []Reviews  `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
		Likes     []Likes    `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
		Comments  []Comments `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	}
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Users{}

	err = db.Model(Users{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Users) SaveUser(db *gorm.DB) (*Users, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Users{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}
