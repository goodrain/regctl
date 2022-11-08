/*
Copyright © 2022 Qi Zhang <smallqi1@163.com>

*/
package registry

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type RepoInfo struct {
	Url    string
	Repo   string
	User   string
	Passwd string
}

func GetRepoList(r RepoInfo) string {
	var ri = RepoInfo{
		Url:    r.Url + "/v2/_catalog",
		User:   r.User,
		Passwd: r.Passwd,
	}
	rs := HttpRequest("GET", ri)
	return rs
}

func GetRepoTags(r RepoInfo) string {
	var ri = RepoInfo{
		Url:    r.Url + "/v2/" + r.Repo + "/tags/list",
		User:   r.User,
		Passwd: r.Passwd,
	}
	rs := HttpRequest("GET", ri)
	return rs
}

func GetRepoManifest(r RepoInfo) []string {
	var ri = RepoInfo{
		User:   r.User,
		Passwd: r.Passwd,
	}
	grtResult := GetRepoTags(r)
	tags := map[string]interface{}{}
	err := json.Unmarshal([]byte(grtResult), &tags)
	if err != nil {
		logrus.Error(err)
	}
	var rs []string
	for k, v := range tags {
		if k == "tags" {
			for _, vv := range v.([]interface{}) {
				ri.Url = r.Url + "/v2/" + r.Repo + "/manifests/" + vv.(string)
				rs = append(rs, HttpRequest("GET", ri))
			}
		}
	}
	return rs
}
