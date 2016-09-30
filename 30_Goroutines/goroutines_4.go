package main

import (
	"io"
	"log"
	"net" //Nos permite utilizar el http para crear servidores webs
	"time"
)

func main() {
	/*Se crea un listener escucha en la interfaz a que haya nuevas conexiones*/
	listener, err := net.Listen("tcp", "localhost:8000")
	//Se verifica si hubo error
	if err != nil {
		log.Fatal(err)
	}
	for {
		//Devuelve una conexión y un error
		conn, err := listener.Accept()
		if err != nil {
			/*Este error es de la conexión entrante */
			log.Print(err) //Por ejemplo conexion abortada
			continue
		}
		//se manejan las conexiones de manera concurrente
		//go manejarCon(conn)
	}
}

func manejarCon(c net.Conn) {
	//Siempre se debe cerrar
	defer c.Close()

	for {
		//con Format le decimos como queremos que time devuelva los datos
		_, err := io.WriteString(c, time.Now().Format("15:04:05\r"))
		if err != nil {
			return //ejemplo si el cliente se desconecta
		}
		time.Sleep(1 * time.Second)
	}

}
