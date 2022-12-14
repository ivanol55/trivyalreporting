// Sets the package name to import from the helper runner
package infraHelp

// Imports necessary packages for the function to print text into the console
import (
	"fmt"
)

// Declare a function to show help for the infrastructure commandlet
func ShowHelp() {
	fmt.Println("This is the infrastructure scanner help! Here's the options available for you under 'infra':")
	fmt.Println("    - 'latest [region] [Alert levels, comma-separated] [scanned services, comma-separated]' will run a report on your infrastructure and write the report to ./webfiles/infra/reports/latest/{Cloud Service}.")
	fmt.Println("    - 'ondemand [region] [Alert levels, comma-separated] [scanned services, comma-separated]' will run a report on your infrastructure and put this report in a timestamped folder by service: ./webfiles/infra/reports/{YYYY-MM-DD-HH-MM}/{Cloud Service}")
	fmt.Println("    - 'help' will display this help page again")
	fmt.Println("")
}
