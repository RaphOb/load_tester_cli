package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aquasecurity/table"
)

type Rapport struct {
	mu      sync.Mutex
	mapping map[string]map[string][]float32
}

func (r *Rapport) display(nbRequest int, d time.Duration) {
	t := table.New(os.Stdout)
	t.SetHeaderStyle(table.StyleBold)
	t.SetDividers(table.UnicodeRoundedDividers)
	t.AddHeaders("Request", "Method", "N° Req", "Success", "Failure", "Avg", "Max", "Min", "Med")
	t.AddFooters("Summary", fmt.Sprintf("Total elapsed time: %s", d))
	for k, v := range r.mapping {
		parts := strings.Split(k, "__")
		url, method := parts[0], parts[1]
		times := v["times"]
		status := v["status"]
		min := slices.Min(times)
		max := slices.Max(times)
		sum := float32(0.0)
		for i := 0; i < len(times); i++ {
			sum += times[i]
		}
		avg := sum / float32(len(times))
		slices.Sort(times)
		mid := len(times) / 2
		var med float32
		if len(times)%2 == 0 {
			med = (times[mid-1] + times[mid]) / 2
		} else {
			med = times[mid]
		}

		success, failure := countStatus(status)

		t.AddRow(url, method, strconv.Itoa(nbRequest), strconv.Itoa(success), strconv.Itoa(failure), fmt.Sprintf("%f", avg), fmt.Sprintf("%f", max), fmt.Sprintf("%f", min), fmt.Sprintf("%f", med))
	}
	t.Render()
}
