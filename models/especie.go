/*
 * @File: models.token.go
 * @Description: Struct da tabela "pmi_especie_bem_movel"
 * @Author: Carlos Henrique Lemos (chenriquelemos@gmail.com)
 */
package models

type Especie struct {
	ID                     uint   `gorm:"column:id_especie_bem_movel;primary_key"`
	Descricao              string `gorm:"column:nome_especie_bem_movel"`
	ChassiObrigatorio      bool   `gorm:"column:indi_chassi_obrig"`
	NumeroSerieObrigatorio bool   `gorm:"column:indi_serie_obrig"`
	ClasseID               uint   `gorm:"column:id_classe_bem_movel"`
	Classe                 Classe
	EspecieSigilosa        bool `gorm:"column:indi_especie_sigilosa"`
}

func (e Especie) TableName() string {
	return "pmi_especie_bem_movel"
}
