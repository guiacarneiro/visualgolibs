/*
 * Created on: 15/01/2019
 *     Author: guilhermehenrique
 */

package main

import "github.com/guiacarneiro/visualgolibs/logger"

func main() {
	logger.Inicializa(logger.Debug, "C:\\temp\\log", 1)
	logger.LogError.Println("teste")
}
