package kchat
 
import (
  "appengine"
  "appengine/datastore"

  "encoding/json"
  "html/template"
  "net/http"
  "time"
  "fmt"
)

type Message struct {
  Name  string
  Body  string 
  Date  time.Time
}

func init() {
  http.HandleFunc("/", root)
  http.HandleFunc("/chat", chat)
}

func root(w http.ResponseWriter, r *http.Request) {
  var mainTemplate = template.Must(template.ParseFiles("main.html"))

  if err := mainTemplate.Execute(w, map[string] string {}); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func chat(w http.ResponseWriter, r *http.Request) {
  context := appengine.NewContext(r)

  if r.Method == "GET" {
    q := datastore.NewQuery("Message").Order("-Date").Limit(50)
    messages := make([] Message, 0, 50)

    if _, err := q.GetAll(context, &messages); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, json.Marshal(messages))
    return
  }
 
  if r.Method == "POST" {

  }
}

