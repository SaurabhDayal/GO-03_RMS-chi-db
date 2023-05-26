package dbHelper

import (
	"06_RMS-chi-db/database"
	"06_RMS-chi-db/errorHandling"
	"06_RMS-chi-db/models"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

const userCredit = models.UserCredit

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(usr *models.Users) (*models.UsersClient, error, string) {
	var id int
	err := database.RMS.Get(&id, "SELECT id FROM users WHERE user_email=$1 AND archived_at IS NULL", usr.UserEmail)
	if id != 0 {
		return nil, errorHandling.UserAlreadyPresent(), ""
	}
	password := usr.UserPassword
	hash, err := HashPassword(password)
	if err != nil {
		return nil, errorHandling.UnableToGenerateToken(), ""
	}
	var user models.UsersClient
	SQL1 := `INSERT INTO users (user_name, user_password, user_email, credit) VALUES ($1,$2,$3,$4) RETURNING id, user_name, user_password, user_email, credit`
	err = database.RMS.Get(&user, SQL1, usr.UserName, hash, usr.UserEmail, userCredit)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB(), ""
	}
	SQL2 := `INSERT INTO user_roles (user_id, role) VALUES ($1, $2)`
	_, err = database.RMS.Exec(SQL2, user.Id, "user")
	if err != nil {
		return nil, errorHandling.UnableToAccessDB(), ""
	}
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return nil, errorHandling.UnableToGenerateToken(), ""
	}
	b = []byte(hex.EncodeToString(b))
	var token string
	SQL3 := `INSERT INTO auths (user_id, user_token) VALUES ($1, $2) RETURNING user_token`
	err = database.RMS.Get(&token, SQL3, user.Id, b)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB(), ""
	}
	return &user, nil, token
}

func LoginUser(email string, pwd string) (*models.Auths, error) {
	SQL1 := `SELECT user_password FROM users WHERE user_email=$1`
	var hashPassword string
	err := database.RMS.Get(&hashPassword, SQL1, email)
	if err != nil {
		return nil, errorHandling.IncorrectEmail()
	}
	match := CheckPasswordHash(pwd, hashPassword)
	if match {
		b := make([]byte, 6)
		if _, err := rand.Read(b); err != nil {
			return nil, errorHandling.UnableToGenerateToken()
		}
		b = []byte(hex.EncodeToString(b))

		var userId int
		SQL4 := `SELECT id FROM users WHERE user_email=$1`
		err = database.RMS.Get(&userId, SQL4, email)
		if err != nil && err != sql.ErrNoRows {
			return nil, errorHandling.UnableToAccessDB()
		} else if err == sql.ErrNoRows {
			SQL2 := `INSERT INTO auths (user_id, user_token) VALUES ($1, $2)`
			_, err = database.RMS.Exec(SQL2, userId, b)
			if err != nil {
				return nil, errorHandling.UnableToAccessDB()
			}
		} else {
			SQL5 := `DELETE FROM auths WHERE user_id=$1`
			_, err = database.RMS.Exec(SQL5, userId)
			if err != nil {
				return nil, errorHandling.UnableToAccessDB()
			}
			SQL2 := `INSERT INTO auths (user_id, user_token) VALUES ($1, $2)`
			_, err = database.RMS.Exec(SQL2, userId, b)
			if err != nil {
				return nil, errorHandling.UnableToAccessDB()
			}
		}

		var usrSn models.Auths
		SQL3 := `SELECT user_id, user_token FROM auths WHERE user_id=$1`
		err = database.RMS.Get(&usrSn, SQL3, userId)
		if err != nil {
			return nil, errorHandling.UnableToAccessDB()
		}
		return &usrSn, nil
	}
	return nil, errorHandling.UserCredentialNotMatch()
}

func LogoutUser(token string) error {
	SQL := `DELETE FROM auths WHERE user_token = $1`
	_, err := database.RMS.Exec(SQL, token)
	if err != nil {
		return errorHandling.UnableToAccessDB()
	}
	return nil
}
