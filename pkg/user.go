package franky

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           string
	Email        string
	Password     string
	ApiKey       string
	CreationDate string
}

func (user *User) saltAndHashPassword() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return nil
}

func (user *User) setCreationDate() {
	user.CreationDate = fmt.Sprint(time.Now().UnixNano())
}

func (user *User) generateApiKey() {
	user.ApiKey = "andykhv.franky." + uuid.NewString()
}
