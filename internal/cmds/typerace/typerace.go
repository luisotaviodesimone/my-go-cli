package typerace

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/luisotaviodesimone/my-go-cli/internal/colors"
)

func StartTyperace(objective string) {
	fd := int(os.Stdin.Fd())

	// Enable raw mode
	oldState, err := setRawMode(fd)
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}
	defer restoreMode(fd, oldState) // Restore terminal mode on exit

	i := 0.0

	inputChan := make(chan string)
	timeChan := make(chan string)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("Type the following text: ")
	fmt.Println(colors.Green + "\"" + objective + "\"" + colors.Reset)
	fmt.Println("Press Enter to start")

	bufio.NewReader(os.Stdin).ReadBytes('\n')

	go promptUser(inputChan, timeChan)
	go getSignalInterruption(signalChan, timeChan)
	go startTimer(timeChan, i)

	buf := make([]byte, 1)
	whole := make([]byte, 0)

	for {
		os.Stdin.Read(buf)

		// If it is Enter keystroke, ignore iteration
		if buf[0] == 0x0A {
			continue
		}

		whole = append(whole, buf...)

		// 0x7F is the backspace character
		if buf[0] == 0x7F {

			/* TODO: Cover the case where punctuation is present, since it has two bytes */

			// Only remove the last character if the sentence is not empty
			if len(whole) > 1 {
				fmt.Print(strings.Repeat("\b", +1), " ", strings.Repeat("\b", 1)) // Remove the last character, add a space in place and then move the cursor back
				whole = whole[:len(whole)-2]
			}
			// 0x08 is the ctrl+backspace character
		} else if buf[0] == 0x08 {
			for i := len(whole) - 1; i >= 0; i-- {
				// If the current character is a space and it is the last character, ignore it and remove it in the next iteration
				if whole[i] == 0x20 && i == len(whole)-2 {
					continue
				}

				// If the for loop finds a space, remove everything after it
				if whole[i] == 0x20 {
					fmt.Print(strings.Repeat("\b", len(whole)-i), strings.Repeat(" ", len(whole)-i), strings.Repeat("\b", len(whole)-i))
					whole = whole[:i+1]
					break
				}

				// If the sentence is empty, remove everything
				if i == 0 {
					fmt.Print(strings.Repeat("\b", len(whole)), strings.Repeat(" ", len(whole)), strings.Repeat("\b", len(whole)))
					whole = whole[:0]
				}
			}
		}

		userInput := fmt.Sprint(string(whole))
		inputChan <- userInput

		if string(whole) == objective {
			fmt.Println("\nNice!")
			fmt.Println("Time spent:", <-timeChan, "\bsecs")
			os.Exit(0)
		}
	}
}
