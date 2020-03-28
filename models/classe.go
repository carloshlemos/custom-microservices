/*
 * @File: models.token.go
 * @Description: Struct da tabela "pmi_classe_bem_movel"
 * @Author: Carlos Henrique Lemos (chenriquelemos@gmail.com)
 */
package models

type Classe struct {
	ID        uint   `gorm:"column:id_classe_bem_movel;primary_key"`
	Descricao string `gorm:"column:nome_classe_bem_movel"`
	GrupoID   uint   `gorm:"column:id_grupo_bem_movel"`
	Grupo     Grupo
}

func (c Classe) TableName() string {
	return "pmi_classe_bem_movel"
}
