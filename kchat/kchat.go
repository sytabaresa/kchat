package kchat
 
import (
  "appengine"
  "appengine/datastore"

  "encoding/json"
  "html/template"
  "net/http"
  "time"
  "fmt"
  //"io/ioutil"
)

type Message struct {
  Name  string     `json:"name"`
  Body  string     `json:"body"`
  Date  time.Time  `json:"date"`
}

func init() {
  http.HandleFunc("/", root)
  http.HandleFunc("/chat", chat)
}

func root(w http.ResponseWriter, r *http.Request) {
  var t = template.Must(template.ParseFiles("/home/dv/Dropbox/code/kchat/kchat/main.html"))
  var page struct {}
  
  if err := t.Execute(w, page); err != nil {
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
    payload, _ := json.Marshal(messages)
    fmt.Fprint(w, string(payload))
    return
  }
 
  if r.Method == "POST" {
   // payload, _ := ioutil.ReadAll(r.Body)

    msg := Message{
      Name:  "johan",
      Body:  "hello",
      Date:  time.Now(),
    }

    datastore.Put(context, datastore.NewIncompleteKey(context, "Message", nil), &msg)
    fmt.Fprintf(w, "ok")
  }
}

