package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/AndrewDonelson/GoJobQue/models"
)

// MaxWorker allow setting the maximum number of workers
var MaxWorker int

// MaxQueue allows setting the maxmimum que size
//var MaxQueue int

var QueMgr *models.QueManager

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	MaxWorker, err := strconv.Atoi(os.Getenv("MAX_WORKERS"))
	if err != nil {
		MaxWorker = 16
	}

	// MaxQueue, err := strconv.Atoi(os.Getenv("MAX_QUEUE"))
	// if err != nil {
	// 	MaxQueue = 16
	// }
	QueMgr, err = models.NewQueManager(MaxWorker)
	http.HandleFunc("/", HelloServer)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
