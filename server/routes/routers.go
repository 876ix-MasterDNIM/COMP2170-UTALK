package routes

import (
	"U-Talk/server"
	"U-Talk/server/authenticator"
	"U-Talk/server/repository"
	"U-Talk/server/utilities/htmlserver"
	"U-Talk/server/utilities/sessions"
	"html/template"
	"log"

	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/securecookie"
)

const (
	get  = "GET"
	post = "POST"
)

var db = repository.Repository{}
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// RunServer starts server
func RunServer() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../../public")))
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/threads", threads)
	mux.HandleFunc("/posts", posts)
	mux.HandleFunc("/categories", categories)
	http.ListenAndServe(":3000", mux)
}

func index(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case get:
		htmlserver.ServeHTML("../../public/views/index.html", response)
	}
}

func register(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		db.Repository("u-talk", "users")
		request.ParseMultipartForm(32 << 20)
		file, handler, err := request.FormFile("img")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(response, "%v", handler.Header)
		f, err := os.OpenFile("../../public/uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		user := new(datastructures.User)
		user.User(strings.Join(request.Form["email"], ""), strings.Join(request.Form["password"], ""), strings.Join(request.Form["username"], ""), "../../public/uploads/"+handler.Filename)
		db.StoreUser(user)
	}
}

func login(response http.ResponseWriter, request *http.Request) {
	db.Repository("u-talk", "users")
	username := request.FormValue("username")
	password := request.FormValue("password")

	redirectTarget := "/index"
	err := authenticator.AuthenticateUser(username, password)
	if err != nil {
		fmt.Println(err)
		response.Write([]byte("User does not exist."))
	} else {
		sessions.SetSession(username, response)
		redirectTarget = "/categories"
		http.Redirect(response, request, redirectTarget, 302)
	}

}

func logout(response http.ResponseWriter, request *http.Request) {
	sessions.ClearSession(response)
	http.Redirect(response, request, "/", 302)
}

func categories(response http.ResponseWriter, request *http.Request) {
	htmlserver.ServeHTML("../../public/views/categories.html", response)
}

func threads(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case get:
		// userName := sessions.UserName(request)
		// if userName != "" {
		// 	htmlserver.ServeHTML("../../public/views/threads.html", response)
		// } else {
		// 	http.Redirect(response, request, "/index", 302)
		// }
		//htmlserver.ServeHTML("../../public/views/threads.html", response)
		categoryName := request.FormValue("category")
		data := struct {
			Category string
			Threads  []repository.DbThread
		}{
			categoryName,
			db.Threads(categoryName),
		}

		tpl, err := template.ParseFiles("../../public/views/threads.html")
		if err != nil {
			log.Fatal(err)
		}
		templateErr := tpl.Execute(response, data)
		if templateErr != nil {
			log.Fatal(err)
		}

	case post:
		request.ParseMultipartForm(32 << 20)
		file, handler, err := request.FormFile("img")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(response, "%v", handler.Header)
		f, err := os.OpenFile("../../public/uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		desc := request.FormValue("desc")
		topic := request.FormValue("topic")
		categoryName := request.FormValue("category")
		thread := datastructures.Thread{}
		thread.Thread(desc, sessions.UserName(request), handler.Filename, topic)
		db.AddThread(&thread, categoryName, request)
	}

}

func posts(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case get:
		htmlserver.ServeHTML("../../public/views/posts.html", response)
	}
}
