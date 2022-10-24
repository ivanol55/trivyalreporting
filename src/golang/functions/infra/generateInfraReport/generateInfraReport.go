// Sets the package name to import from the helper runner
package generateInfraReport

// Imports necessary packages for the function to:
// `bytes`: Used to create file buffers and store them as objects
// `fmt`: Print text into the terminal and run OS commands
// `io/ioutil`: reads writes files to the system storage
// `os`: Run host system tasks like folder creation
// `os/exec`: Execute host system commands
// `strings`: String manipulation, here splitting strings into arrays by key character
// `trivyalreporting/src/golang/functions/helpers/datetimeString`: Retrieve a datetime string
// `trivyalreporting/src/golang/functions/helpers/generateIndex`: Update the report index after every execution

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"trivyalreporting/src/golang/functions/helpers/datetimeString"
	"trivyalreporting/src/golang/functions/helpers/generateIndex"
)

// Main function that runs the reports into HTML format, calls the rest of the helpers in order
// Receives the report kind expected, the arguments array and the required argument count
func RunReport(reportKind string, args []string, requiredCount int) {
	// Gets the number of arguments sent into the script
	var argumentCount int = len(args)
	// If the argument count is not the required one, print a message to inform about it and exit gracefully
	if argumentCount < requiredCount {
		fmt.Println("")
		fmt.Println("Seems like some arguments are missing here! Please check for the proper 'help' usage page to learn more")
		fmt.Println("You need to provide the list of severities you want to use as comma-separated (CRITICAL,HIGH,MEDIUM,LOW) and the services you want to scan")
		fmt.Println("See the list of available services supported by trivy with `trivy help aws`")
		os.Exit(0)
	} else {
		// If the argument count is correct, run the report
		fmt.Println("Preparing to run reports...")
		//Get a date and time string from the chooseDatetime helper function, sending down the reportKind (can be either 'latest' or 'ondemand')
		var datetime = chooseDatetime(reportKind)
		// Store the necessary values from the arguments array received from the helper
		// In the case of severities and services, split by commas into an array
		var region string = args[3]
		var severities string = args[4]
		var severitiesList = strings.Split(severities, ",")
		var services string = args[5]
		var servicesList = strings.Split(services, ",")
		// Create all necessary directories to store reports, both temporary and final
		fmt.Println("Creating necessary directories...")
		createDirs(datetime, servicesList)
		// Generate the report blocks for header and footer (from staticBlocks), and dynamic per-service and per-priority blocks (DynamicBlocks)
		fmt.Println("Generating report blocks based on templates...")
		generateStaticBlocks(datetime, region)
		generateDynamicBlocks(datetime, servicesList, severitiesList, region)
		// Compress report blocks into single HTML files and put them into their final locations
		fmt.Println("Finished reports iteration as blocks. Splicing into finished reports...")
		spliceBlocksIntoReports(datetime, servicesList, severitiesList)
		// Clean up temporary directories to clean out cache and stray report blocks
		fmt.Println("Cleaning up temporary files...")
		cleanupTempFiles(datetime)
		fmt.Println("Reports are ready! You can find them available on ./webfiles/infra/reports/" + datetime + "/[service]/")
		// Updates the index file for the infrastructure reports reference
		generateIndex.UpdateIndex("infra")
	}
}

// Declares a middleware function to choose the folder name based on the report Kind sent down from the helper, and returns it as a result
func chooseDatetime(reportKind string) string {
	if reportKind == "latest" {
		// If the reportKind is `latest`, report folder is latest and overwrites the last report
		var datetime string = "latest"
		return datetime
	} else {
		// If the reportKind is not `latest` a datetime string in the form of YYYY-MM-DD-HH-mm is retrieved from the datetimeString helper package
		var datetime string = datetimeString.GetCurrentDatetimeString()
		return datetime
	}
}

// Declares a function to create all necessary directories, both temporary and final
func createDirs(datetime string, services []string) {
	// Creates the base temporary folder in case it doesn't exist yet
	var tmpdir string = "/tmp/trivyalreporting/"
	os.Mkdir(tmpdir, 0755)
	// Creates the second level directory in case it doesn't exist yet, this is to separate infra from other types of reports
	tmpdir = "/tmp/trivyalreporting/infra/"
	os.Mkdir(tmpdir, 0755)
	// Creates the report target folder for the temporary files based on the report datetime value
	os.Mkdir(tmpdir+datetime, 0755)
	// Creates the base target directory where this report will be stored
	var finaldir string = "webfiles/infra/reports/"
	os.Mkdir(finaldir+datetime, 0755)
	// For each service, it creates a sub-folder to store the final report
	for _, service := range services {
		os.Mkdir(tmpdir+datetime+"/"+service, 0755)
		os.Mkdir(finaldir+datetime+"/"+service, 0755)
	}
}

// Declares a function that generates the static report blocks, in this case, the header and footer for each HTML report
func generateStaticBlocks(datetime string, region string) {
	// Prepares the base command strings to run
	var headerCommand string = "trivy aws --cache-dir /tmp/trivyalreporting/infra/" + datetime + "/cache --service s3 --severity CRITICAL --format template --template @src/trivy/infra/header.tpl --output /tmp/trivyalreporting/infra/" + datetime + "/header.html "
	var footerCommand string = "trivy aws --cache-dir /tmp/trivyalreporting/infra/" + datetime + "/cache --service s3 --severity CRITICAL --format template --template @src/trivy/infra/footer.tpl --output /tmp/trivyalreporting/infra/" + datetime + "/footer.html "
	// Creates the region flag for the command string
	var regionFlag string = " --region " + region
	// Prepares and runs the header generation command
	var command string = headerCommand + regionFlag
	var commandObject = exec.Command("sh", "-c", command)
	commandObject.Run()
	// Prepares and runs the footer generation command
	command = footerCommand + regionFlag
	commandObject = exec.Command("sh", "-c", command)
	commandObject.Run()
}

// Declares a function that runs a report for each service and specified severity in independent HTML blocks
func generateDynamicBlocks(datetime string, services []string, severities []string, region string) {
	// Iterate over each service in the service list passed in as a variable
	for _, service := range services {
		// Communicate which report is being generated
		fmt.Println("Running reports for service " + service + " for AWS region " + region + "...")
		// For each severity, in the service for this loop execution
		for _, severity := range severities {
			// Generate the base command that will be run, setting the output location and the template that will be used
			var baseCommand string = "trivy aws --cache-dir /tmp/trivyalreporting/infra/" + datetime + "/cache --format template --template @src/trivy/infra/misconfigurations.tpl"
			// Prepare variable strings for selected region, current service, current severity for the loop and the output location of the report block
			var regionFlag string = " --region " + region
			var serviceFlag string = " --service " + service
			var severityFlag string = " --severity " + severity
			var outputFlag string = " --output /tmp/trivyalreporting/infra/" + datetime + "/" + service + "/index-" + severity + ".html"
			// Build the command with all the finished flags and execute it to save the report into a known location
			var command string = baseCommand + regionFlag + severityFlag + serviceFlag + outputFlag
			var commandObject = exec.Command("sh", "-c", command)
			commandObject.Run()
		}
	}
}

// Declares a function to generate a list of all report severity blocks into per-service lists and passes it down to the `combineFiles` function
func spliceBlocksIntoReports(datetime string, services []string, severities []string) {
	// Declare target locations of header and footer files
	var headerFile string = "/tmp/trivyalreporting/infra/" + datetime + "/header.html"
	var footerFile string = "/tmp/trivyalreporting/infra/" + datetime + "/footer.html"
	// Loop used to generate a report file location entry for each requested service
	for _, service := range services {
		// Initialize a report block variable for report locations
		var reportBlockList []string
		// Add the header file location to the array
		reportBlockList = append(reportBlockList, headerFile)
		// For each registered severity, add the service report block location to the array
		for _, severity := range severities {
			var reportBlockFile string = "/tmp/trivyalreporting/infra/" + datetime + "/" + service + "/index-" + severity + ".html"
			reportBlockList = append(reportBlockList, reportBlockFile)
		}
		// Add the footer file as the last report block list element to close the document
		reportBlockList = append(reportBlockList, footerFile)
		// call the `combineFiles` function to generate the output files by splicing contents of all reports for a service
		combineFiles(reportBlockList, datetime, service)
	}
}

// Declares a function that, given an array of file locations, generates finished service reports and splices them into a single file
func combineFiles(reportBlockList []string, datetime string, service string) {
	// Create a mmeory buffer to store the final file contents in
	var finalBuffer bytes.Buffer
	// For every file in the report block list, read its contents and append them to the file buffer
	for _, reportBlockFile := range reportBlockList {
		reportBlockBuffer, _ := ioutil.ReadFile(reportBlockFile)
		finalBuffer.Write(reportBlockBuffer)
	}
	// Set the final location of the requested report based on the datetime value and service, and save the buffer contents into that destination
	var outputFile string = "webfiles/infra/reports/" + datetime + "/" + service + "/index.html"
	ioutil.WriteFile(outputFile, finalBuffer.Bytes(), 0644)
}

// Declare a function to clean up unnecessary files after the report has been finalized
func cleanupTempFiles(datetime string) {
	// Declare the location to delete based on the datetime variable, then remove it recursively
	var tmpdir string = "/tmp/trivyalreporting/infra/" + datetime
	os.RemoveAll(tmpdir)
}
