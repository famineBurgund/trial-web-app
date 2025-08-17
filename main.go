package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html")) //Вводим прерывание при ошибках

func indexHandler(w http.ResponseWriter, r *http.Request) { //По факту оставляем как есть, только через библиотеку html/template
	tpl.Execute(w, nil) //Сейчас нет файла для создания шаблона,поэтому пока что nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1488"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
