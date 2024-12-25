package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	// اتصال به Redis
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10, // تعداد شغلی ک همزمان میتونن پردازش بشن
		},
	)

	//  تعریف handler
	mux := asynq.NewServeMux()
	mux.HandleFunc("email:send", handleEmailTask) // تعریف handler برای jobهای email:send

	//  راه‌اندازی سرور
	if err := server.Run(mux); err != nil {
		log.Fatalf("خطا در اجرای سرور: %v", err)
	}
}

// funcبرای job handler
func handleEmailTask(ctx context.Context, t *asynq.Task) error {
	// برگردوندن اطلاعات
	var payload map[string]interface{}
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err //
	}

	//برگردوندن مقدار از payload
	userID, ok := payload["user_id"].(float64) // تبدیل اعداد از float 64 be int توسط json unmarshal
	if !ok {
		return &json.UnmarshalTypeError{} // خطای برگردوندن user id
	}

	taskName, ok := payload["task_name"].(string)
	if !ok {
		return &json.UnmarshalTypeError{} // خطای برگردوندن task name
	}

	// پردازش job
	log.Printf("proccesing job : task_name=%s, user_id=%d", taskName, int(userID))
	log.Println("the job succes")
	return nil
}
