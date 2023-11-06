package main

import (
	"log"

	"asdrome.com/csv-to-pb/pocketbase"
)

func main() {
	collection := "perfumeria"
	// Crear una instancia de PocketBase
	pb := pocketbase.NewPocketBase("https://asdrome-pos-pb.pockethost.io").Collection(collection)

	// Datos que deseas enviar a PocketBase
	/* 	data := map[string]interface{}{
		"brand":    "test",
		"image":    "https://example.com",
		"name":     "test",
		"price":    123,
		"discount": 123,
	} */

	//_, err := pocketbase.ReadFromCSV("./cuidados-diarios.csv")
	// Llamar al método create en la instancia de PocketBase
	err := pb.LoadFromCSV("./" + collection + ".csv")
	if err != nil {
		// Manejar el error si es necesario
		log.Fatalln(err)
	}
	//log.Println("JSON:", string(data))
}
