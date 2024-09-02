package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	log.Println("Hola soy un log") // (TOMI) loggea en el archivo creado e imprime por pantalla

	// ======= loggear "Hola soy un log" usando la biblioteca log

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// ======= validar que la config este cargada correctamente

	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración") // (TOMI) termina el proceso si no puede leer el config
	}

	// ======= loggeamos el valor de la config

	log.Println(globals.ClientConfig.Mensaje) // (TOMI) logeo el mensaje leido del archivo de config.

	// ======= leemos consola

	// utils.LeerConsola() // (TOMI) lee la consola (en principio, una sola vez)

	// ======= ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// ======= enviar un mensaje al servidor con el valor de la config

	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// ======= leer de la consola el mensaje

	utils.LeerConsola()

	// ======= generamos un paquete y lo enviamos al servidor

	utils.GenerarYEnviarPaquete()
}
