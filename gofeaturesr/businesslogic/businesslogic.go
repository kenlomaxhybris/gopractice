package businesslogic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = []Article{
	Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func homePage(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}

func returnAllArticles(w http.ResponseWriter, q *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	//	fmt.Fprintf(w, "%+v", string(reqBody))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, _ := strconv.Atoi(id)
	// we then need to loop through all our articles
	for _, article := range Articles {
		// if our id path parameter matches one of our
		// articles

		if article.Id == i {
			// updates our Articles array to remove the
			// article

		}
	}
}

var myRouter *mux.Router

func SetUpRouterServer() {
	myRouter = mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
}

func Serve() {
	//http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}
