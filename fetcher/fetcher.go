package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	// 直接使用封装的 Get 方法可能会出现403错误，原因是访问的url需要设置 header
	// 解决方案是自定义创建一个 client 去进行 request
	//resp, err := http.Get(url)

	// 手动创建一个 client
	client := http.DefaultClient
	// 设置代理
	//client := proxyHttpClient("http://127.0.0.1:8000")
	// 手动创建一个 request
	req, err := http.NewRequest("GET", url, nil)
	// 错误处理
	if err != nil {
		return nil, err
	}
	// 设置 Header
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36")
	// 发起请求，得到 Response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 自定义一个 error 可以使用 errors.New() or fmt.Errorf()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	// 处理202等待请求响应
	if resp.StatusCode == http.StatusAccepted {
		time.Sleep(3 * time.Millisecond)
	}
	// 读取 1024 byte 进行编码猜测
	bodyReader := bufio.NewReader(resp.Body)
	e, err := determineEncoding(bodyReader)
	if err != nil {
		return nil, fmt.Errorf("read 1024 bytes error %s", err)
	}
	// 编码转换包 golang.org/x/text
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	// 返回读取结果 utf-8 编码
	return ioutil.ReadAll(utf8Reader)
}

// 设置代理
func proxyHttpClient(proxyUrl string) *http.Client {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return nil
	}

	netTransport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Transport: netTransport,
	}
}

// 封装 golang.org/x/net/html 包，进行编码猜测
func determineEncoding(r *bufio.Reader) (encoding.Encoding, error) {
	// 通过 *Reader 读取 1024 byte 得到 []byte 类型的参数
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8, err
	}
	// 猜测编码
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e, nil
}
