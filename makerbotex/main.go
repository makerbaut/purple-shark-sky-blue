package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

const (
	redisURL      = "redis:6379"
	answerBaseURL = "http://answer:3000"
)

// This function will iterate over all the keys from the redis server. It will only
// decipher "list" and "set" types due to the requirements of this exercise.
// If an error occurs, a panic() is thrown to abort execution. Upon getting a successful
// response from the server, the program will exit gracefully with exit code 0.
func main() {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	var checksum int
	for _, k := range client.Keys("*").Val() {
		var rows []string
		switch client.Type(k).Val() {
		case "set":
			rows = client.SMembers(k).Val()
		case "list":
			rows = client.LRange(k, 0, -1).Val()
		default:
			log.Println("Ignoring unknown type")
		}

		if hasAnagram(rows) || hasTwoNumbersDividesTo177(rows) {
			continue
		}

		checksum += getMaxMinDifference(rows)
	}

	url := fmt.Sprintf("%s/%d", answerBaseURL, checksum)
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("Received OK from server: %s", str)
		os.Exit(0)
	} else {
		log.Printf("Did not receive successful response from server: %s", str)
		// NOTE: could also throw a panic() here to be consistent
		os.Exit(-1)
	}
}
