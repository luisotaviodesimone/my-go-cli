package typerace

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/luisotaviodesimone/my-go-cli/internal/colors"
)

func timer(timePrint float64, interval time.Duration) string {
	timeStr := strconv.FormatFloat(timePrint, 'f', 2, 64)
	time.Sleep(interval)
	return timeStr
}

func getSignalInterruption(signalChan chan os.Signal, timeChan chan string) {

	for sig := range signalChan {
		fmt.Println("\nGot signal:", sig)
		fmt.Printf("Time spent: %vsecs\n", <-timeChan)
		os.Exit(0)
	}
}

func startTimer(timeChan chan string, timerTime float64) {
	for {
		time := timer(timerTime, 10*time.Millisecond)
		timeChan <- time
		timerTime += 0.01
	}
}

func promptUser(inputChan, timeChan chan string) {
	input := ""
	timer := "0.000"
	fmt.Print("\r", colors.Blue+"Start typing: "+colors.Reset, input, "█ ", colors.Red, timer, "secs", colors.Reset)
	for {
		// TODO: Find a solution to the problem where the cursor not being at the end of the typed text
		select {
		case input = <-inputChan:
			fmt.Print("\r", colors.Blue+"Start typing: "+colors.Reset, input, "█ ", colors.Red, timer, "secs", colors.Reset)
		case timer = <-timeChan:
			fmt.Print("\r", colors.Blue+"Start typing: "+colors.Reset, input, "█ ", colors.Red, timer, "secs", colors.Reset)
		default:
		}
	}
}
