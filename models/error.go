/*
 * @File: models.token.go
 * @Description: Define quais informações de erro serão retornadas aos clientes
 * @Author: Carlos Henrique Lemos (chenriquelemos@gmail.com)
 */
package models

type Error struct {
	Code    int    `json:"código" exemplo:"27"`
	Message string `json:"mensagem" exemplo:"Mensagem de Erro"`
}
