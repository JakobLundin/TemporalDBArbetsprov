package main

import (
    "log"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"

    "lieroarbetsprov/app"
)

func main() {
    c, err := client.NewClient(client.Options{})
    if err != nil {
        log.Fatalln("unable to create Temporal client", err)
    }
    defer c.Close()
    w := worker.New(c, app.IncrementTaskQueue, worker.Options{})
    w.RegisterWorkflow(app.GetIncWorkflow)
    w.RegisterActivity(app.GetAndIncrement)

    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("unable to start Worker", err)
    }
}
