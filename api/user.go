package franky

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           string
	Email        string
	Password     string
	ApiKey       string
	CreationDate int64
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
	user.CreationDate = time.Now().Unix()
}

func (user *User) generateApiKey() {
	user.ApiKey = "franky.api." + uuid.NewString()
}

func (user *User) generateId() {
	user.Id = "franky.user." + uuid.NewString()
}
