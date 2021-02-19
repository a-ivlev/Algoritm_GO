package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Job struct {
	req string
	payload []byte
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
}

func main() {
	var url, reqMetod string
	var numThreads, numRequest, count, countBad int

	flag.StringVar(&reqMetod, "m", "GET", "http-method")
	flag.StringVar(&url, "u", "http://localhost:8081/items", "http address")
	flag.IntVar(&numThreads, "t", 2, "Number of threads to send a request, default 10")
	flag.IntVar(&numRequest, "r", 100, "Number of requests to the server, default 100")

	flag.Parse()

	switch reqMetod {
	case "GET":
	case "POST":
	case "PUT":
	case "DELETE":
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

	var totalDuration time.Duration

	if numThreads > numRequest {
		numThreads = numRequest
	}
	for i := 0; i < numThreads; i++ {
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.Handle()
	}

	for j := 0; j < numRequest; j++{
		startReq := time.Now()
		//resp, err := http.Get("http://localhost:8081/items")
		req, _ := http.NewRequest(reqMetod, "http://localhost:8081/items", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		elapsed := time.Since(startReq)*time.Nanosecond
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != 200 {
			countBad++
		}
		totalDuration += elapsed
		count++
		text := fmt.Sprintf("\nУспешных запросов %d, время отклика на запрос %v, Статус код = %s\n", count-countBad, elapsed*time.Nanosecond, resp.Status)
		jobChan <- &Job{
			req: "resp.Body",
			payload: []byte(text),
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
		log.Printf("worker %d processing job with payload %s", w.num, string(job.payload))
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	fmt.Printf("NewWorker num %d \n", num)
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}
