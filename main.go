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
	HealthyStatus  string
	CrashOnRequest bool
	CrashTimerMin  int
	CrashTimerMax  int
}

var settings Settings

func loadSettings() {
	settings = Settings{
		HealthyStatus:  "Healthy",
		CrashOnRequest: false,
		CrashTimerMin:  -1,
		CrashTimerMax:  60,
	}
	if len(os.Getenv("PETSTORE_FAIL")) != 0 {
		settings.HealthyStatus = "Unhealthy"
	}
	if len(os.Getenv("PETSTORE_CRASH")) != 0 {
		settings.CrashOnRequest = true
	}
	if len(os.Getenv("PETSTORE_CRASHTIMER_MIN")) != 0 {
		i, err := strconv.Atoi(os.Getenv("PETSTORE_CRASHTIMER_MIN"))
		if err == nil {
			settings.CrashTimerMin = i
		}
	}
	if len(os.Getenv("PETSTORE_CRASHTIMER_MAX")) != 0 {
		i, err := strconv.Atoi(os.Getenv("PETSTORE_CRASHTIMER_MAX"))
		if err == nil {
			settings.CrashTimerMax = i
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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	loadSettings()

	if settings.CrashTimerMin > -1 {
		randomSeconds := rand.Intn(settings.CrashTimerMax - settings.CrashTimerMin)
		fmt.Printf("Crashing randomly in %d second(s)...\n", settings.CrashTimerMin+randomSeconds)
		timer1 := time.NewTimer(time.Duration(settings.CrashTimerMin+randomSeconds) * time.Second)

		go func() {
			<-timer1.C
			fmt.Println("\nCrash timer expired")
			os.Exit(1)
		}()
	}
}

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "port number")
	flag.Parse()

	fmt.Printf("Starting server on port %d...", port)
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthcheck", healthcheck)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		panic(err)
	}
}
