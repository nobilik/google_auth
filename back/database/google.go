package database

import "google_auth/models"

func CreateGoogleProfile(profile *models.GoogleProfile) error {
	query := `INSERT INTO google_profiles (created_at, updated_at, user_id, uid) VALUES (NOW(), NOW(), ?, ?)`
	res, err := DB.Exec(query, profile.UserID, profile.UID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	profile.ID = uint(id)
	return err
}
