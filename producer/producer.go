package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	//  connecting to  Redis
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close() //بستن redis

	//   payload explain 
	payload := map[string]interface{}{
		"user_id":   42,
		"task_name": "send_email",
	}

	//  setting payload in json array
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("invalid to send paylaoad into json%v ", err)
	}

	//  make a job and sending in database
	task := asynq.NewTask("email:send", payloadBytes)

	//  seding job oin queue
	info, err := client.Enqueue(task, asynq.Queue("default"), asynq.MaxRetry(5), asynq.Timeout(30*time.Second))
	if err != nil {
		log.Fatalf("invalid to send job %v", err)
	}

	//  show jobs informations
	log.Printf("job sent : ID=%s", info.ID)
}
