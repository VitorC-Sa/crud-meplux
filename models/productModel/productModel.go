package productmodel

import (
	"time"

	"github.com/VitorC-sa/crud-meplux/database"
	"gorm.io/gorm"
)

type Product struct {
	Id              uint      `json:"id" gorm:"primarykey,autoincrement"`
	Descricao       string    `json:"descricao"`
	Preco           float64   `json:"preco"`
	Quantidade      uint      `json:"quantidade"`
	DataCriacao     time.Time `json:"data_criacao" gorm:"autoCreateTime:milli"`
	DataAtualizacao time.Time `json:"data_atualizacao" gorm:"autoUpdateTime:milli"`
}

var db *gorm.DB

func init() {
	db = database.GetConnection()
	db.AutoMigrate(&Product{})
}

func GetOne(id int) *Product {
	panic("Not implemented")
}

func GetAll() []Product {
	panic("Not implemented")
}

func Insert(cli Product) {
	panic("Not implemented")
}

func Update(cli Product) {
	panic("Not implemented")
}
