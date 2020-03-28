package databases

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"../common"
	"../models"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// PostgresDB manages Postgres connection
type PostgresDB struct {
	Connection   *gorm.DB
	Databasename string
}

func (pg *PostgresDB) Connect() {
	pg.Databasename = common.Config.PgDbName

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		common.Config.PgAddrs,
		common.Config.PgPort,
		common.Config.PgDbUsername,
		pg.Databasename,
		common.Config.PgDbPassword)

	fmt.Println(psqlInfo)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Debug("Não é possível conectar-se ao banco, erro: ", err)
		panic(err)
	}

	// Enable Logger, show detailed log
	db.LogMode(true)
	db.Exec("SET search_path TO pmi")

	var especie models.Especie
	var classe models.Classe

	res := db.Model(&especie).Related(&classe).Find(&especie, &models.Especie{ID: 215856})

	if res.RecordNotFound() {
		fmt.Println(res)
	}

	fmt.Println(especie)
}

// Close the existing connection
func (pg *PostgresDB) Close() {
	if pg.Connection != nil {
		pg.Connection.Close()
	}
}
