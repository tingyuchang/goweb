package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func AddAuth(username, password string) bool {
	db.Create(&Auth{
		Username: username,
		Password: password,
	})

	return true
}

func ExistAuthByUsername(username string) bool {
	var auth Auth
	db.Select("username").Where("username = ?", username).First(&auth)

	if auth.Username != "" {
		return true
	}

	return false
}
