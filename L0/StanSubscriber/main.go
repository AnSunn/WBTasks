package main

import (
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/cache"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/stan"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/storage"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
	"time"
)

var wg, wg1 sync.WaitGroup

func main() {
	//Take credentials from the file
	err := godotenv.Load("configs/api.env")
	if err != nil {
		log.Println("Cannot find config file.:", err)
	}

	// Connection to DB. If DB is unavailable, then there are 'maxRetries' attempts within 'IntervalRetry' seconds
	maxRetries := 2
	IntervalRetry := 2 * time.Second
	storage.ConnectToDB(os.Getenv("database_uri"), maxRetries, IntervalRetry)

	//Close connection with DB
	defer storage.Db.Close()

	//Initialize cache
	cache.InitCache()

	//Load data from DB to cache
	cache.LoadCacheFromDB()

	//Launch Stan
	wg.Add(1)
	go stan.NATSStreaming(&wg)

	//Launch web-server
	wg1.Add(1)
	go service.StartWebServer(&wg1)
	wg.Wait()
	wg1.Wait()
}
