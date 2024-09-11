package registermiddlewaresubfunction

import (
	"errors"
	"fmt"
	"net/http"
	model "social-network/Model"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterVerification(register model.Register, w http.ResponseWriter) error {
	nw := model.ResponseWriter{
		ResponseWriter: w,
	}

	if register.Auth.Password != register.Auth.ConfirmPassword {
		nw.Error("Password and password confirmation do not match")
		return errors.New("password and password confirmation do not match")
	}

	if register.Auth.Email == "" || register.Auth.Password == "" || register.FirstName == "" || register.LastName == "" || register.BirthDate == "" {
		nw.Error("There is an empty field")
		return errors.New("there is an empty field")
	}

	return nil
}

func CreateUuidAndCrypt(register *model.Register, w http.ResponseWriter) error {
	nw := model.ResponseWriter{
		ResponseWriter: w,
	}

	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Auth.Password), 12)
	if err != nil {
		fmt.Println(err)

		nw.Error("Internal Error: There is a probleme with bcrypt")
		return errors.New("there is a probleme with bcrypt")
	}
	register.Auth.Password = string(cryptedPassword)

	uuid, err := uuid.NewV7()
	if err != nil {
		fmt.Println(err)
		nw.Error("Internal Error: There is a probleme with the generation of the uuid")
		return errors.New("there is a probleme with the generation of the uuid")
	}
	register.Auth.Id = uuid.String()

	return nil
}
