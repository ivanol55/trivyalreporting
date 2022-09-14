package requiredArgCount

import (
	"fmt";
	"os";
)

func CheckForArgs(args []string, requiredCount int) {
	var argumentCount int = len(args)
    if argumentCount < requiredCount {
		fmt.Println("")
		fmt.Println("Seems like no arguments were provided for the script! Please check for the proper 'help' usage page to learn more.")
		os.Exit(0)
    }
}