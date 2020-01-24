// Sample program to demonstrate usage of an unbuffered channel
// to simulate a relay race between four goroutines

package main

import(
    "fmt"
    "sync"
    "time"
)


var wg sync.WaitGroup

func main() {
    // Create an unbuffered channel
    track := make(chan int)

    // Add a count of one for the last runner.
    wg.Add(1)

    // First Runner to his mark.
    go Runner(track)
    // Start the race. 
    track <- 1

    // Wait for the race to finish 
    wg.Wait()
}

// Runner simulates a person running in the relay race
func Runner(track chan int) {
    // Wait to receive the baton 
    baton := <-track
    const maxExchanges = 4
    var exchange int
    fmt.Printf("Runner %d Running with Baton\n", baton)

    // New runner to the line.
    if baton < maxExchanges {
        exchange = baton + 1
        fmt.Printf(" Runner %d To The Line\n", exchange)
        go Runner(track)
    }

    // Running around the track
    time.Sleep(100 * time.Millisecond)
    // Is the race over 
    if baton == maxExchanges {
        fmt.Printf("Runner %d Finished, Race Over\n", baton)
        wg.Done()
        return
    }

    // Exchange the baton for the next Runner.
    fmt.Printf("Runner %d Exchane with Runner %d\n",baton, exchange)
    track <- exchange 
}
