package scanner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type sniSummary struct {
	sniProxy []string
	sync.RWMutex
}

type counterSummary struct {
	lastTime int64
	count    int64
}

var SNISummary = sniSummary{}
var CounterSummary = counterSummary{}

const counterIntervalSec = 5

func (s *sniSummary) Add(sniProxy string) {
	s.Lock()
	s.sniProxy = append(s.sniProxy, sniProxy)
	s.Unlock()
}

func (s *sniSummary) Len() int {
	s.Lock()
	n := len(s.sniProxy)
	s.Unlock()
	return n
}

func (s *sniSummary) Output(output string) {
	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("failed writing to file: %s\n", err)
	}

	wr := bufio.NewWriter(file)
	for _, data := range s.sniProxy {
		_, _ = wr.WriteString(data + "\n")
	}

	_ = wr.Flush()
	_ = file.Close()
}

func (s *counterSummary) Add(sniProxy string) {
	t := time.Now().Unix()
	s.count += 1
	if s.lastTime == 0 {
		s.lastTime = t
		return
	}
	dt := t - s.lastTime
	if dt > counterIntervalSec {
		fmt.Printf("scanner summary @@ scan speed: %d/s find: %d current:%s\n",
			s.count/dt, SNISummary.Len(), sniProxy)
		s.lastTime = t
		s.count = 0
	}
}
