package main

import "net/http"
import "encoding/json"
import "fmt"
import "contact"

func main() {

	http.HandleFunc("/hello",hello)
	http.HandleFunc("/contacts/save",saveContact)
	http.HandleFunc("/contacts/get",getContact)
	http.ListenAndServe(":9000",nil)
}


func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello goleng"))
}
func saveContact(w http.ResponseWriter, r *http.Request) {
	var res map[string]string
	json.NewDecoder(r.Body).Decode(&res)
	
	// fmt.Println(res)

	contact.Add(res["name"],contact.Contact{
		Tel		: res["tel"],
		Email	: res["mail"],
	})

	fmt.Println(contact.Get(res["name"]))

}

func getContact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		name := r.URL.Query().Get("name")
		c, _ := contact.Get(name)
		json.NewEncoder(w).Encode(c)
	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}