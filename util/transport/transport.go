package transport

import (
	"github.com/golang/protobuf/proto"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/gen"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Transport interface {
	Connect() error

	StartConsume() error

	StartProduce() error

	Close() error
}

type AMQPTransport struct {
	ConsumeChannel chan<- *PacketContainer
	ProduceChannel chan *PacketContainer

	config  Config
	channel *amqp.Channel
	queue   *amqp.Queue
}

func (t *AMQPTransport) Connect() error {
	if t.config.UseTLS {
		// TODO: enable TLS
		return nil
	}

	conn, err := amqp.Dial(t.config.String())

	if err != nil {
		return err
	}

	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		t.config.ExtraData["queue_id"],
		false,
		false,
		true,
		false,
		nil,
	)

	err = ch.ExchangeDeclare(
		"luxor_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	t.channel = ch
	t.queue = &q
	return nil
}

func (t *AMQPTransport) StartConsume() error {
	msgs, err := t.channel.Consume(
		t.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	t.ConsumeChannel = make(chan<- *PacketContainer)

	go func() {
		for d := range msgs {
			var data *raft.Packet

			if err := proto.Unmarshal(d.Body, data); err != nil {
				log.Println(err)
				continue
			}

			t.ConsumeChannel <- &PacketContainer{
				ReplyTo:  d.ReplyTo,
				Packet:   data,
				Received: time.Now().UnixNano() / 1000,
			}
		}
	}()

	return nil
}

func (t *AMQPTransport) StartProduce() error {
	t.ProduceChannel = make(chan *PacketContainer)
	go func() {
		for pc := range t.ProduceChannel {
			data, err := proto.Marshal(pc.Packet)

			if err != nil {
				log.Println(err)
				continue
			}

			err = t.channel.Publish(
				"luxor_topic",
				pc.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "application/protobuf",
					CorrelationId: uuid.NewV4().String(),
					Body:          data,
				})

			if err != nil {
				log.Println(err)
				continue
			}
		}
	}()
	return nil
}

func (t *AMQPTransport) Close() error {
	close(t.ProduceChannel)
	close(t.ConsumeChannel)
	_, err := t.channel.QueueDelete(t.queue.Name, false, false, false)

	if err != nil {
		return err
	}

	return t.channel.Close()
}