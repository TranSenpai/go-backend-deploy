package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-backend/internal/common/env"
	"log"
	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

// doc https://www.rabbitmq.com/tutorials/tutorial-six-go
// mẫu code cho client: https://github.com/rabbitmq/rabbitmq-tutorials/blob/main/go/rpc_client.go
func NewRabbitMQ(env *env.Env) *RabbitMQ {
	conn, err := amqp.DialConfig(env.RabbitMQURL, amqp.Config{
		Properties: amqp.Table{
			"connection_name": "go-backend",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ [RABBIT_MQ] Connection To RabbitMQ Successfully")

	return &RabbitMQ{
		Conn: conn,
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func queueDeclareMain(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	// tạo queue nếu chưa có
	// nếu đã có queue, nó sẽ kiểm tra đúng các setting hay không, nếu không đúng => báo lỗi, nếu đúng thì đi tiếp
	return ch.QueueDeclare(
		queueName, // name
		true,      // durability có tồn tại queue hay không khi rabbitmq restart
		false,     // delete when unused, vì là queue chính cho ên không tự động xoá
		false,     // exclusive độc quyền, có khoá với connection hiện tại hay không, vì là còn B xử lý => false
		false,     // noWait có đợi queue tạo thành công hay không
		nil,       // arguments
	)
}

func queueDeclareTemp(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	// tạo queue nếu chưa có
	// nếu đã có queue, nó sẽ kiểm tra đúng các setting hay không, nếu không đúng => báo lỗi, nếu đúng thì đi tiếp
	return ch.QueueDeclare(
		queueName, // name
		true,      // durability có tồn tại queue hay không khi rabbitmq restart
		true,      // delete when unused, vì là queue tạm cho sẽ phải tự động bí xoá
		false,     // exclusive độc quyền, có khoá với connection hiện tại hay không
		false,     // noWait có đợi queue tạo thành công hay không
		nil,       // arguments
	)
}

func publishNotReply(ctx context.Context, ch *amqp.Channel, queueName string, corrId string, body []byte) error {
	return ch.PublishWithContext(
		ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,

			// với những message quan trọng, thì không chỉ lưu trữ trong RAM
			// giúp khi rabbitmq restart thì message vẫn còn
			// đảm bảo queue sẽ được giữ: durable = true
			DeliveryMode: amqp.Persistent,

			Body: body,
		})
}

func publishReply(ctx context.Context, ch *amqp.Channel, queueName string, replyTo string, corrId string, body []byte) error {
	return ch.PublishWithContext(
		ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			// với những message quan trọng, thì không chỉ lưu trữ trong RAM
			// giúp khi rabbitmq restart thì message vẫn còn
			// đảm bảo queue sẽ được giữ: durable = true
			DeliveryMode: amqp.Persistent,

			ReplyTo: replyTo,
			Body:    body,
		})
}

type Reply struct {
	ErrorString string
	Data        json.RawMessage
}

func (r *RabbitMQ) Send(ctx context.Context, queueName string, payload any) (err error) {
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch, err := r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()

	// tạo queue nếu chưa có
	// nếu đã có queue, nó sẽ kiểm tra đúng các setting hay không, nếu không đúng => báo lỗi, nếu đúng thì đi tiếp
	queueMain, err := queueDeclareMain(ch, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}

	corrId := randomString(32)
	err = publishNotReply(ctx, ch, queueMain.Name, corrId, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (r *RabbitMQ) On(ctx context.Context, queueName string, handler func(context.Context, []byte) error) (err error) {
	ch, err := r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	// tạo queue nếu chưa có
	// nếu đã có queue, nó sẽ kiểm tra đúng các setting hay không, nếu không đúng => báo lỗi, nếu đúng thì đi tiếp
	queueMain, err := queueDeclareMain(ch, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := ch.Consume(
		queueMain.Name, // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		defer ch.Close()
		for d := range msgs {
			// handler
			err := handler(ctx, d.Body)
			if err != nil {
				d.Nack(false, false)
				continue
			}
			d.Ack(false)
		}
	}()

	return
}

func (r *RabbitMQ) Request(ctx context.Context, queueName string, payload any, result any) (err error) {
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch, err := r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()

	// TẠM ==================================
	queueTemp, err := queueDeclareTemp(ch, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs, err := ch.Consume(
		queueTemp.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ==================================

	// CHÍNH ==================================
	// tạo queue nếu chưa có
	// nếu đã có queue, nó sẽ kiểm tra đúng các setting hay không, nếu không đúng => báo lỗi, nếu đúng thì đi tiếp
	queueMain, err := queueDeclareMain(ch, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}

	corrId := randomString(32)
	err = publishReply(ctx, ch, queueMain.Name, queueTemp.Name, corrId, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ==================================

	for {
		select {
		case <-ctx.Done():
			err = errors.New("Request timeout")
			fmt.Println(err)
			return
		case d := <-msgs:
			if d.CorrelationId != corrId {
				continue
			}

			var bodyReply Reply
			err = json.Unmarshal(d.Body, &bodyReply)
			if err != nil {
				fmt.Println(err)
				return
			}

			if bodyReply.ErrorString != "" {
				err = errors.New(bodyReply.ErrorString)
				fmt.Println(err)
				return
			}

			err = json.Unmarshal(bodyReply.Data, result)
			if err != nil {
				fmt.Println(err)
				return
			}

			return
		}
	}
}

func (r *RabbitMQ) OnReply(ctx context.Context, queueName string, handler func(context.Context, []byte) (any, error)) (err error) {
	ch, err := r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	// tạo queue nếu chưa có
	// nếu đã có queue, nó sẽ kiểm tra đúng các setting hay không, nếu không đúng => báo lỗi, nếu đúng thì đi tiếp
	queueMain, err := queueDeclareMain(ch, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := ch.Consume(
		queueMain.Name, // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		defer ch.Close()
		for d := range msgs {
			// hàm gửi lại lỗi cho các logic phái dưới
			replyError := func(replyError error) {
				d.Nack(false, false)

				bodyReply := Reply{
					ErrorString: replyError.Error(),
					Data:        nil,
				}

				body, err := json.Marshal(bodyReply)
				if err != nil {
					fmt.Println(err)
					return
				}

				err = publishNotReply(ctx, ch, d.ReplyTo, d.CorrelationId, body)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			// handler
			result, err := handler(ctx, d.Body)
			if err != nil {
				replyError(err)
				continue
			}

			data, err := json.Marshal(result)
			if err != nil {
				replyError(err)
				continue
			}

			bodyReply := Reply{
				ErrorString: "",
				Data:        data,
			}

			body, err := json.Marshal(bodyReply)
			if err != nil {
				replyError(err)
				continue
			}

			err = publishNotReply(ctx, ch, d.ReplyTo, d.CorrelationId, body)
			if err != nil {
				d.Nack(false, false)
				continue
			}

			d.Ack(false)
		}
	}()

	return
}
