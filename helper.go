// Sets the package name for the main script
package main

// Imports necessary packages for the main logic loop to run the necessary helpers and tools based on script arguments
import (
	"os"
	"trivyalreporting/src/golang/functions/helpers/generalHelp"
	"trivyalreporting/src/golang/functions/helpers/greeting"
	"trivyalreporting/src/golang/functions/helpers/requiredArgCount"
	"trivyalreporting/src/golang/functions/infra/generateInfraReport"
	"trivyalreporting/src/golang/functions/infra/infraHelp"
)

// Function that runs when the program is started, executes the main application logic
func main() {
	// Show the application greeting with ascii art title
	greeting.ShowGreeting()
	// Check if enough args were provided. If not, the program shows an error and exits
	requiredArgCount.CheckForArgs(os.Args, 2)
	// Check which tool to run depending on the first script argument
	switch os.Args[1] {
	// If "help" is provided as the first script argument, show the application general help page and exit the program
	case "help":
		generalHelp.ShowHelp()
	// If "infra" is provided as the first script argument, run the infra application logic check
	case "infra":
		// Check if the required amount of arguments were sent. If not, show an error to the user and exit the program
		requiredArgCount.CheckForArgs(os.Args, 3)
		// check for the second provided argument to see which infra report tool needs to be executed
		switch os.Args[2] {
		// If "help" is requested as the second program argument, display the infrastructure report-specific help and exit the program
		case "help":
			infraHelp.ShowHelp()
		// If "latest" is requested as the second program argument, generate a report that overwrites the one located in "latest"
		case "latest":
			generateInfraReport.RunReport("latest", os.Args, 6)
		// If "ondemand" is requested as the second program argument, generate a report with the current date and time as a name
		case "ondemand":
			generateInfraReport.RunReport("ondemand", os.Args, 6)
		// If the script is sent a non-supported argument, show the infrastructure report-specific help and exit the program
		default:
			infraHelp.ShowHelp()
		}
	// If the first argument doesn't match any supported arguments, show the general application help and exit the program
	default:
		generalHelp.ShowHelp()
	}
}
