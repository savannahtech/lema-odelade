package external

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"accessment.com/microservice/utils"
)

type GitRestApiCall struct {
}

var RestCall GitRestApiCall

func (restcal *GitRestApiCall) ApiCall(endPoint string, data *bytes.Buffer, method string) (*[]byte, error) {

	url := fmt.Sprintf("%s%s", utils.GitHubBaseUrl, endPoint)
	req, err1 := http.NewRequest(method, url, data)
	if err1 != nil {
		return nil, errors.New(err1.Error())
	}

	// set the request header Content-Type for json
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	authorization := fmt.Sprintf("Bearer %s", utils.GetEnv("GITACCESSTOKEN", ""))
	req.Header.Add("Authorization", authorization)
	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		return nil, errors.New(err2.Error())
	}

	defer resp.Body.Close()
	respBody, nErr := io.ReadAll(resp.Body)

	if nErr != nil {
		return nil, errors.New(nErr.Error())
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return &respBody, nil
}

type RestCallService interface {
	ApiCall(endPoint string, data *bytes.Buffer, method string) (*[]byte, error)
}
