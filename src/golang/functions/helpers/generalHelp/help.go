package generalHelp

import (
	"fmt";
)

func ShowHelp() {
	fmt.Println("This tool offers several scanning options depending on your needs. Infrastructure, repositories or docker images, we've got you covered. Here's what's available as first-level options:")
	fmt.Println("    - 'start' will run the local webserver so you can explore reports generated locally")
	fmt.Println("    - 'stop' will stop the local webserver after your'e done")
	fmt.Println("    - 'infra' will let you scan and generate reports on your currently running infrastructure. run this script with 'infra' or 'infra help' to get more information")
	fmt.Println("")	
}