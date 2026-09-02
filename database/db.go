package database

import (
	"log"
<<<<<<< HEAD
	"os"
=======
>>>>>>> 6eeb98f6a68e30563c9ab9fd373e4c18d158ae56

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
<<<<<<< HEAD
	stringDeConexao := "host=" + os.Getenv("HOST") + " user=" + os.Getenv("USER") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("PORT") + " sslmode=disable"
=======
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
>>>>>>> 6eeb98f6a68e30563c9ab9fd373e4c18d158ae56
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
