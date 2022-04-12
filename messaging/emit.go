/*
 * Created on: 07/12/2018
 *     Author: guilhermehenrique
 */

package messaging

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"visualgolibs/logger"
)

//SendMessage - send message
func SendMessage(routingKey string, exchange string, message string) error {
	opts := mqtt.NewClientOptions().AddBroker(strings.Join([]string{"ws://", ipMensageria, "/ws"}, ""))
	opts.SetClientID(exchange)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		logger.LogError.Println("Failed to connect to RabbitMQ", token.Error())
		return token.Error()
	}

	token := c.Publish(routingKey, 0, false, message)
	token.Wait()

	c.Disconnect(250)

	return nil
}