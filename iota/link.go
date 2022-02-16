package iota

import (
	IotaApI "github.com/loveandpeople-DAG/goClient/api"
)

var Lin *Link

type Link struct {
	API *IotaApI.API
}

func InitLink(NodeUrl, User, Pwd string) (*Link, error) {
	Lin = &Link{}
	api, err := IotaApI.ComposeAPI(IotaApI.HTTPClientSettings{
		URI:  NodeUrl,
		User: User,
		Pwd:  Pwd,
	})
	if err != nil {
		return nil, err
	}
	Lin.API = api
	return Lin, nil
}
