package infraAutoReport

import (
	"fmt";
	"strings";
	"os";
    "os/exec";
)

func RunReport(args []string, requiredCount int) {
	var argumentCount int = len(args)
    if argumentCount < requiredCount {
		fmt.Println("")
		fmt.Println("Seems like some arguments are missing here! Please check for the proper 'help' usage page to learn more")
		fmt.Println("You need to provide the list of severities you want to use as comma-separated (CRITICAL,HIGH,MEDIUM,LOW) and the services you want to scan")
		fmt.Println("See the list of available services supported by trivy with `trivy help aws`")
		os.Exit(0)
    }  else {
		var region string = args[3]
		var severities string = args[4]
		var services string = args[5]
		var servicesList = strings.Split(services, ",")
		var baseCommand = "trivy aws --format template --template @src/trivy/infra/template.tpl --output ./webfiles/infra/reports/latest/index.html"
		var regionFlag = " --region " + region
		var severityFlag string = " --severity " + severities
		var serviceFlag string = ""
		for _, service := range servicesList {
    		serviceFlag = serviceFlag + " --service " + service
		}
		var command string = baseCommand + regionFlag + severityFlag + serviceFlag
		var commandObject = exec.Command("sh", "-c", command)
		fmt.Println("")
		fmt.Println("Running a report into /infra/reports/latest/ with your requested values...")
		commandObject.Run()
		fmt.Println("Report is done! it has been generated as ./webfiles/infra/reports/latest/index.html")	
    }
}