/*
 * @File: models.token.go
 * @Description: Define quais informações de erro serão retornadas aos clientes
 * @Author: Carlos Henrique Lemos (chenriquelemos@gmail.com)
 */
package models

// Message defines the response message
type Message struct {
	Message string `json:"mensagem" exemplo:"mensagem"`
}
