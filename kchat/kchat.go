package kchat
 
import (
  "appengine"
  "appengine/channel"

  "html/template"
  // "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/", root)
}

var mainTemplate = template.Must(template.ParseFile("main.html"))

func root(w http.ResponseWriter, r *http.Request) {
  context := appengine.NewContext(r)
  token, err := channel.Create(context, "kchat")
  if err != nil {
    http.Error(w, "Could not create channel", http.StatusInternalServerError)
    context.Errorf("channel.Create: %v", err)
    return
  }

  err = mainTemplate.Execute(w, map[string] string { "token": token })
  if err != nil {
    context.Errorf("mainTemplate: %v", err)
  }
  
}

