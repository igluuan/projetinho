package models

//Importação do pacote time para manipulação de datas e horas
import (
	"log"
	"time"
)

var (
	Timezone   *time.Location
	DateLayout = "02/01/2006"
)

func init() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatal("Erro no timezone", err)
	}
	Timezone = loc
}

// Definição da estrutura User com campos ID, Name, Email e CreatedAT
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Declaração de uma variável Users que é uma fatia de User, inicializada com alguns dados de exemplo(Banco de dados local)
var Users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com", CreatedAt: time.Now()},
	{ID: 2, Name: "Bob", Email: "bob@example.com", CreatedAt: time.Now()},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com", CreatedAt: time.Now()},
}

func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.In(Timezone).Format(DateLayout)
}
