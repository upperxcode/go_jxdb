package main

import (
	"fmt"
	"log"

	"github.com/upperxcode/go_jxdb/jxdb"
	"github.com/upperxcode/go_jxdb/pkg/models"
)

func main() {
	// Configuração do banco de dados
	driver := jxdb.Postgres
	host := "172.18.0.2"
	user := "postgres"
	dbname := "upper"
	password := "postgres"
	port := 5432

	DB, err := jxdb.InitInstance(driver, host, user, dbname, password, port)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	// Configuração do repositório genérico para usuários
	userRepo := &jxdb.GenRepository[models.User]{
		Db:        DB.Conn.DB,
		TableName: "pessoa.usuario",
		Fields:    []string{"pessoa.usuario.id", "pessoa.usuario.pessoa", "p pessoa_nome", "pessoa.usuario.login", "pessoa.usuario.perfil", "pessoa.usuario.acesso", "pessoa.usuario.ativo"},
		Order:     "ORDER BY id",
		Limit:     250,
	}
	userRepo.Joins = append(userRepo.Joins, "LEFT JOIN pessoa.pessoa p ON usuario.pessoa = p.id")

	users, err := userRepo.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users:", users)

}
