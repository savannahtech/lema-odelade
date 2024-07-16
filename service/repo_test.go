package service

import (
	"fmt"
	"net/http"
	"testing"

	"accessment.com/microservice/utils"
	"github.com/gin-gonic/gin"
)

func TestFetchRepo(t *testing.T) {
	var repoService RepoService
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8085/api/services/repo", nil)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	authorization := fmt.Sprintf("Bearer %s", utils.GetEnv("GITTOKEN", ""))
	req.Header.Add("Authorization", authorization)

	var ctx gin.Context
	ctx.Request = req
	repoService.GetCommits(&ctx)
	if ctx.Request.Response.StatusCode != 400 {
		t.Errorf("Unexpected response body: %v", ctx.Request.Response)
	}

}
