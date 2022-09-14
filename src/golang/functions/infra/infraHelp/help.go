package infraHelp

import (
	"fmt";
)

func ShowHelp() {
	fmt.Println("This is the infrastructure scanner help! Here's the options available for you under 'infra':")
	fmt.Println("    - 'latest [region] [Alert levels, comma-separated] [scanned services, comma-separated]' will run a report on your infrastructure and write the report to /infra/reports/latest/ on your web server.")
	fmt.Println("    - 'ondemand [region] [Alert levels, comma-separated] [scanned services, comma-separated]' will run a report on your infrastructure and put this report in a timestamped folder by service: /infra/reports/{YYYY-MM-DD-HH-MM}/{Cloud Service}")
	fmt.Println("    - 'help' will display this help page again")
	fmt.Println("")	
}
