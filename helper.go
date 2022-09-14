package main

import (
	"os"
	"trivyalreporting/src/golang/functions/helpers/greeting";
	"trivyalreporting/src/golang/functions/helpers/requiredArgCount";
	"trivyalreporting/src/golang/functions/helpers/generalHelp";
	"trivyalreporting/src/golang/functions/helpers/manageWebserver";
	"trivyalreporting/src/golang/functions/infra/infraHelp";
	"trivyalreporting/src/golang/functions/infra/infraAutoReport";
	"trivyalreporting/src/golang/functions/infra/infraOnDemandReport";
)

func main() {
	greeting.ShowGreeting()
	requiredArgCount.CheckForArgs(os.Args, 2)
	switch os.Args[1] {
	case "help":
		generalHelp.ShowHelp()
	case "start":
		manageWebserver.StartStopReportsServer("start")
	case "stop":
		manageWebserver.StartStopReportsServer("stop")
	case "infra":
		requiredArgCount.CheckForArgs(os.Args, 3)
		switch os.Args[2] {
		case "help":
			infraHelp.ShowHelp()
		case "latest":
			infraAutoReport.RunReport(os.Args, 6)
		case "ondemand":
			infraOnDemandReport.RunReport(os.Args, 6)
		default:
			infraHelp.ShowHelp()
		}


	default:
		generalHelp.ShowHelp()
	}
}
