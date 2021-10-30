package gym

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/google/uuid"
)

// ErrInvalidUserID USER ID ------------
var ErrInvalidUserID = errors.New("invalid User ID")

// UserID  represent the user that own the session
type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{
		value: v.String(),
	}, nil

}

func (id UserID) String() string {
	return id.value
}

// ErrInvalidEmail EMAIL------------
var ErrInvalidEmail = errors.New("invalid email")

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	v, err := mail.ParseAddress(value)
	if err != nil {
		return Email{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return Email{
		value: v.String(),
	}, nil

}

func (v Email) String() string {
	return v.value
}

// PASSWORD------------ FOR TEST ONLY IN RAW
var ErrEmptyPassword = errors.New("the field Pasword can not be empty")

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	if value == "" {
		return Password{}, ErrEmptyPassword
	}

	return Password{
		value: value,
	}, nil

}

// Domanin Object  -------

type User struct {
	id       UserID
	email    Email
	password Password
}

func NewUser(userID, email, password string) (User, error) {

	userIdVO, err := NewUserID(userID)
	if err != nil {
		return User{}, err
	}

	emailVO, err := NewEmail(email)
	if err != nil {
		return User{}, err
	}

	passwordVO, err := NewPassword(password)
	if err != nil {
		return User{}, err
	}

	user := User{
		id:       userIdVO,
		email:    emailVO,
		password: passwordVO,
	}

	return user, nil

}

func (u User) ID() UserID {
	return u.id
}

func (u User) Email() Email {
	return u.email
}

func (u User) Password() Password {
	return u.password
}
