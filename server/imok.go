package local_server

import (
	"fmt"
	"os"
	"time"
)

// Watcher contains data structure for watching
// whether the frontend is in normal state.
type Watcher struct {
	// used records how many times it
	// did not hear from the frontend.
	used int

	// tolerance is the max times of
	// acceptable error (no message from frontend).
	tolerance int

	// period is the time duration
	// between two times of checking.
	period time.Duration
}

// ImokWatch a Watcher object with necessary parameters.
func ImokWatch(stepLen time.Duration, tolerance int) Watcher {
	return Watcher{
		used:      0,
		tolerance: tolerance,
		period:    stepLen,
	}
}

// listen to the frontend site
func (w *Watcher) listen(ch chan bool) {
	for data := range ch {
		if data {
			w.used = 0
		}
	}
}

// check whether it is in normal state.
// If it lose connection with the frontend and out of tolerance,
// then end the waiting group to exit the program,
// in order to prevent if from computational cost.
func (w *Watcher) check() {
	for {
		time.Sleep(w.period)
		w.used++
		if w.used == w.tolerance {
			loseTolerance()
		}

		// delete me after all the test finished.
		if w.used != 0 {
			fmt.Printf("now used: %v\n", w.used)
		}
	}
}

func loseTolerance() {
	fmt.Println("Lose connection with frontend and exit.")
	os.Exit(300)
}

// Watch the chan which may send checking message from the frontend.
// Also count the time each period, as a timer.
// Once out of tolerance, this function will end.
// ---
// Notice: This loop won't block anything!
func (w *Watcher) Watch(ch chan bool) {
	go w.listen(ch)
	go w.check()
}
