package user_repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/matheuswww/mystream/src/logger"
	rest_err "github.com/matheuswww/mystream/src/restErr"
	"golang.org/x/crypto/bcrypt"
)

func (ur *userRepository) Signin(email, password string) (string,*rest_err.RestErr ) {
	ctx,cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := "SELECT id,password FROM users WHERE email = $1"
	var id,encryptedPassword string
	err := ur.sql.QueryRowContext(ctx, query, email).Scan(&id, &encryptedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "",rest_err.NewNotFoundError("email not found")
		}
		logger.Error(fmt.Sprintf("Error trying QueryRowContext: %v", err))
		return "",rest_err.NewInternalServerError("server error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "",rest_err.NewBadRequestError("wrong password")
		}
		logger.Error(fmt.Sprintf("Error trying CompareHashAndPassword: %v", err))
		return "",rest_err.NewInternalServerError("server error")
	}
	return id, nil
}