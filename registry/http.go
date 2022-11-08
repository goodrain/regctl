/*
Copyright © 2022 Qi Zhang <smallqi1@163.com>

*/
package registry

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HttpRequest(method string, r RepoInfo) string {
	var client = &http.Client{}

	req, err := http.NewRequest(method, r.Url, nil)
	if err != nil {
		logrus.Error("request fail %s", err)
		return "request errors"
	}

	if r.User != "" && r.Passwd != "" {
		req.SetBasicAuth(r.User, r.Passwd)
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return "response errors"
	}
	if resp.StatusCode != 200 {
		logrus.Error(method+" "+r.Url+" ", resp.Status)
		return "Errors: response status code not 200"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("body is %s", resp.Status)
		return "errors"
	}

	bs := string(body)
	return bs
}
