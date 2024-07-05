// user.go
package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int       `json:"id" db:"id" pk:"s"`
	Pessoa     int       `json:"pessoa" db:"pessoa"`
	PessoaNome string    `json:"nome" db:"pessoa_nome" exclude:"s"` //isso garante que o campo seja retirado de insert e update
	Login      string    `json:"login" db:"login" validate:"required,email"`
	Perfil     int       `json:"perfil" db:"perfil"`
	Acesso     time.Time `json:"acesso" exclude:"s"`
	Ativo      string    `json:"ativo" db:"ativo" validate:"required"`
}

func scanUser(rows *sql.Rows) (User, error) {
	var user User
	err := rows.Scan(&user.Id, &user.Pessoa, &user.Login, &user.Ativo)
	return user, err
}

func userValues(user User) []interface{} {
	return []interface{}{user.Id, user.Pessoa, user.Ativo}
}

func userIDValue(user User) interface{} {
	return user.Id
}
