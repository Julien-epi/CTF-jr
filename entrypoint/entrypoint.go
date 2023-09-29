package entrypoint

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func CurrentPort() (int, error) {
	var dynamicPort int
	var mutex sync.Mutex
	portFound := false

	var wg sync.WaitGroup

	for port := 1000; port <= 10000; port++ {
		if portFound {
			log.Println("Port found, stopping search.")
			break
		}

		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			url := fmt.Sprintf("http://10.49.122.144:%d/ping", port)

			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Error on port %d: %v", port, err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				mutex.Lock()
				dynamicPort = port
				mutex.Unlock()
				log.Printf("Port found: %d", dynamicPort)
				portFound = true
				return
			}
		}(port)
	}

	wg.Wait()

	if dynamicPort == 0 {
		log.Println("No dynamic port found")
		return 0, fmt.Errorf("no dynamic port found")
	}
	return dynamicPort, nil
}
