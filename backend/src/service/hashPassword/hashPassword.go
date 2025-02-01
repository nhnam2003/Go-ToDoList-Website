package hashpassword

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error ) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}
