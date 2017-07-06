package Auditoria

import (
	//Se importan las librerias necesarias para que el Elastic funcione
	"github.com/Sirupsen/logrus"
	//"gopkg.in/olivere/elastic.v3"
	//"gopkg.in/sohlich/elogrus.v1"
	//"os"
	"time"
)

// func main() {
// 	//Se coloca un tick para observar notificaciones al correr el programa
// 	t := time.Tick(2 * time.Second)
//
// 	//Imprime la dirección del Elastic, ya que fue encapsulada como Variable de Entorno
// 	println(os.Getenv("ELASTICSEARCH_URL"))
// 	println("Imprimio la url")
//
// 	//Configurando el cliente del Elastic Search
// 	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
// 	if err != nil {
// 		panic("error elasticsearch " + err.Error())
// 	}
// 	println("cliente elastic creado")
// 	println(client)
//
// 	//Se crea un hook
// 	hook, err := elogrus.NewElasticHook(client, "loghost", logrus.DebugLevel, "logstash")
// 	if err != nil {
// 		panic("error elogrus" + err.Error())
// 	}
// 	println("hook elogrus creado")
//
// 	//Se agrega un hook al logrus
// 	logrus.AddHook(hook)
// }

//Se crea un bucle for, teniendo en cuenta el tick que se realiza al principio, y muestra la fecha y la hora actual
func insertar_elastic2(usuario string, peticion string) {
	now := time.Now().String()
	logrus.WithFields(logrus.Fields{
		"Fecha_operacion": now,
		"Usuario":         "prueba",
		"Ip":              "10.20.0.256",
		"Valor":           "test",
		"Petición":        "GET",
	}).Info("Now!")
}
