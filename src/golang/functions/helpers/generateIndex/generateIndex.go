// Sets the package name to import from the helper runner
package generateIndex

// Imports necessary packages for the function to print text into the console, read and write system files and folders, and work with golang templates
import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"
)

// Declare a function to run all the helpers declared below that each run a procedural task to generate an index file and save it to the system
func UpdateIndex(reportType string) {
	fmt.Println("Regenerating the index with existing reports...")
	// Get the list of existing reports and store them in `reportsList` as an array
	fmt.Println("Getting report directory list")
	var reportsList []string = getReports(reportType)
	// Get the resources scanned on each individual report (here, AWS services) and store them in relation to their report in a map
	fmt.Println("Getting scanned resources for each report")
	var reportServiceMap map[string][]string = buildReportServiceMap(reportType, reportsList)
	// Use the gathered information to generate an index based on a template and store it on its final location
	fmt.Println("Updating the report index")
	generateIndexFile(reportType, reportServiceMap)
	fmt.Println("Report index has been updated")
}

// Declare a function that gets all the existing reports on the filesystem by folder name for the given trivyalreporting report kind (infra? containers?)
func getReports(reportType string) []string {
	// Set the reports folder we want to scan based on reportType
	var targetDir string = "webfiles/" + reportType + "/reports/"
	// Read all existing folders on the target directory into a reports object
	reports, _ := ioutil.ReadDir(targetDir)
	// Prepare an array to store just the folder names, to get rid of unnecessary info and sanitize the returned data
	var reportsList []string
	// Iterate over each folder object in the reports struct and store the folder name in `reportsList`
	for _, reportDirectoryObject := range reports {
		var reportDirectory string = reportDirectoryObject.Name()
		// Checks if a filename contains a dot (hidden files, files with extensions)
		hasExtension, _ := regexp.MatchString(`.*[.].*`, reportDirectory)
		// If the filename doesn't have a dot, it's a folder we want to add, so we consider it a report and add it to the array
		if hasExtension == false {
			reportsList = append(reportsList, reportDirectory)
		}
	}
	// Return all folder names in an array to the parent function
	return reportsList
}

// Declare a function to build a map with all the reports and their scanned elements to pass into the template parser
func buildReportServiceMap(reportType string, reportsList []string) map[string][]string {
	// Prepare a map object and initialize it, with object key as a string, and object value as an array of strings
	var reportsElementMap map[string][]string
	reportsElementMap = make(map[string][]string)
	// For each report name on the list passed to the function, read its children elements and store them in the map with their parent report
	for _, report := range reportsList {
		// Prepare the target report directory for this iteration based on report type (infra? containers?) and report name
		var reportDir string = "webfiles/" + reportType + "/reports/" + report + "/"
		// Read the directory contents to get all the children elements of this report
		reportElementsObject, _ := ioutil.ReadDir(reportDir)
		// Prepare an array to store the element list we want to store for a particular report
		var reportElementsList []string
		// For each report child element, add it to the map array value to prepare the final array for that report
		for _, reportElement := range reportElementsObject {
			var reportElementDirectory string = reportElement.Name()
			// Checks if a filename contains a dot (hidden files, files with extensions)
			hasExtension, _ := regexp.MatchString(`.*[.].*`, reportElementDirectory)
			// If the filename doesn't have a dot, it's a folder we want to add, so we consider it a report element and add it to the array
			if hasExtension == false {
				reportElementsList = append(reportElementsList, reportElementDirectory)
			}
		}
		// Add the reports elements array to the report elements map with the report name as the key for future classification
		reportsElementMap[report] = reportElementsList
	}
	// When iteration over all reports is complete, return the finished map object
	return reportsElementMap
}

// Declare a function that generates an index based on a template provided the report map for the target report type
func generateIndexFile(reportType string, reportServiceMap map[string][]string) {
	// Create the template object that we will use to generate the index
	var templateObject = template.New("index.tpl")
	// Prepare a slice of templates we want to execute and add the desired template to this array. This is the required format for `ParseFiles``
	var indexTemplates []string
	indexTemplates = append(indexTemplates, "src/trivy/helpers/index.tpl")
	// Parse the template to ensure it is properly read and ready
	var templateParse, templateParse_errors = templateObject.ParseFiles(indexTemplates...)
	// Prepare a handler for the template execution, run error checking for the template
	var templateHandler = template.Must(templateParse, templateParse_errors)
	// Prepare a struct used to pass data down to the template that we'll use to fill in values. We send in the report type (infra? containers?) and the previously built report mapping
	type templateDataStruct struct {
		ReportType string
		ReportMap  map[string][]string
	}
	var templateData = templateDataStruct{
		ReportType: reportType,
		ReportMap:  reportServiceMap,
	}
	// Set the target path where the index for this report type will be stored and create the target file or truncate it if it doesn't exist
	var targetPath = "webfiles/" + reportType + "/reports/index.html"
	targetFile, _ := os.Create(targetPath)
	// Execute the template and store it on the target file we set earlier
	templateHandler.Execute(targetFile, templateData)
}
