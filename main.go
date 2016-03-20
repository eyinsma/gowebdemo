package main
import (
	"net/http"
	"fmt"
	"strings"
	"log"
"html/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse the parameter
	fmt.Println(r.Form)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ","))
	}

	fmt.Fprintf(w, "hello, world!")
}


func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	//simple
	http.HandleFunc("/", sayHello)
	//add login
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}