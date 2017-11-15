package greynoise

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const greynoise_api = "http://api.greynoise.io:8888"
const query_ip_uri = greynoise_api + "/v1/query/ip"
const list_tags_uri = greynoise_api + "/v1/query/list"
const query_tag_uri = greynoise_api + "/v1/query/tag"

type httpClient interface {
	Post(string, string, io.Reader) (*http.Response, error)
}

type Greynoise struct {
	client httpClient
}

type IPQueryResponse struct {
	IP      string     `json:"ip"`
	Status  string     `json:"status"`
	Records []IPRecord `json:"records"`
}

type IPRecord struct {
	Name        string    `json:"name"`
	FirstSeen   time.Time `json:"first_seen"`
	LastUpdated time.Time `json:"last_updated"`
	Confidence  string    `json:"confidence"`
	Intention   string    `json:"intention"`
	Category    string    `json:"category"`
	IP          string    `json:"ip,omitempty"`
}

func (g *Greynoise) QueryIP(ip string) (IPQueryResponse, error) {
	buf := strings.NewReader(fmt.Sprintf("ip=%s", ip))
	resp, err := g.client.Post(query_ip_uri, "application/json", buf)
	if err != nil {
		return IPQueryResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IPQueryResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return IPQueryResponse{}, errors.New(string(body))
	}
	var ipResponse IPQueryResponse
	err = json.Unmarshal(body, &ipResponse)
	return ipResponse, err
}

type TagsResponse struct {
	Status string   `json:"status"`
	Tags   []string `json:"tags"`
}

func (g *Greynoise) ListTags() (TagsResponse, error) {
	resp, err := g.client.Post(list_tags_uri, "application/json", bytes.NewReader(nil))
	if err != nil {
		return TagsResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return TagsResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return TagsResponse{}, errors.New(string(body))
	}
	var tagResponse TagsResponse
	err = json.Unmarshal(body, &tagResponse)
	return tagResponse, err
}

type QueryTagResponse struct {
	Status  string     `json:"status"`
	Tag     string     `json:"tag"`
	Records []IPRecord `json:"records"`
}

func (g *Greynoise) QueryTag(tag string) (QueryTagResponse, error) {
	buf := strings.NewReader(fmt.Sprintf("tag=%s", tag))
	resp, err := g.client.Post(query_tag_uri, "application/json", buf)
	if err != nil {
		return QueryTagResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return QueryTagResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return QueryTagResponse{}, errors.New(string(body))
	}
	var tagResponse QueryTagResponse
	err = json.Unmarshal(body, &tagResponse)
	return tagResponse, err
}
