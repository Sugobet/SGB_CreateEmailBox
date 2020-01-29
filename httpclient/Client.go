package httpclient

import (
	"net/http"
	"net/url"
)


// Client 类型
type Client struct{
	Client *http.Client
}

// NewProxyClient 用于创建代理客户端
func (c *Client) NewProxyClient(proxyAddr string) (*http.Client, error) {
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy:http.ProxyURL(proxy),
	}

	c.Client = &http.Client{
		Transport: transport,
	}

	return c.Client, nil
}
