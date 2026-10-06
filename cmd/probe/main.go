package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

type target struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func probe(tgt target) {
	start := time.Now()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(tgt.Url)
	duration := time.Since(start)
	if err != nil {
		if errors.Is(err, syscall.ECONNREFUSED) {
			fmt.Printf("%s  REFUSED \n", tgt.Name)
			return
		}
		var ne net.Error
		if errors.As(err, &ne) && ne.Timeout() {
			fmt.Printf("%s  TIMEOUT \n", tgt.Name)
			return
		}

		fmt.Printf("%s other Error : %s", tgt.Name, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("Name :%s , Response Code : %d , Duration %d UNHEALTHY \n", tgt.Name, resp.StatusCode, duration.Milliseconds())
		return
	}
	fmt.Printf("Name :%s , Response Code : %d , Duration %d \n", tgt.Name, resp.StatusCode, duration.Milliseconds())

}

func main() {

	configPath := flag.String("config", "target.json", "enter the path for the config file")
	flag.Parse()

	//fmt.Println(*configPath)

	data, err := os.ReadFile(*configPath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Config file not Found \n")
		os.Exit(1)
	}

	var targets []target
	err = json.Unmarshal(data, &targets)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in the input json file %s\n", err)
		os.Exit(1)
	}

	for _, t := range targets {
		//fmt.Printf("name: %s url:%s \n", t.Name, t.Url)
		probe(t)
	}

}
