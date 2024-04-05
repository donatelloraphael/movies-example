package data

import "database/sql"

type Models struct {
	Users  UserModel
	Tokens TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:  UserModel{DB: db},
		Tokens: TokenModel{DB: db},
	}
}
