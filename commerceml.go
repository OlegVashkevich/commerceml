package main

import (
	"commerceml/schema_import"
	"commerceml/schema_offers"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//get godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envLogin := os.Getenv("login")
	envPassword := os.Getenv("password")

	//A - http://<сайт>/<путь> /1c_exchange.php?type=catalog&mode=checkauth.
	if r.URL.Query().Get("type") == "catalog" &&
		r.URL.Query().Get("mode") == "checkauth" {
		//Check Basic HTTP Auth
		username, password, ok := r.BasicAuth()
		if ok && envLogin == username && envPassword == password {
			fmt.Printf("%#v\n", username)
			fmt.Printf("%#v\n", password)
			fmt.Fprintf(w, "success")
		} else {
			fmt.Fprintf(w, "fail")
		}
	}
}

func main() {
	//import0_1
	// Open our xmlFile
	xmlFile, err := os.Open("xml_data/import0_1.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened import0_1.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xmlFile)

	//d := xml.NewDecoder(xmlFile)

	data := new(schema_import.КоммерческаяИнформация)
	err = xml.Unmarshal(byteValue, data)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%#v\n", data.Каталог.Товары.Товар[0])

	//offers0_1
	// Open our xmlFile
	xmlFile2, err := os.Open("xml_data/offers0_1.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened offers0_1.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile2.Close()

	// read our opened xmlFile as a byte array.
	byteValue2, _ := io.ReadAll(xmlFile2)

	//d := xml.NewDecoder(xmlFile)

	data2 := new(schema_offers.КоммерческаяИнформация)
	err = xml.Unmarshal(byteValue2, data2)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%#v\n", data2.ИзмененияПакетаПредложений.Предложения.Предложение[0])

	http.HandleFunc("/exchange1c.php", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
