package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/solrac87/rest/src/api/security"
)

// User struct
type User struct {
	ID        int64     `bson:"id" json:"id"`
	NickName  string    `bson:"nickname" json:"nickname"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// HashPassword encode user password
func (u *User) HashPassword(p string) error {
	hashedPassword, err := security.Hash(p)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.NickName = html.EscapeString(strings.TrimSpace(u.NickName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.HashPassword(u.Password)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {

	if u.NickName == "" {
		return errors.New("Required Nickaname")
	}

	if u.Password == "" && action == "create" {
		return errors.New("Required Password")
	}

	if u.Email == "" {
		return errors.New("Required Email")
	}

	if strings.Contains(u.Email, "@") == false {
		return errors.New("Invalid Email")
	}

	return nil
}
