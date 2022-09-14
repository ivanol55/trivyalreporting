package infraOnDemandReport

import (
	"fmt";
	"trivyalreporting/src/golang/functions/helpers/datetimeString";
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
		var datetime string = datetimeString.GetCurrentDatetimeString()
		var region string = args[3]
		var severities string = args[4]
		var services string = args[5]
		var servicesList = strings.Split(services, ",")
		var baseCommand = "trivy aws --format template --template @src/trivy/infra/template.tpl"
		var regionFlag = " --region " + region
		var severityFlag string = " --severity " + severities
		os.Mkdir("webfiles/infra/reports/" + datetime, 0755)
		for _, service := range servicesList {
			var serviceFlag string = " --service " + service
			var outputFlag string = " --output ./webfiles/infra/reports/" + datetime + "/" + service + "/index.html"
			os.Mkdir("webfiles/infra/reports/" + datetime + "/" + service, 0755)
			var command string = baseCommand + regionFlag + severityFlag + serviceFlag + outputFlag
			var commandObject = exec.Command("sh", "-c", command)
			fmt.Println("")
			fmt.Println("Running a report into /infra/reports/" + datetime + "/" + service + "/ with your requested values...")
			commandObject.Run()
			fmt.Println("Report for service " + service + " is done! it has been generated as ./webfiles/infra/reports/"+ datetime + "/" + service + "/index.html")
		}
		fmt.Println("Finished reports iteration.")	
    }
}