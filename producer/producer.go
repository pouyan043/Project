package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	//  اتصال به Redis
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close() //بستن redis

	//  تعریف payload
	payload := map[string]interface{}{
		"user_id":   42,
		"task_name": "send_email",
	}

	//  قرار دادن payload تو ارایه json
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("invalid to send paylaoad into json%v ", err)
	}

	//  ایجاد شغل و فرساتادن دیتا
	task := asynq.NewTask("email:send", payloadBytes)

	//  فرستادن شغل به صف
	info, err := client.Enqueue(task, asynq.Queue("default"), asynq.MaxRetry(5), asynq.Timeout(30*time.Second))
	if err != nil {
		log.Fatalf("invalid to send job %v", err)
	}

	//  نشون دادن اطلاعات شغل
	log.Printf("job sent : ID=%s", info.ID)
}
