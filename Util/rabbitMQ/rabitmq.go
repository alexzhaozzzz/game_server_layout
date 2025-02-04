package rabbitMQ

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

// 连接信息amqp://kuteng:kuteng@127.0.0.1:5672/kuteng这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host

// rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	MqUrl string
}

// 创建结构体实例
func NewRabbitMQ(mqUrl string, queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: mqUrl}
}

// 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(mqUrl string, queueName string) (*RabbitMQ, error) {
	if mqUrl == "" || queueName == "" {
		return nil, errors.New("param is nil")
	}

	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(mqUrl, queueName, "", "")

	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	if err != nil {
		return nil, err
	}

	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return nil, err
	}

	return rabbitmq, nil
}

// 直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) error {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return err
	}

	//调用channel 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}

	return nil
}

// simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple() (<-chan amqp.Delivery, error) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return nil, err
	}

	//接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		false, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
