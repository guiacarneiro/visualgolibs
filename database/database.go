/*
 * Created on: 29/10/2018
 *     Author: guilhermehenrique
 */

package database

import (
	"regexp"

	"github.com/jmoiron/sqlx"
)

var stringConexaoBancoDados = ""

//Inicializa - inicializa configuracoes
func Inicializa(stringConexao string) {
	stringConexaoBancoDados = stringConexao
}

//CriaConexao - cria conexao com o banco de dados
func CriaConexao() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", stringConexaoBancoDados)
	return db, err
}

//VerificaErrorNoResult
func VerificaErrorNoResult(err error) bool {
	if err != nil && err.Error() == "sql: no rows in result set" {
		return true
	} else {
		return false
	}
}

//VerificaErrorTableNotFound
func VerificaErrorTableNotFound(err error) bool {
	if err != nil {
		matched, _ := regexp.MatchString("Error \\d*: Table '.*' doesn't exist", err.Error())
		if matched {
			return true
		}
	}
	return false
}
