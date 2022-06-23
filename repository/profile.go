package repository

import "database/sql"

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

// func (p *ProfileRepository) GetProfile() ([]User, error) {
// 	var profiles []User

// 	rows, err := p.db.Query("SELECT * FROM users WHERE email = ?")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var profile User
// 		err := rows.Scan(&profile.Nama, &profile.Email)
// 		if err != nil {
// 			return nil, err
// 		}
// 		profiles = append(profiles, profile)
// 	}

// 	return profiles, nil
// }

func (p *ProfileRepository) FetchProfileByID(id int64) (User, error) {
	var profile User
	err := p.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&profile.ID, &profile.Nama, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return profile, err
	}

	return profile, nil
}
