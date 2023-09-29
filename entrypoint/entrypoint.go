package entrypoint

import (
	"fmt"
	"net/http"
	"sync"
)

func CurrentPort() (int, error) {
	var wg sync.WaitGroup
	var dynamicPort int
	var mutex sync.Mutex

	for port := 1000; port <= 65000; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			url := fmt.Sprintf("http://10.49.122.144:%d/ping", port)
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				mutex.Lock()
				dynamicPort = port
				mutex.Unlock()
			}
		}(port)
	}

	wg.Wait()

	if dynamicPort == 0 {
		return 0, fmt.Errorf("no dynamic port found")
	}
	return dynamicPort, nil
}
