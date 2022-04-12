/*
 * Created on: 21/01/2019
 *     Author: guilhermehenrique
 */

package messaging

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"visualgolibs/logger"
)

//ReceiveMessage - receive message
func ReceiveMessage(receiveHandlerInterface ReceiveHandlerInterface, routingKey string, exchange string) error {
	opts := mqtt.NewClientOptions().AddBroker(strings.Join([]string{"ws://", ipMensageria, "/ws"}, ""))
	opts.SetClientID(exchange)
	opts.AutoReconnect = true
	opts.OnConnect = func(c mqtt.Client) {
		c.Subscribe(routingKey, 0, func(client mqtt.Client, msg mqtt.Message) {
			logger.LogInfo.Println("Mensagem recebida: ", string(msg.Payload()))
			receiveHandlerInterface.Execute(msg.Payload())
		})
	}

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		logger.LogError.Println("Failed to connect to RabbitMQ", token.Error())
		return token.Error()
	}

	return nil
}
