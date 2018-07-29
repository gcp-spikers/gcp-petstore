package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Settings struct {
	HealthyStatus          string
	CrashOnRequest         bool
	CrashAfterTime         int
	CrashTimeRandomSeconds int
}

var settings Settings

func loadSettings() {
	settings = Settings{
		HealthyStatus:          "Healthy",
		CrashOnRequest:         false,
		CrashAfterTime:         -1,
		CrashTimeRandomSeconds: 60,
	}
	if len(os.Getenv("PETSTORE_FAIL")) != 0 {
		settings.HealthyStatus = "Unhealthy"
	}
	if len(os.Getenv("PETSTORE_CRASH")) != 0 {
		settings.CrashOnRequest = true
	}
	if len(os.Getenv("PETSTORE_CRASH_AFTER_TIME")) != 0 {
		i, err := strconv.Atoi(os.Getenv("PETSTORE_CRASH_AFTER_TIME"))
		if err == nil {
			settings.CrashAfterTime = i
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	type Person struct {
		Health      string   `json:"health"`
		Environment []string `json:"environment"`
	}

	data := Person{
		Health:      settings.HealthyStatus,
		Environment: os.Environ(),
	}
	json, err := json.MarshalIndent(data, "", "\t")
	if err != nil || settings.CrashOnRequest {
		log.Fatal(err)
	}
	io.WriteString(w, fmt.Sprintf(string(json[:])))
}

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "port number")
	flag.Parse()

	loadSettings()

	if settings.CrashAfterTime > -1 {
		timer1 := time.NewTimer(settings.CrashAfterTime * time.Second)
		go func() {
			<-timer1.C
			randomSeconds := rand.Intn(settings.CrashTimeRandomSeconds)
			fmt.Printf("\nCrash timer expired, starting random timer of %d seconds...", randomSeconds)
			time.Sleep(time.Duration(settings.CrashAfterTime+randomSeconds) * time.Second)
			fmt.Println("\nTimer expired, crashing...")
			os.Exit(1)
		}()
	}

	fmt.Printf("Starting server on port %d...", port)
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthcheck", healthcheck)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		panic(err)
	}
}
