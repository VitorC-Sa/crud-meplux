package clientmodel

import (
	"time"

	"github.com/VitorC-sa/crud-meplux/database"
	"gorm.io/gorm"
)

type Client struct {
	Id              uint      `json:"id" gorm:"primarykey,autoincrement"`
	Nome            string    `json:"nome" gorm:"not null"`
	Cnpj            string    `json:"cnpj" gorm:"not null"`
	Endereco        string    `json:"endereco,omitempty"`
	Telefone        string    `json:"telefone,omitempty"`
	Email           string    `json:"email,omitempty"`
	DataCriacao     time.Time `json:"data_criacao" gorm:"autoCreateTime:milli"`
	DataAtualizacao time.Time `json:"data_atualizacao" gorm:"autoUpdateTime:milli"`
}

var db *gorm.DB

func init() {
	db = database.GetConnection()
	db.AutoMigrate(&Client{})
}

func GetOne(id int) *Client {
	panic("Not implemented")
}

func GetAll() []Client {
	panic("Not implemented")
}

func Insert(cli Client) {
	panic("Not implemented")
}

func Update(cli Client) {
	panic("Not implemented")
}
