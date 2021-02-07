package main

import (
  "net/http"
  "fmt"
  "log"
  "context"

  "go.temporal.io/sdk/client"
  "lieroarbetsprov/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main(){
  r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
    log.Println("Request received")
		result := startActivity()
    w.Write([]byte(fmt.Sprintf("<h1>%s</h1>", result)))
	})
  port := ":3333"
  log.Println("Serving Http on port", port)
	err := http.ListenAndServe(port, r)
  if err != nil {
    log.Fatalln("Could not Initialize HTTP Server: ", err)
  }
}

func startActivity() string {
  c, err := client.NewClient(client.Options{})
  if err != nil {
    log.Fatalln("Kunde int skapa klient")
  }
  defer c.Close()
  options := client.StartWorkflowOptions{
    ID: "increment-workflow",
    TaskQueue: app.IncrementTaskQueue,
  }
  we, err := c.ExecuteWorkflow(context.Background(), options, app.GetIncWorkflow, 1)
  if err != nil {
    log.Fatalln("Kunde int slutföra workflow")
  }
  var result string
  err = we.Get(context.Background(), &result)
  if err != nil {
    log.Fatalln("Kunde int få resultat från workflow")
  }
  return result
}
