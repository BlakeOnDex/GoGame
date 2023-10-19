package main

import (
	"fmt"
	"os"
)

func main() {
	
	Direction := ""
	Answer := 0
	Riddle := ""
	fmt.Println("Welcome to my game!")
	fmt.Println(" Choose:left or right")
	fmt.Scanln(&Direction)
	switch Direction {
	case "left":
		fmt.Println("GameOver")
		fmt.Scanln()
		os.Exit(0)
	case "right":
		fmt.Println("Solve this equation")
		fmt.Println("120 + 678")
		fmt.Scanln(&Answer)
		switch Answer {
		case 798:
			fmt.Println("move on to the next level")
			fmt.Println("left or right")
			fmt.Scanln(&Direction)
			switch Direction {
			case "left":
				fmt.Println("Solve this riddle")
				fmt.Println("I’m often created weak so I’m not forgotten. Sometimes I’m reset, if my owner forgets me. What am I?")
				fmt.Scanln(&Riddle)
				switch Riddle {
				case "password":
					
				    fmt.Println("Exiting program...")
					fmt.Scanln()
					os.Exit(0)
					
				}

			case "right":
				fmt.Println("GameOver")
				fmt.Scanln()
				os.Exit(0)
			}
		default:
			fmt.Println("GameOver")
			fmt.Scanln()
			os.Exit(0)
		}

	}

}
