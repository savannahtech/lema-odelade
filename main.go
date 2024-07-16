package main

import (
	"accessment.com/microservice/routers"
	cronjob "accessment.com/microservice/service/cron"
)

func main() {
	//Cron Job
	cron := cronjob.CronJobService{}
	go cron.UpdateCommitEveryHour()

	//Router
	routers.Routers()
}
