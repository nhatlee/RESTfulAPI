package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Id int `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["key"]
	var1 := vars["var1"]
	var2 := vars["var2"]

	fmt.Println("Var 1: " + var1)
	fmt.Println("Var 2: " + var2)
	fmt.Fprintf(w, "Key: " + key)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage of Nhatlee!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/all", returnAllArticles)
	//log.Fatal(http.ListenAndServe(":8081", nil))
	//myRouter := mux.NewRouter().StrictSlash(true)
	//myRouter.HandleFunc("/", homePage)
	//myRouter.HandleFunc("/all", returnAllArticles)
	//myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	//log.Fatal(http.ListenAndServe(":10000", myRouter))
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{key}/{var1}/{var2}/", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}



func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
