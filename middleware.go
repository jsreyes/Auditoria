package Auditoria

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	amqp "github.com/streadway/amqp"
)

//Variables para la conexión y el canal
//var connection *amqp.Connection
//var chl *amqp.Channel

type Auditoria struct {
	User   string
	Metodo string
	Ip     string
}

/*func failOnError(err error, msg string) {
	if err != nil {
		beego.Info("%s: %s", msg, err)
		beego.Info(fmt.Sprintf("%s: %s", msg, err))
	}
}*/

func FunctionBeforeStatic(ctx *context.Context) {
	beego.Info("beego.BeforeStatic: Before finding the static file")
}
func FunctionBeforeRouter(ctx *context.Context) {
	beego.Info("beego.BeforeRouter: Executing Before finding router")
}
func FunctionBeforeExec(ctx *context.Context) {

	beego.Info("beego.BeforeExec: After finding router and before executing the matched Controller")
}

func FunctionAfterExec(ctx *context.Context) {
	//beego.config("appname")
	fmt.Println(reflect.TypeOf(ctx.Request))
	fmt.Printf("%+v\n", ctx.Request)
	fmt.Println(ctx.Request.URL)
	fmt.Println(ctx.Request.Method)
	fmt.Println(ctx.Request.Body)

	//fmt.Println(ctx.Request.RemoteAddr)

	beego.Info("beego.AfterExec: After executing Controller")

	var mensaje = fmt.Sprintf("{'User': 'userWSO2', 'Metodo': '%s', 'IP': '10.20.0.15', 'Body':'%s'}", string(ctx.Request.Method), ctx.Request.Body)

	sentToRabbit(mensaje)

}

func sentToRabbit(msj string) {

	//p := beego.AppConfig.Strings("RABBIT_MQ_URI")

	//Obtengo el parametro de configuración del API

	con, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer con.Close()

	connection = con

	chanel, err := con.Channel()
	failOnError(err, "Failed to open a channel")
	defer chanel.Close()

	chl = chanel

	cha := beego.AppConfig.Strings("RABBIT_MQ_CHANNEL")

	q, err := chl.QueueDeclare(
		cha[0], // name
		false,  // durable
		false,  // delete when usused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	beego.Info(q)
	failOnError(err, "Failed to declare a queue")

	body := msj
	err = chl.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
func FunctionFinishRouter(ctx *context.Context) {
	beego.Info("beego.FinishRouter: After finishing router")
}

func InitMiddleware() {

	//beego.InsertFilter("*", beego.BeforeStatic, FunctionBeforeStatic, false)
	//beego.InsertFilter("*", beego.BeforeRouter, FunctionBeforeRouter, false)
	//beego.InsertFilter("*", beego.BeforeExec, FunctionBeforeExec, false)
	beego.InsertFilter("*", beego.AfterExec, FunctionAfterExec, false)
	//beego.InsertFilter("*", beego.FinishRouter, FunctionFinishRouter, false)
}
