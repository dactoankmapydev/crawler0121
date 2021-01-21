package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	errUnexpectedResponse = "unexpected response: %s"
)

type HTTPClient struct{}

var (
	HttpClient = HTTPClient{}
)

func RequestProxy() {
	proxyURL, _ := url.Parse("http://127.0.0.1:3130")
	http.DefaultTransport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
}

func (c HTTPClient) GetRequestVirustotal(api string) ([]byte, error) {
	//RequestProxy()
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Set("X-Apikey", os.Getenv("VIRUS_TOTAL_API_KEY"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.info(fmt.Sprintf("GET %s -> %d", api, resp.StatusCode))
	if resp.StatusCode != 200 {
		respErr := fmt.Errorf(errUnexpectedResponse, resp.Status)
		fmt.Sprintf("request failed: %v", respErr)
		return nil, respErr
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c HTTPClient) GetRequestOtx(api string) ([]byte, error) {
	//RequestProxy()
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Set("X-OTX-API-KEY", os.Getenv("OTX_API_KEY"))
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
	return ioutil.ReadAll(resp.Body)
}

func (c HTTPClient) GetRequestMirrorH(pathURL string) ([]byte, error) {
	//RequestProxy()
	req, _ := http.NewRequest("GET", pathURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
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
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c HTTPClient) info(msg string) {
	log.Printf("[client] %s\n", msg)
}
