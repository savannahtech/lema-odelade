package cronjob

import (
	"encoding/json"
	"fmt"
	"time"

	Entity "accessment.com/microservice/db/entity"
	Repository "accessment.com/microservice/db/repository"
	Dto "accessment.com/microservice/dto"
	"accessment.com/microservice/external"
	"accessment.com/microservice/utils"
	"github.com/carlescere/scheduler"
)

type CronJobService struct {
}

func (cron *CronJobService) GetCommits(repoName, owner string) {

	url := fmt.Sprintf("%s%s%s%s%s", utils.GitHubBaseUrl, "/repos/", utils.GetEnv("OWNER", ""), "/", utils.GetEnv("REPONAME", ""))
	response, errr := external.RestCall.ApiCall(url, nil, "GET")
	if errr != nil {
		return
	}

	var commits []Dto.Commit
	errT := json.Unmarshal(*response, &commits)
	if errT != nil {
		return
	}

	if len(commits) == 0 {
		return
	}

	var shaList []string
	for _, element := range commits {
		shaList = append(shaList, element.Sha)
	}

	var commitRepo Repository.CommitRepository = &Repository.CommitRepo{}
	commitList, err := commitRepo.GetCommitInSha(shaList)
	if err != nil {
		return
	}

	if len(commitList) == 0 {
		commitList = []Entity.Commit{}
		for _, elem := range commits {
			var commit Entity.Commit
			commit.Author = elem.Commit.Author.Name
			commit.Message = elem.Commit.Message
			commit.Sha = elem.Sha
			commit.Url = elem.Url
			commit.Date = elem.Commit.Author.Date
			commitList = append(commitList, commit)
		}

		if len(commitList) > 0 {
			err := commitRepo.StoreList(commitList)
			if err != nil {
				return
			}
		}

		return
	}

	commitList = []Entity.Commit{}
	persistToDb := func(com Entity.Commit) {
		for _, comt := range commits {
			if comt.Sha == com.Sha {
				continue
			}
			var commit Entity.Commit
			commit.Author = comt.Commit.Author.Name
			commit.Message = comt.Commit.Message
			commit.Sha = comt.Sha
			commit.Url = comt.Url
			commit.Date = comt.Commit.Author.Date
			commitList = append(commitList, commit)
		}
	}

	for _, com := range commitList {
		persistToDb(com)
	}

	if len(commitList) == 0 {
		return
	}

	errc := commitRepo.StoreList(commitList)
	if errc != nil {
		return
	}
}

func (cron *CronJobService) UpdateCommitEveryHour() {

	job := func() {
		var repoDetailRepo Repository.RepoDetailRepository = &Repository.RepoDetailRepo{}
		repoDetails, err := repoDetailRepo.GetAll()
		if err != nil {
			return
		}

		count := 0
		for _, details := range repoDetails {
			count++
			if count == 5000 {
				//Rate limit 5000 per hour for authenticated users
				time.Sleep(1 * time.Hour)
				count = 0
			}
			go cron.GetCommits(details.Name, details.Owner)
		}
	}

	scheduler.Every(1).Hours().Run(job)
}
