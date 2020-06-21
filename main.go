package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"

	"github.com/mengskysama/sni-scanner/scanner"
)

func main() {
	taskFile := flag.String("f", "ip.txt", "ip list")
	target := flag.String("s", "https://www.baidu.com", "target site for test")
	threadStr := flag.String("t", "100", "thread for test")
	sniPort := flag.String("p", "443", "SNI port")
	output := flag.String("o", "result.txt", "output result result.txt")
	flag.Parse()

	thread, err := strconv.Atoi(*threadStr)
	if err != nil {
		fmt.Println("invalid thread")
		return
	}

	tasks, n := scanner.LoadTask(*taskFile)
	if n == 0 {
		fmt.Println("no task found")
		return
	}

	wg := sync.WaitGroup{}
	taskCh := make(chan string)
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go scanner.Worker(*target, taskCh, &wg)
	}

	scanner.Dispatcher(tasks, *sniPort, taskCh)
	wg.Wait()

	fmt.Println("all of worker quited")
	scanner.SNISummary.Output(*output)
}
