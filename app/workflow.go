package app

import (
    "time"

    "go.temporal.io/sdk/workflow"
)

func GetIncWorkflow(ctx workflow.Context, id int) (string, error) {
    options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 30,
    }
    ctx = workflow.WithActivityOptions(ctx, options)
    var result string
    err := workflow.ExecuteActivity(ctx, GetAndIncrement, id).Get(ctx, &result)
    return result, err
}
