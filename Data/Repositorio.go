package data

import (
	"GoRestApi/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func CriarConexaoBanco() {
	strConn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConn), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar ao DB")
	}
}

func BuscarPersonalidades() []models.Personalidade {
	if DB == nil {
		CriarConexaoBanco()
	}

	var p []models.Personalidade
	DB.Find(&p)

	return p
}

func BuscarPersonalidadePorId(id string) models.Personalidade {
	if DB == nil {
		CriarConexaoBanco()
	}

	var p models.Personalidade
	DB.Where("Id = ?", id).First(&p)

	return p
}

func InserirPersonalidade(p models.Personalidade) models.Personalidade {
	if DB == nil {
		CriarConexaoBanco()
	}

	DB.Create(&p)

	return p
}

func DeletaPersonalidade(id string) {
	if DB == nil {
		CriarConexaoBanco()
	}
	var p models.Personalidade

	DB.Delete(&p, id)
}

func EditaPersonalidade(p models.Personalidade){
	if DB == nil {
		CriarConexaoBanco()
	}

	DB.Save(p)
}
