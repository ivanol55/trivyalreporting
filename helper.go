package main

import (
	"os"
	"trivyalreporting/src/golang/functions/helpers/generalHelp"
	"trivyalreporting/src/golang/functions/helpers/greeting"
	"trivyalreporting/src/golang/functions/helpers/requiredArgCount"
	"trivyalreporting/src/golang/functions/infra/generateInfraReport"
	"trivyalreporting/src/golang/functions/infra/infraHelp"
)

func main() {
	greeting.ShowGreeting()
	requiredArgCount.CheckForArgs(os.Args, 2)
	switch os.Args[1] {
	case "help":
		generalHelp.ShowHelp()
	case "infra":
		requiredArgCount.CheckForArgs(os.Args, 3)
		switch os.Args[2] {
		case "help":
			infraHelp.ShowHelp()
		case "latest":
			generateInfraReport.RunReport("latest", os.Args, 6)
		case "ondemand":
			generateInfraReport.RunReport("ondemand", os.Args, 6)
		default:
			infraHelp.ShowHelp()
		}

	default:
		generalHelp.ShowHelp()
	}
}
