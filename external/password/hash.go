package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func MatchPassword(hashedPassword, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currentPassword))
	return err == nil
}
