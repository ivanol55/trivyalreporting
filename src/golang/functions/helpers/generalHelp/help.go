// Sets the package name to import from the helper runner
package generalHelp

// Imports necessary packages for the function to print text into the terminal
import (
	"fmt"
)

// Declares a function that prints the general system help
func ShowHelp() {
	fmt.Println("This tool offers several scanning options depending on your needs. Infrastructure, repositories or docker images, we've got you covered. Here's what's available as first-level options (more features in development!):")
	fmt.Println("    - 'infra' will let you scan and generate reports on your currently running infrastructure. run this script with 'infra' or 'infra help' to get more information")
	fmt.Println("")
}
