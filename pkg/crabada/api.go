package crabada

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

const TeamsPrefix = "public/idle/teams"
const MinesPrefix = "public/idle/mines"

type Api interface {
	GetTeamsByAddress(address string) (*TeamsResult, error)
	GetMineHistory(address string) ([]*MineInfo, error)
	GetLastOpenMine() (*MineInfo, error)
}

type apiClient struct {
	HttpClient *fasthttp.Client
	BaseUrl    *url.URL
}

func NewApiClient(host string) (Api, error) {
	client := &fasthttp.Client{
		WriteTimeout:              time.Second * 5,
		ReadTimeout:               time.Second * 5,
		MaxIdemponentCallAttempts: 1,
	}

	baseUrl, err := url.Parse(host)

	if err != nil {
		return nil, err
	}

	return &apiClient{
		HttpClient: client,
		BaseUrl:    baseUrl,
	}, nil
}

func (c *apiClient) GetTeamsByAddress(address string) (*TeamsResult, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodGet)
	rel := &url.URL{Path: TeamsPrefix}
	rel.RawQuery = "user_address=" + address

	req.SetRequestURI(c.BaseUrl.ResolveReference(rel).String())

	resp := fasthttp.AcquireResponse()

	err := c.HttpClient.Do(req, resp)
	if err != nil {
		return nil, err
	}

	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("api error: statusCode %d body: %s", resp.StatusCode(), string(resp.Body()))
	}

	var response GetTeamsResponse
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (c *apiClient) GetMineHistory(address string) ([]*MineInfo, error) {

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodGet)
	rel := &url.URL{Path: MinesPrefix}

	query := rel.Query()
	query.Add("page", "1")
	query.Add("status", "close")
	query.Add("user_address", address)
	rel.RawQuery = query.Encode()

	req.SetRequestURI(c.BaseUrl.ResolveReference(rel).String())

	resp := fasthttp.AcquireResponse()

	err := c.HttpClient.Do(req, resp)
	if err != nil {
		return nil, err
	}

	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("api error: statusCode %d body: %s", resp.StatusCode(), string(resp.Body()))
	}

	var response GetMineHistoryResponse
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, err
	}

	return response.Result.Data, nil
}

func (c *apiClient) GetLastOpenMine() (*MineInfo, error) {

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodGet)
	rel := &url.URL{Path: MinesPrefix}

	query := rel.Query()
	query.Add("page", "1")
	query.Add("status", "open")
	query.Add("limit", "1")
	rel.RawQuery = query.Encode()

	req.SetRequestURI(c.BaseUrl.ResolveReference(rel).String())

	resp := fasthttp.AcquireResponse()

	err := c.HttpClient.Do(req, resp)
	if err != nil {
		return nil, err
	}

	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("api error: statusCode %d body: %s", resp.StatusCode(), string(resp.Body()))
	}

	var response GetMineHistoryResponse
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, err
	}

	return response.Result.Data[0], nil
}
