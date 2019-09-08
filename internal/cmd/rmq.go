package main

import (
	"os"

	"rmq/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

var (
	Revision = "unknown"
	Version  = "unknown"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	c, err := config.Load("config.yml")
	if err != nil {
		log.Panic().Err(err).Msg("can't load config")
	}

	log.Info().
		Str("version", Version).
		Str("revision", Revision).
		Int("pid", os.Getpid()).
		Interface("config", c).
		Msg("started")

	conn, err := amqp.Dial(c.MQ.URL)
	if err != nil {
		log.Panic().Err(err).Msg("rabbitmq connection")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic().Err(err).Msg("rabbitmq channel")
	}

	q, err := ch.QueueDeclare(
		"test",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Panic().Err(err).Msg("rabbitmq queue")
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello world"),
		})

}
