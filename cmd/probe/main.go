package main

import ( 
         "flag"
     	 "fmt"
	 "os"
	 "encoding/json"
	 "net/http"
	 "time"
       )

type target struct {
	Name string `json:"name"`
	Url string `json:"url"`
}


func main() {

	configPath := flag.String("config" , "target.json", "enter the path for the config file")
	flag.Parse()

	fmt.Println (*configPath)

	data, err := os.ReadFile(*configPath)

	if err != nil {
		panic(err)
	}

	var targets []target
	err = json.Unmarshal(data, &targets)
	if err !=nil {
		panic(err)
	}

	client :=&http.Client{Timeout: 5 * time.Second}

	for _, t := range targets {
		fmt.Printf("name: %s url:%s \n",t.Name,t.Url)
		start := time.Now()
		resp, err  := client.Get(t.Url)
		duration := time.Since(start)
		if err != nil {
			fmt.Printf("Name : %s Errror= %s\n",t.Name, err)
			continue
		}
		defer resp.Body.Close()
		fmt.Printf("name: %s  status: %d  duration: %s\n", t.Name, resp.StatusCode, duration)
	}

}

