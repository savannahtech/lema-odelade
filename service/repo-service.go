package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	Entity "accessment.com/microservice/db/entity"
	Repository "accessment.com/microservice/db/repository"
	Dto "accessment.com/microservice/dto"
	"accessment.com/microservice/external"
	"accessment.com/microservice/utils"
	"github.com/gin-gonic/gin"
)

type RepoService struct {
}

var RepService RepoService
var repoDetailRepo Repository.RepoDetailRepository = &Repository.RepoDetailRepo{}

func (reps *RepoService) GetRepoDetails(ctx *gin.Context) {
	if len(ctx.Query("repo")) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "repo is required"})
		return
	}

	repoDetails, err := repoDetailRepo.GetByName(ctx.Query("repo"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if repoDetails.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": repoDetails})
		return
	}

	url := fmt.Sprintf("%s%s%s%s%s", utils.GitHubBaseUrl, "/repos/", utils.GetEnv("OWNER", ""), "/", utils.GetEnv("REPONAME", ""))
	response, errr := external.RestCall.ApiCall(url, nil, "GET")
	if errr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var repoDetail Dto.RepoDetail
	errT := json.Unmarshal(*response, &repoDetail)
	if errT != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if repoDetail.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Repo not found"})
		return
	}

	repo := Entity.RepoDetail{}
	repo.DateCreated = repoDetail.DateCreated
	repo.UpdatedDate = repoDetail.DateUpdated
	repo.Name = repoDetail.Name
	repo.Description = repoDetail.Description
	repo.Url = repoDetail.Url
	repo.ForkCount = repoDetail.ForkCount
	repo.StarCount = repoDetail.StarCount
	repo.OpenIssueCount = repoDetail.OpenIssueCount
	repo.WatcherCount = repoDetail.WatcherCount
	repo.Owner = repoDetail.Owner.Login

	errp := repoDetailRepo.Store(repo)
	if errp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errp.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": repo})
}

func (reps *RepoService) GetCommits(ctx *gin.Context) {
	if len(ctx.Query("repo")) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "repo is required"})
		return
	}

	var commitRepo Repository.CommitRepository = &Repository.CommitRepo{}
	commits, err := commitRepo.GetCommit(ctx.Query("repo"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(commits) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": commits})
}
