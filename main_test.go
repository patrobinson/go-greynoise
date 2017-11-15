package greynoise

import (
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

type MockQueryIPHttpClient struct{}

func (m *MockQueryIPHttpClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	rBody := strings.NewReader(`{"ip":"198.20.69.74","status":"exists","records":[{"name":"VOIP Scanner","first_seen":"2017-09-27T15:20:20.253Z","last_updated":"2017-09-27T15:20:20.253Z","confidence":"low","intention":"","category":"activity"},{"name":"Elasticsearch Scanner","first_seen":"2017-09-27T15:20:20.253Z","last_updated":"2017-09-27T15:20:20.253Z","confidence":"low","intention":"","category":"activity"}]}`)
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(rBody),
	}

	return resp, nil
}

func TestQueryIP(t *testing.T) {
	greynoise := Greynoise{
		client: &MockQueryIPHttpClient{},
	}
	timeSeen, _ := time.Parse(time.RFC3339Nano, "2017-09-27T15:20:20.253Z")
	resp, err := greynoise.QueryIP("1.2.3.4")
	if err != nil {
		t.Errorf("Got error calling QueryIP: %s", err)
	}
	expectedResp := IPQueryResponse{
		IP:     "198.20.69.74",
		Status: "exists",
		Records: []IPRecord{
			IPRecord{
				Name:        "VOIP Scanner",
				FirstSeen:   timeSeen,
				LastUpdated: timeSeen,
				Confidence:  "low",
				Intention:   "",
				Category:    "activity",
			},
			IPRecord{
				Name:        "Elasticsearch Scanner",
				FirstSeen:   timeSeen,
				LastUpdated: timeSeen,
				Confidence:  "low",
				Intention:   "",
				Category:    "activity",
			},
		},
	}
	if !reflect.DeepEqual(resp, expectedResp) {
		t.Errorf("Expected %s\nGot %s", expectedResp, resp)
	}
}

type MockListTagsHttpClient struct{}

func (m *MockListTagsHttpClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	rBody := strings.NewReader(`{"status":"ok","tags":["VNC_SCANNER_HIGH","PING_SCANNER_LOW","JBOSS_WORM","GOOGLEBOT","SNMP_SCANNER_LOW","CPANEL_SCANNER_LOW","WINRM_SCANNER_HIGH","MASSCAN_CLIENT","REDIS_SCANNER_HIGH","RABBITMQ_SCANNER_HIGH"]}`)
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(rBody),
	}

	return resp, nil
}

func TestListTags(t *testing.T) {
	greynoise := Greynoise{
		client: &MockListTagsHttpClient{},
	}
	resp, err := greynoise.ListTags()
	if err != nil {
		t.Errorf("Got error calling QueryIP: %s", err)
	}
	expectedResp := TagsResponse{
		Status: "ok",
		Tags: []string{
			"VNC_SCANNER_HIGH",
			"PING_SCANNER_LOW",
			"JBOSS_WORM",
			"GOOGLEBOT",
			"SNMP_SCANNER_LOW",
			"CPANEL_SCANNER_LOW",
			"WINRM_SCANNER_HIGH",
			"MASSCAN_CLIENT",
			"REDIS_SCANNER_HIGH",
			"RABBITMQ_SCANNER_HIGH",
		},
	}
	if !reflect.DeepEqual(resp, expectedResp) {
		t.Errorf("Expected %s\nGot %s", expectedResp, resp)
	}
}

type MockQueryTagHttpClient struct{}

func (m *MockQueryTagHttpClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	rBody := strings.NewReader(`{"tag":"YANDEX_SEARCH_ENGINE","status":"ok","records":[{"ip":"5.255.250.2","name":"Yandex Search Engine","first_seen":"2017-09-27T15:20:20.253Z","last_updated":"2017-09-27T15:20:20.253Z","confidence":"high","Intention":"benign","category":"search_engine"},{"ip":"5.255.250.6","name":"Yandex Search Engine","first_seen":"2017-09-27T15:20:20.253Z","last_updated":"2017-09-27T15:20:20.253Z","confidence":"high","Intention":"benign","category":"search_engine"}]}`)
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(rBody),
	}

	return resp, nil
}

func TestQueryTag(t *testing.T) {
	greynoise := Greynoise{
		client: &MockQueryTagHttpClient{},
	}
	timeSeen, _ := time.Parse(time.RFC3339Nano, "2017-09-27T15:20:20.253Z")
	resp, err := greynoise.QueryTag("YANDEX_SEARCH_ENGINE")
	if err != nil {
		t.Errorf("Got error calling QueryIP: %s", err)
	}
	expectedResp := QueryTagResponse{
		Tag:    "YANDEX_SEARCH_ENGINE",
		Status: "ok",
		Records: []IPRecord{
			IPRecord{
				IP:          "5.255.250.2",
				Name:        "Yandex Search Engine",
				FirstSeen:   timeSeen,
				LastUpdated: timeSeen,
				Confidence:  "high",
				Intention:   "benign",
				Category:    "search_engine",
			},
			IPRecord{
				IP:          "5.255.250.6",
				Name:        "Yandex Search Engine",
				FirstSeen:   timeSeen,
				LastUpdated: timeSeen,
				Confidence:  "high",
				Intention:   "benign",
				Category:    "search_engine",
			},
		},
	}
	if !reflect.DeepEqual(resp, expectedResp) {
		t.Errorf("Expected %s\nGot %s", expectedResp, resp)
	}
}
