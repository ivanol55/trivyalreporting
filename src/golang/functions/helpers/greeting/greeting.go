// Sets the package name to import from the helper runner
package greeting

// Imports necessary packages for the function to print text into the terminal
import (
	"fmt"
)

// Declares a function that prints a stylized application name and a welcome message
func ShowGreeting() {
	fmt.Println("")
	fmt.Println(" ############################################################################")
	fmt.Println(" ############################################################################")
	fmt.Println("  _        _                   _                           _   _             ")
	fmt.Println(" | |      (_)                 | |                         | | (_)            ")
	fmt.Println(" | |_ _ __ ___   ___   _  __ _| |_ __ ___ _ __   ___  _ __| |_ _ _ __   __ _ ")
	fmt.Println(" | __| '__| \\ \\ / / | | |/ _` | | '__/ _ \\ '_ \\ / _ \\| '__| __| | '_ \\ / _` |")
	fmt.Println(" | |_| |  | |\\ V /| |_| | (_| | | | |  __/ |_) | (_) | |  | |_| | | | | (_| |")
	fmt.Println(" \\__|_|  |_| \\_/  \\__, |\\__,__|_|_|  \\___| .__/ \\___/|_|   \\__|_|_| |_|\\__, | ")
	fmt.Println(" 		   __/ |                | |                            __/  | ")
	fmt.Println(" 		  |___/                 |_|                           |____/  ")
	fmt.Println(" ############################################################################")
	fmt.Println(" ############################################################################")
	fmt.Println("")
	fmt.Println("Welcome to TrivyalReporting! this is a helper tool used to generate on-demand reports for infrastructure, docker images and repository configuration. Let's avoid incidents the fancy way!")
}
