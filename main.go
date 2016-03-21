package main
import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
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

// enctype
// parseMultipartForm
// FormFile

func upload(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET") {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("static/upload.gtpl")
		t.Execute(w, token)
	}else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("browsefile")
		if err != nil {
			fmt.Println("formfile error",err)
			return
		}
		defer file.Close()

		f, err := os.OpenFile("C:/EYINSMA/home/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("openfile error",err)
			return
		}
		defer f.Close()

		io.Copy(f, file)

	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/login.gtpl")
		t.Execute(w, nil)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}