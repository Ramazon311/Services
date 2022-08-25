package main 

import (
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err!= nil{
		fmt.Println(err)
		return
	}
	name := r.FormValue("name")
	fmt.Println(w,name)
	addres:= r.FormValue("addres")
	fmt.Println(w,addres)
}
func helloHandler(w http.ResponseWriter,r *http.Request){
    if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method !="GET"{
		http.Error(w,"missing method GET",http.StatusNotFound)
		return
	}
	fmt.Println(w,"hello")
}
func main(){
   fileServer := http.FileServer(http.Dir("./static"))
   http.Handle("/",fileServer)
   http.HandleFunc("/form",formHandler)
   http.HandleFunc("/hello",helloHandler)
   fmt.Println("Starting server at port 8080")
   err:=http.ListenAndServe(":8080",nil)
   if err != nil{
	log.Fatal(err)
   }
}