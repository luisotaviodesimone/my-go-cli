package typerace

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
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

	fmt.Println(colors.Cyan, "\rStart typing: ", colors.Reset)
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

		var userInput string
		if string(buf[0]) != objective[len(whole)-1:len(whole)] {
			userInput = fmt.Sprint(colors.Red, string(whole), colors.Reset)
		} else {
			userInput = fmt.Sprint(colors.Green, string(whole), colors.Reset)
		}

		inputChan <- userInput

		if string(whole) == objective {
			fmt.Println("Time spent:", <-timeChan, "\bsecs")
			timeFloat, _ := strconv.ParseFloat(<-timeChan, 64)
			wordsCount := float64(len(strings.Fields(string(whole))))
			fmt.Println("WPM:", strconv.FormatFloat(wordsCount/(timeFloat/60), 'f', 2, 64), "\bwpm")
			charsCount := float64(len(string(whole)))
			fmt.Println("CPS:", strconv.FormatFloat(charsCount/(timeFloat/60), 'f', 2, 64), "\bcpm")
			os.Exit(0)
		}
	}
}
