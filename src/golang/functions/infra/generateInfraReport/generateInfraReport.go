package generateInfraReport

import (
	"fmt";
	"trivyalreporting/src/golang/functions/helpers/datetimeString";
	"strings";
	"os";
    "os/exec";
	"bytes"
	"io/ioutil"
)

func RunReport(reportKind string, args []string, requiredCount int) {
	var argumentCount int = len(args)
    if argumentCount < requiredCount {
		fmt.Println("")
		fmt.Println("Seems like some arguments are missing here! Please check for the proper 'help' usage page to learn more")
		fmt.Println("You need to provide the list of severities you want to use as comma-separated (CRITICAL,HIGH,MEDIUM,LOW) and the services you want to scan")
		fmt.Println("See the list of available services supported by trivy with `trivy help aws`")
		os.Exit(0)
    }  else {
		fmt.Println("Preparing to run reports...")	
		var datetime = chooseDatetime(reportKind)
		var region string = args[3]
		var severities string = args[4]
		var severitiesList = strings.Split(severities, ",")
		var services string = args[5]
		var servicesList = strings.Split(services, ",")
		fmt.Println("Creating necessary directories...")	
		createDirs(datetime, servicesList)
		fmt.Println("Generating report blocks based on templates...")	
		generateStaticBlocks(datetime, region)
		cleanupCache(datetime)
		generateDynamicBlocks(datetime, servicesList, severitiesList, region)
		fmt.Println("Finished reports iteration as blocks. Splicing into finished reports...")
		spliceBlocksIntoReports(datetime, servicesList, severitiesList)
		spliceBlocksIntoReports(datetime, servicesList, severitiesList)
		fmt.Println("Cleaning up temporary files...")
		cleanupTempFiles(datetime)
		fmt.Println("Reports are ready! You can find them available on ./webfiles/infra/reports/" + datetime + "/[service]/")
    }
}

func chooseDatetime(reportKind string) string {
    if reportKind == "latest" {
		var datetime string = "latest"
		return datetime
    }  else {
		var datetime string = datetimeString.GetCurrentDatetimeString()
		return datetime
    }
}

func createDirs(datetime string, services []string) {
	var tmpdir string = "/tmp/trivyalreporting/"
	os.Mkdir(tmpdir, 0755)
	tmpdir = "/tmp/trivyalreporting/infra/"
	os.Mkdir(tmpdir, 0755)
	os.Mkdir(tmpdir + datetime, 0755)
	var finaldir string = "webfiles/infra/reports/"
	os.Mkdir(finaldir + datetime, 0755)
	for _, service := range services {
		os.Mkdir(tmpdir + datetime + "/" + service, 0755)
		os.Mkdir(finaldir + datetime + "/" + service, 0755)
	}
}

func generateStaticBlocks(datetime string, region string) {
	var headerCommand string = "trivy aws --cache-dir /tmp/trivyalreporting/infra/" + datetime + "/cache --service s3 --severity CRITICAL --format template --template @src/trivy/infra/header.tpl --output /tmp/trivyalreporting/infra/" + datetime + "/header.html "
	var footerCommand string = "trivy aws --cache-dir /tmp/trivyalreporting/infra/" + datetime + "/cache --service s3 --severity CRITICAL --format template --template @src/trivy/infra/footer.tpl --output /tmp/trivyalreporting/infra/" + datetime + "/footer.html "
	var regionFlag string = " --region " + region
	var command string = headerCommand + regionFlag
	var commandObject = exec.Command("sh", "-c", command)
	commandObject.Run()
	command = footerCommand + regionFlag
	commandObject = exec.Command("sh", "-c", command)
	commandObject.Run()
}

func generateDynamicBlocks(datetime string, services []string, severities []string, region string) {
	for _, service := range services {
		fmt.Println("Running reports for service " + service + " for AWS region " + region + "..." )
		for _, severity := range severities {
			var baseCommand string = "trivy aws --cache-dir /tmp/trivyalreporting/infra/" + datetime + "/cache --format template --template @src/trivy/infra/misconfigurations.tpl"
			var regionFlag string = " --region " + region
			var serviceFlag string = " --service " + service
			var severityFlag string = " --severity " + severity
			var outputFlag string = " --output /tmp/trivyalreporting/infra/" + datetime + "/" + service + "/index-" + severity + ".html"
			var command string = baseCommand + regionFlag + severityFlag + serviceFlag + outputFlag
			var commandObject = exec.Command("sh", "-c", command)
			commandObject.Run()
		}
	}
}

func spliceBlocksIntoReports(datetime string, services []string, severities []string) {
	var headerFile string = "/tmp/trivyalreporting/infra/" + datetime + "/header.html"
	var footerFile string = "/tmp/trivyalreporting/infra/" + datetime + "/footer.html"

	for _, service := range services {
		var reportBlockList []string
		reportBlockList = append(reportBlockList, headerFile)
		for _, severity := range severities {
			var reportBlockFile string = "/tmp/trivyalreporting/infra/" + datetime + "/" + service + "/index-" + severity + ".html"
			reportBlockList = append(reportBlockList, reportBlockFile)
		}
		reportBlockList = append(reportBlockList, footerFile)
		combineFiles(reportBlockList, datetime, service)
	}
}

func combineFiles (reportBlockList []string, datetime string, service string) {
	var finalBuffer bytes.Buffer
	for _, reportBlockFile := range reportBlockList {
		reportBlockBuffer, _ := ioutil.ReadFile(reportBlockFile)
		finalBuffer.Write(reportBlockBuffer)
	}
	var outputFile string = "webfiles/infra/reports/" + datetime + "/" + service + "/index.html"
	ioutil.WriteFile(outputFile, finalBuffer.Bytes(), 0644)
}

func cleanupTempFiles(datetime string) {
	var tmpdir string = "/tmp/trivyalreporting/infra/" + datetime
	os.RemoveAll(tmpdir)
}

func cleanupCache(datetime string) {
	var tmpdir string = "/tmp/trivyalreporting/infra/" + datetime + "/cache/"
	os.RemoveAll(tmpdir)
}