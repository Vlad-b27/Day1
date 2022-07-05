package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		router.Get("/Person", GetPerson)
		router.Post("/Person", CreatePerson)
		http.ListenAndServe(":8082", router)
	})
	router.Route("/v2", func(r chi.Router) {
		router.Get("/Person", GetPerson)
		router.Post("/Person", CreatePerson)
		http.ListenAndServe(":8082", router)
	})
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	type person struct {
		name :"Name",
		gender:"Male",
		age   :1,
	}
    value := fmt.Sprintf("Person is %v\n", person)
	fmt.Fprintf(w,Person)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
 body:=r.Body
 decodedBody,_:=io.ReadAll(body)
 json.Unmarshal(decodedBody,&person)
 value := fmt.Sprintf("Person is %v\n", person)
 fmt.Fprintf(w,Person)
}
