/*
 * @File: models.token.go
 * @Description: Struct da tabela "pmi_grupo_bem_movel"
 * @Author: Carlos Henrique Lemos (chenriquelemos@gmail.com)
 */
package models

type Grupo struct {
	ID        uint   `gorm:"column:id_grupo_bem_movel;primary_key"`
	Descricao string `gorm:"column:nome_grupo_bem_movel"`
}

func (g Grupo) TableName() string {
	return "pmi_grupo_bem_movel"
}
