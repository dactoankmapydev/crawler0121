package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	errUnexpectedResponse = "unexpected response: %s"
)

type HTTPClient struct{}

var (
	HttpClient = HTTPClient{}
)

var timeout = time.Duration(10 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func (c HTTPClient) GetRequestVirustotal(api string) ([]byte, error) {
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Add("X-Apikey", "7d42532bd1dea1e55f7a8e99cdee23d9b26c386a6485d6dcb4106b9d055f9277")
	proxyURL, _ := url.Parse("http://127.0.0.1:3128")
	transport := http.Transport{
		Dial:  dialTimeout,
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: &transport,
	}
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

func (c HTTPClient) GetRequestOtx(api string) ([]byte, error) {
	req, _ := http.NewRequest("GET", api, nil)

	req.Header.Add("X-OTX-API-KEY", "779cc51038ddb07c5f6abe0832fed858a6039b9e8cdb167d3191938c1391dbba")
	proxyURL, _ := url.Parse("http://127.0.0.1:3128")
	transport := http.Transport{
		Dial:  dialTimeout,
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: &transport,
	}
	//client := &http.Client{}
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
	return ioutil.ReadAll(res.Body)
}

func (c HTTPClient) GetRequestMirrorH(pathURL string) ([]byte, error) {
	req, _ := http.NewRequest("GET", pathURL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	/*transport := http.Transport{
		Dial: dialTimeout,
	}
	client := &http.Client{
		Transport: &transport,
	}*/
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
