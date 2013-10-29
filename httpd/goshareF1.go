package abkhttpd

import (
  "net/http"
  "html/template"
)

func F1(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  t, _ := template.ParseFiles("httpd/public/help.html")
  t.Execute(w, nil)
}

func HelpHTTP(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  t, _ := template.ParseFiles("public/index.html")
  t.Execute(w, nil)
}

func HelpZMQ(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  t, _ := template.ParseFiles("public/index.html")
  t.Execute(w, nil)
}

func Status(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  t, _ := template.ParseFiles("public/index.html")
  t.Execute(w, nil)
}
