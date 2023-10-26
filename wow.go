package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var minutes int
	// perl -ne 'if (/td..strong.*style=3D...>(\d+)</)
	// r := regexp.MustCompile(`/td..strong.*style=3D...>(\d+)</`)
	r := regexp.MustCompile(`.*td.*strong.*style=3D.*>(\d+)<.*`)

	for s.Scan() {
		res := r.FindAllStringSubmatch(s.Text(), -1)
		if len(res) > 0 {
			if i, err := strconv.Atoi(res[0][1]); err == nil {
				minutes += i
			}
		}

	}
	fmt.Printf("minutes: %d  hours: %d\n", minutes, minutes/60.0)
}
