package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	var config *globals.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func LeerConsola() {
	// Leer de la consola
	var text string
	log.Println("Ingrese los mensajes")
	for text != "\n" {
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		log.Print(text)
	}
}

func GenerarYEnviarPaquete() {
	paquete := Paquete{}

	// Leemos y cargamos el paquete
	var text string
	log.Println("Ingrese los mensajes a enviar")
	for text != "\n" {
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		log.Print(text)
		last := len(text) - 1
		paquete.Valores = append(paquete.Valores, text[:last], ",")
	}
	log.Printf("Paquete a enviar: %+v", paquete)

	// Enviamos el paquete

	EnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, paquete)
}

func GenerarYEnviarPaquete2() { // Envia las lineas por separado

	// Leemos y cargamos el paquete
	var text string
	log.Println("Ingrese los mensajes a enviar")
	for text != "\n" {
		// paquete := Paquete{}
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		log.Print(text)
		last := len(text) - 1
		// paquete.Valores = append(paquete.Valores, text[:last])
		EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, text[:last])
	}
	// log.Printf("Paquete a enviar: %+v", paquete)

	// Enviamos el paquete

}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
