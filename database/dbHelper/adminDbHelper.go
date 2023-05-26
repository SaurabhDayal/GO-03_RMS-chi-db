package dbHelper

import (
	"06_RMS-chi-db/database"
	"06_RMS-chi-db/errorHandling"
	"06_RMS-chi-db/models"
)

const subAdminCredit = models.SubAdminCredit
const role = "subAdmin"

func CreateNewSubAdmin(subAdm *models.Users) (*models.UsersClient, error) {
	var subAdmin models.UsersClient
	SQL := `INSERT INTO users (user_name, user_password, user_email, credit) VALUES ($1,$2,$3,$4) RETURNING id,user_name,user_password,user_email,credit`
	err := database.RMS.Get(&subAdmin, SQL, subAdm.UserName, subAdm.UserPassword, subAdm.UserEmail, subAdminCredit)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	SQL2 := `INSERT INTO user_roles (user_id, role) VALUES ($1, $2)`
	_, err = database.RMS.Exec(SQL2, subAdmin.Id, role)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	return &subAdmin, nil
}

func GetAllSubAdmin() ([]models.Users, error) {
	subAdmins := make([]models.Users, 0)
	SQL := `SELECT u.user_name, u.user_password, u.user_email, u.credit FROM users u 
            INNER JOIN user_roles r ON u.id = r.user_id
            WHERE r.role = $1 AND u.archived_at IS NULL`
	err := database.RMS.Select(&subAdmins, SQL, role)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	return subAdmins, nil
}