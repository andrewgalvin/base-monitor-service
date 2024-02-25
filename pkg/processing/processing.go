package processing

import "net/http"

// Replace with what each "worker" will
// do on their own thread

// Worker is a function that processes a job
// It is important to note that the function
// should be what "one" worker does on their
// own thread, and not the entire worker pool

func ProcessTask(client *http.Client) error {
	// Do some processing here
	return nil
}
