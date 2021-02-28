package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Job struct {
	req http.Request
	mu *sync.RWMutex
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
}

var totalDuration time.Duration
var count, countBad int
var mutex sync.RWMutex

func main() {
	var url, reqMetod string
	var numThreads, numRequest  int

	flag.StringVar(&reqMetod, "m", "GET", "http-method")
	flag.StringVar(&url, "u", "http://localhost:8081/items", "http address")
	flag.IntVar(&numThreads, "t", 2, "Number of threads to send a request, default 10")
	flag.IntVar(&numRequest, "r", 100, "Number of requests to the server, default 100")

	flag.Parse()

	switch reqMetod {
	case "GET", "POST", "PUT", "DELETE":
	default:
		fmt.Println("Введён не корректный http-method")
		os.Exit(1)
	}

	if numThreads < 1 || numRequest < 1 {
		fmt.Println("Количество запросов и количество потоков не может быть меньше 1")
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)

	if numThreads > numRequest {
		numThreads = numRequest
	}
	for i := 0; i < numThreads; i++ {
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.Handle()
	}

	for j := 0; j < numRequest; j++{
		req, _ := http.NewRequest(reqMetod, "http://localhost:8081/items", nil)
		jobChan <- &Job{
			req: *req,
		}
	}
	close(jobChan)
	wg.Wait()
	fmt.Printf("Общее количество запросов %d Успешных запросов %d Запросы завершённые с ошибкой %d\n", count, count-countBad, countBad)
	fmt.Println("RPS = ", time.Duration(totalDuration.Nanoseconds()/int64(count))*time.Nanosecond)
}

func (w *Worker) Handle() {
	defer w.wg.Done()
	for job := range w.jobChan {
		startReq := time.Now()
		resp, _ := http.DefaultClient.Do(&job.req)
		elapsed := time.Since(startReq)*time.Nanosecond
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != 200 {
			mutex.Lock()
			countBad++
			mutex.Unlock()
		}
		mutex.Lock()
		totalDuration += elapsed
		count++
		mutex.Unlock()
		fmt.Printf("\nВремя отклика на запрос %v, Статус код = %s\n", elapsed*time.Nanosecond, resp.Status)
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}
