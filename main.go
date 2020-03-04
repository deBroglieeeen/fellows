package main

import(
	"fmt"
	"net/http"
)

// func handler(writer http.ResponseWriter,request *http.Request){
// 	fmt.Fprintf(writer,"Hello boy, %s!",request.URL.Path[1:])
// }

// func main(){
// 	http.HandleFunc("/",handler)
// 	http.ListenAndServe(":8080",nil)
// }

func main(){
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/",http.StripPrefix("/static/",files))

	mux.HandleFunc("/",index)
	mux.HandleFunc("/err",err)
	mux.HandleFunc("/login",login)
	mux.HandleFunc("/logout",logout)
	mux.HandleFunc("/signup",signup)
	mux.HandleFunc("/signup_account",signup_account)
	mux.HandleFunc("/authenticate",authenticate)

	mux.HandleFunc("/thread/new",newThread)
	mux.HandleFunc("/thread/create",createThread)
	mux.HandkeFunc("/thread/post",postThread)
	mux.HandleFunc("/thread/read",readThread)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter,r * http.Request){
	files := []string{"templates/layout.html",
										"templates/navbar.html",
										"templates/index.html",}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads(); if err == nil{
		templates.ExecuteTemplate(w,"layout",threads)
	}
}