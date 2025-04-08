package main

import (
	"dictionaries/cmd/dictionaries"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	port := 8080
	http.HandleFunc("/", handleForm)
	http.HandleFunc("/calculate", handleCalculation)
	log.Printf("Server started on :%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/templates/form.html"))
	tmpl.Execute(w, nil)
}

func handleCalculation(w http.ResponseWriter, r *http.Request) {
	// Обработка данных формы
	r.ParseForm()
	birthDate := r.Form.Get("birthdate")
	birthTime := r.Form.Get("birthtime")
	location := r.Form.Get("location")
	userName := r.Form.Get("name")

	fmt.Println(birthDate, birthTime, location, userName)

	inputAdress := dictionaries.Adress{
		Id:      1,
		Country: location,
		//UpdatedDateTime: time.Now(),
	}

	inputBodigraf := dictionaries.Bodigraf{
		Id:               1,
		BirthDateTimeUtc: birthDate + birthTime,
		Agress:           inputAdress,
		//CreatedDateTime:  time.Now(),
	}

	fmt.Println(inputBodigraf)
}
