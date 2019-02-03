package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var globalPage string

func main() {
	go readData()

	// Handle and serve front page
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, globalPage)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
	log.Fatal(http.ListenAndServe(":9688", nil))
}

func readData() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 128)
	r := regexp.MustCompile("1-0:1\\.7\\.0\\((\\d{2}\\.\\d{3})\\*kW\\)")

	for {
		fulltext := ""
		for {
			n, err := s.Read(buf)
			if err != nil {
				fmt.Println(err)
			}

			reception := string(buf[:n])
			fulltext += reception
			if strings.Contains(reception, "!") {
				break
			}
		}

		if res := r.FindStringSubmatch(fulltext); res != nil {
			globalPage = fmt.Sprintf(
				"# HELP emeter_pwr_delivered Actual electricity power delivered.\n# TYPE emeter_pwr_delivered gauge\nemeter_pwr_delivered %s", res[1])
		}
	}
}
