package main

import (
	"log"

	"asdrome.com/csv-to-pb/pocketbase"
)

func main() {
	// Crear una instancia de PocketBase
	pb := pocketbase.NewPocketBase("https://asdrome-pos-pb.pockethost.io").Collection("cuidados_diarios")

	// Datos que deseas enviar a PocketBase
	/* 	data := map[string]interface{}{
		"brand":    "test",
		"image":    "https://example.com",
		"name":     "test",
		"price":    123,
		"discount": 123,
	} */
	var err error

	// Llamar al método create en la instancia de PocketBase
	err = pb.LoadFromCSV("./cuidados-diarios.csv")
	if err != nil {
		// Manejar el error si es necesario
		log.Fatalln(err)
	}
	//log.Println("JSON:", string(data))
}
