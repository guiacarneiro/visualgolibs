/*
 * Created on: 22/01/2019
 *     Author: guilhermehenrique
 */

package messaging

//ReceiveHandlerInterface - interface para executar quando receber mensagem da mensageria
type ReceiveHandlerInterface interface {
	Execute(message []byte)
}
