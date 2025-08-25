package main

import (
	"fmt"
)

type Job struct {
	State string
	done  chan struct{}
}

func (j *Job) Wait() {
	<-j.done
}

func (j *Job) Done() {
	j.State = "done"
	close(j.done)
}

func main() {
	ch := make(chan *Job)
	go func() {
		j := <-ch
		j.Done()
	}()

	job := Job{"ready", make(chan struct{})}
	ch <- &job
	job.Wait()
	fmt.Println(job.State)
}

// The output will be: done
// The job's state is updated to "done" after the goroutine signals completion.
// The main function waits for the job to be done before printing its state.
// This demonstrates how to use channels to synchronize the state of a job in a concurrent environment.
