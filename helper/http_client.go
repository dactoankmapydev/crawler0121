package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	errUnexpectedResponse = "unexpected response: %s"
)

type HTTPClient struct{}

var (
	HttpClient = HTTPClient{}
)

var backoffSchedule = []time.Duration{
	10 * time.Second,
	15 * time.Second,
	20 * time.Second,
	25 * time.Second,
	30 * time.Second,
	35 * time.Second,
	40 * time.Second,
	45 * time.Second,
	50 * time.Second,
	55 * time.Second,
	60 * time.Second,
	70 * time.Second,
	80 * time.Second,
	90 * time.Second,
	100 * time.Second,
}

func (c HTTPClient) GetVirustotal(api string) ([]byte, error) {
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Add("X-Apikey", "7d42532bd1dea1e55f7a8e99cdee23d9b26c386a6485d6dcb4106b9d055f9277")
	//proxyURL, _ := url.Parse("http://127.0.0.1:3128")
	//transport := http.Transport{
	//	//Proxy: http.ProxyURL(proxyURL),
	//}
	//client := &http.Client{
	//	Transport: &transport,
	//}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	c.info(fmt.Sprintf("GET %s -> %d", api, resp.StatusCode))
	if resp.StatusCode != 200 {
		respErr := fmt.Errorf(errUnexpectedResponse, resp.Status)
		fmt.Sprintf("request failed: %v", respErr)
		return nil, respErr
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c HTTPClient) GetVirustotalWithRetries (api string) ([]byte, error){
	var body []byte
	var err error
	for _, backoff := range backoffSchedule {
		body, err = c.GetVirustotal(api)
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stderr, "Request error: %+v\n", err)
		fmt.Fprintf(os.Stderr, "Retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// All retries failed
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c HTTPClient) GetOtx(api string) ([]byte, error) {
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Add("X-OTX-API-KEY", "779cc51038ddb07c5f6abe0832fed858a6039b9e8cdb167d3191938c1391dbba")
	//proxyURL, _ := url.Parse("http://127.0.0.1:3128")
	//transport := http.Transport{
	//	Proxy: http.ProxyURL(proxyURL),
	//}
	//client := &http.Client{
	//	Transport: &transport,
	//}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	c.info(fmt.Sprintf("GET %s -> %d", api, res.StatusCode))
	if res.StatusCode != 200 {
		respErr := fmt.Errorf(errUnexpectedResponse, res.Status)
		fmt.Sprintf("request failed: %v", respErr)
		return nil, respErr
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c HTTPClient) GetOtxWithRetries (api string) ([]byte, error){
	var body []byte
	var err error
	for _, backoff := range backoffSchedule {
		body, err = c.GetOtx(api)
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stderr, "Request error: %+v\n", err)
		fmt.Fprintf(os.Stderr, "Retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// All retries failed
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c HTTPClient) GetMirror(pathURL string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", pathURL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	//proxyURL, _ := url.Parse("http://127.0.0.1:3128")
	//transport := http.Transport{
	//	Proxy: http.ProxyURL(proxyURL),
	//}
	//client := &http.Client{
	//	Transport: &transport,
	//}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	c.info(fmt.Sprintf("GET %s -> %d", pathURL, resp.StatusCode))
	if resp.StatusCode != 200 {
		respErr := fmt.Errorf(errUnexpectedResponse, resp.Status)
		fmt.Sprintf("request failed: %v", respErr)
		return nil, respErr
	}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//return body, nil
	return resp, nil
}

func (c HTTPClient) GetMirrorWithRetries (api string) (*http.Response, error){
	//var body []byte
	var body *http.Response
	var err error
	for _, backoff := range backoffSchedule {
		body, err = c.GetMirror(api)
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stderr, "Request error: %+v\n", err)
		fmt.Fprintf(os.Stderr, "Retrying in %v\n", backoff)
		time.Sleep(backoff)
	}

	// All retries failed
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c HTTPClient) info(msg string) {
	log.Printf("[client] %s\n", msg)
}