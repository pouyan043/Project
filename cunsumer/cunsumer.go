package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	// اconnecting to Redis
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10, // jobs that can proccesing at the same time 
		},
	)

	//  تعریف handler
	mux := asynq.NewServeMux()
	mux.HandleFunc("email:send", handleEmailTask) // explain jobs for handler email:send

	// starting server
	if err := server.Run(mux); err != nil {
		log.Fatalf("invalid to start server: %v", err)
	}
}

// func for job handler
func handleEmailTask(ctx context.Context, t *asynq.Task) error {
	// return informations
	var payload map[string]interface{}
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err //
	}

	//return value from payload 
	userID, ok := payload["user_id"].(float64) // convert number from float 64 to int from json unmarshal
	if !ok {
		return &json.UnmarshalTypeError{} // return eror user id
	}

	taskName, ok := payload["task_name"].(string)
	if !ok {
		return &json.UnmarshalTypeError{} // return eror task name
	}

	// proccesing job
	log.Printf("proccesing job : task_name=%s, user_id=%d", taskName, int(userID))
	log.Println("the job succes")
	return nil
}
