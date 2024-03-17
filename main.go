package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ReqHandler(protocol string, domain string, port int, error_file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	url := fmt.Sprintf("%s://%s:%d", protocol, domain, port)
	// fmt.Printf("URL is %s\n", url)

	req, err := client.Get(url)
	if err != nil {
		var errorMsg string
		if netError, ok := err.(*net.OpError); ok && netError != nil {
			errorMsg = fmt.Sprintf("Domain %s is not reachable: %v\n", domain, netError.Err)

		} else {
			errorMsg = fmt.Sprintf("Error making HTTP request for %s: %v\n", domain, err)
		}

		if _, err := error_file.Write([]byte(errorMsg)); err != nil {
			fmt.Printf("Error writing to error file: %v\n", err)

		}

		return // Exit the function if there is an error
	}

	defer req.Body.Close()

	status := req.StatusCode

	if status == 200 {
		fmt.Printf("%s\n", url)
		// fmt.Printf("The status code for the  protcol %s for %s:%d is %d\n", protocol, domain, port, status)
	}

}

var domains_file = flag.String("f", "domains", "Enter the file name of domains")

/*
here the flag wich we will use let's say -f domains_list it will be pointer stored in domains_file , so now when i want to open a file

fileOpen := os.Open(*domains_file)
*/
var portsFlag = flag.String("p", "80,443", "Enter ports to make request with commas i.e 80,443")

func main() {

	flag.Parse()

	portStrings := strings.Split(*portsFlag, ",")
	var ports []int

	for _, z := range portStrings {
		port, err := strconv.Atoi(z)
		if err != nil {
			fmt.Printf("Invalid port number: %v\n", err)
			return // Exit if there's an invalid port
		}
		ports = append(ports, port)
	}

	fmt.Printf("Main Function Entering ..\n")

	error_file, err := os.OpenFile("error-file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening error log file: %v\n", err)
		return // Exit if the error file cannot be opened
	}

	file, err := os.Open(*domains_file)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	protocols := []string{"http", "https"}
	var wg sync.WaitGroup

	for scanner.Scan() {
		domains := scanner.Text()
		for _, p := range ports {
			for _, prot := range protocols {
				// fmt.Printf("The url on which the request is:= %s://%s:%d\n", prot, d, p)
				wg.Add(1)
				go ReqHandler(prot, domains, p, error_file, &wg)
			}

		}

	}

	wg.Wait()

}
