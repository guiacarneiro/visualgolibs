/*
 * Created on: 23/01/2019
 *     Author: guilhermehenrique
*/

package messaging

var (
	ipMensageria string
	exchange string
)

func Inicializa(ipMessaging string, nomeExchange string) {
	ipMensageria = ipMessaging
	exchange = nomeExchange
}