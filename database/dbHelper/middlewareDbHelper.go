package dbHelper

import (
	"03_RMS/database"
)

func CheckUserRole(token string) string {
	SQL1 := `SELECT user_id FROM auths WHERE user_token=$1`
	var userId int
	err := database.RMS.Get(&userId, SQL1, token)
	if err != nil {
		return ""
	}
	usrRoles := make([]string, 0)
	SQL2 := `SELECT role FROM user_roles WHERE user_id=$1`
	err = database.RMS.Select(&usrRoles, SQL2, userId)
	if err != nil {
		return ""
	}
	for _, v := range usrRoles {
		if v == "admin" {
			return "admin"
		}
		if v == "subAdmin" {
			return "subAdmin"
		}
		if v == "user" {
			return "user"
		}
	}
	return ""
}

func CheckUserId(token string) (int, error) {
	SQL1 := `SELECT user_id FROM auths WHERE user_token=$1`
	var userId int
	err := database.RMS.Get(&userId, SQL1, token)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func CheckOwnerRestaurantsIds(ownerId int, restaId int) (bool, error) {
	SQL := `SELECT id FROM restaurants WHERE user_id=$1`
	var resID []int
	err := database.RMS.Select(&resID, SQL, ownerId)
	if err != nil {
		return false, err
	}
	for _, v := range resID {
		if v == restaId {
			return true, nil
		}
	}
	return false, nil
}
