package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
)

type Job struct {
	req http.Request
	payload []byte
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
}

func main() {
	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)
	for i := 0; i < 2; i++ {
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.Handle()
	}

	for j := 0; j < 100; j++{

		req, err := http.NewRequest("GET", "localhost:8081/items", nil )
		if err != nil {
			log.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := &Handler{}
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			log.Fatal("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		text := fmt.Sprintf("Some message %d status %d", j+1, rr.Code)
		jobChan <- &Job{
			req: *req,
			payload: []byte(text),
		}
	}

	//jobChan <- &Job{
	//	payload: []byte("Some message 1"),
	//}
	//jobChan <- &Job{
	//	payload: []byte("Some message 2"),
	//}
	//jobChan <- &Job{
	//	payload: []byte("Some message 3"),
	//}
	//jobChan <- &Job{
	//	payload: []byte("Some message 4"),
	//}
	//jobChan <- &Job{
	//	payload: []byte("Some message 5"),
	//}
	//jobChan <- &Job{
	//	payload: []byte("Some message 6"),
	//}
	close(jobChan)
	wg.Wait()

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
