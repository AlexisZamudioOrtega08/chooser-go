package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var opts = []string{"yes", "no", "idk"}

func getRandomOption() string {
	rand.Seed(time.Now().UnixNano())
	return opts[rand.Intn(len(opts))]
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["action"] = getRandomOption()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error while encoding json. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	})
	http.ListenAndServe(":3000", nil)
}
