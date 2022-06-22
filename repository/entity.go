package repository

type User struct {
	ID        int     `json:"id"          db:"id"`
	Nama      string  `json:"nama"        db:"nama"`
	Email     string  `json:"email"       db:"email"`
	Password  string  `json:"password"    db:"password"`
	CreatedAt *string `json:"created_at"  db:"created_at"`
	UpdatedAt *string `json:"updated_at"  db:"updated_at"`
}

type Biodata struct {
	ID_Biodata    int    `json:"id_biodata"    db:"id_biodata"`
	Nama          string `json:"nama"          db:"nama"`
	Jenis_Kelamin string `json:"jenis_kelamin" db:"jenis_kelamin"`
	No_HP         string `json:"no_hp"         db:"no_hp"`
	Alamat        string `json:"alamat"        db:"alamat"`
}
