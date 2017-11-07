package main

import (
	"github.com/Sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v3"
	elogrus "gopkg.in/sohlich/elogrus.v1"
	//"os"
	"log"
	"net/http"
	"time"
	//"gopkg.in/olivere/elastic.v3"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type cuerpo struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	// url := r.URL.Path
	params := mux.Vars(r)
	tabla := params["tabla"]
	id := params["id"]
	//strB, _ := json.Marshal("gopher")

	//Leer Cuerpo
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg *cuerpo
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(msg)

	//println(r.Body["Nombre"]);

	Insertar_elastic(tabla+" id:"+id, r.Method, msg)
}

func exec_cliente_elastic() {

	//Configurando el cliente del Elastic Search
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://10.20.2.111:9200"))
	if err != nil {
		panic("error elasticsearch " + err.Error())
	}
	println("cliente elastic creado")
	println(client)

	//Se crea un hook
	hook, err := elogrus.NewElasticHook(client, "loghost", logrus.DebugLevel, "logstash")
	if err != nil {
		panic("error elogrus" + err.Error())
	}
	println("hook elogrus creado")

	//Se agrega un hook al logrus
	logrus.AddHook(hook)

	//Muxer
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/{tabla:[a-zA-Z0-9_]+}/{id:[a-zA-Z0-9_]+}", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func Insertar_elastic(usuario string, peticion string, data *cuerpo) {
	now := time.Now().String()
	logrus.WithFields(logrus.Fields{
		"Fecha_operacion": now,
		"Usuario":         usuario,
		"Ip":              "10.20.0.256",
		"Petición":        peticion,
		"Data":            data,
	}).Info("Now!")
}
