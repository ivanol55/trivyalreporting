// Sets the package name to import from the report generation functions
package datetimeString

// Imports necessary packages for the function to get dates and convert hour and minute ints to strings
import (
	"strconv"
	"time"
)

// Function that returns the current datetime string to create folders as YYYY-MM-DD-HH-mm
func GetCurrentDatetimeString() string {
	// Get date as string in YYYY-MM-DD
	var date string = time.Now().Format("2006-01-02")
	// Get hour and minute from the Clock function, convert to strings
	var hour, _, _ = time.Now().Clock()
	var hourString = strconv.Itoa(hour)
	var _, minute, _ = time.Now().Clock()
	var minuteString = strconv.Itoa(minute)
	// Format the final string with all values as YYYY-MM-DD-HH-mm, return it to the function caller
	var datetime string = date + "-" + hourString + "-" + minuteString
	return datetime
}
