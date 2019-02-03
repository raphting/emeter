package main

import (
	"fmt"
	"github.com/tarm/serial"
	"regexp"
	"strings"
)

func main() {
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
			fmt.Println(res[1])
		}
	}

}
