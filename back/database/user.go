package database

import (
	"fmt"
	"google_auth/models"
)

func GetUser(field string, value interface{}) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf(`SELECT id, created_at, updated_at, full_name, email, telephone, password, 
		(SELECT EXISTS (SELECT * FROM google_profiles WHERE user_id = users.id)) as is_googler FROM users WHERE %s = ?`, field)
	row := DB.QueryRow(query, value)

	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.FullName, &user.Email, &user.Telephone, &user.Password, &user.IsGoogler)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByGoogleUID(uid string) (*models.User, error) {
	var user models.User

	query := `SELECT id, created_at, updated_at, full_name, email, telephone, password 
	FROM users WHERE id = (SELECT user_id FROM google_profiles WHERE uid = ?)`
	row := DB.QueryRow(query, uid)

	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.FullName, &user.Email, &user.Telephone, &user.Password)
	if err != nil {
		return nil, err
	}
	user.IsGoogler = true

	return &user, nil
}

func CreateUser(user *models.User) error {
	query := `INSERT INTO users (created_at, updated_at, full_name, email, telephone, password) VALUES (NOW(), NOW(), ?, ?, ?, ?)`
	res, err := DB.Exec(query, user.FullName, user.Email, user.Telephone, user.Password)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	user.ID = uint(id)
	return err
}

// here is no error if not found
func UpdateUser(user *models.User) error {
	query := `
		UPDATE users
		SET updated_at = NOW(), full_name = ?, email = ?, telephone = ?
		WHERE id = ?
	`
	_, err := DB.Exec(query, user.FullName, user.Email, user.Telephone, user.ID)

	return err
}
