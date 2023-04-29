package main

import (
	"bufio"
	"flag"
	"net/url"
	"strings"
	"fmt"
	"os"
	"sync"
)

func main() {

	concurrency := 20
	var domain string
	flag.IntVar(&concurrency, "c", 20, "Set the concurrency level")
	flag.StringVar(&domain,"d","","Enter domain for which you wan to filter the URLS")
	flag.Parse()

	jobs := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)

		go func() {
			for iurl := range jobs {
				url, err := url.Parse(iurl)
				if err != nil {
					continue
				}
				if strings.HasSuffix(url.Hostname(),"."+domain) == true || url.Hostname() == domain {
					fmt.Println(iurl)
				}
			}

			wg.Done()
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)
	for sc.Scan() {
		jobs <- sc.Text()
	}

	close(jobs)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	wg.Wait()

}
