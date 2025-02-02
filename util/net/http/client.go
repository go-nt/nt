package http

import (
	"bytes"
	"encoding/json"
	"io"
	netHttp "net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func NewClient() (*Client, error) {
	client := Client{}

	return &client, nil
}

type Client struct {
	httpClient  *netHttp.Client
	rootUrl     string
	accessToken string
}

func (client *Client) init(rootUrl string, accessToken string) {
	client.httpClient = &netHttp.Client{
		Timeout: 3 * time.Second,
	}
	client.rootUrl = rootUrl
	client.accessToken = accessToken
}

func (client *Client) Get(urlPath string) ([]byte, error) {
	return client.Request(urlPath, "get", nil)
}

func (client *Client) Post(urlPath string, data map[string]any) ([]byte, error) {
	return client.Request(urlPath, "post", data)
}

func (client *Client) PostJson(urlPath string, data map[string]any) ([]byte, error) {
	return client.Request(urlPath, "post-json", data)
}

func (client *Client) GetAndBind(urlPath string, v any) error {
	bytes, err := client.Request(urlPath, "get", nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) PostAndBind(urlPath string, data map[string]any, v any) error {
	bytes, err := client.Request(urlPath, "post", data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) PostJsonAndBind(urlPath string, data map[string]any, v any) error {
	bytes, err := client.Request(urlPath, "post-json", data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Request(urlPath string, method string, data map[string]any) ([]byte, error) {
	var request *netHttp.Request
	var err error

	switch method {
	case "get":
		request, err = netHttp.NewRequest(netHttp.MethodGet, client.rootUrl+urlPath, nil)
		if err != nil {
			return nil, err
		}

	case "post":

		// form 表单 application/x-www-form-urlencoded
		values := url.Values{}
		for key, value := range data {
			switch t := value.(type) {
			case string:
				values.Set(key, t)
			case bool:
				if t {
					values.Set(key, "true")
				} else {
					values.Set(key, "false")
				}
			case int:
				values.Set(key, strconv.Itoa(t))
			}
		}

		dataReader := strings.NewReader(values.Encode())

		request, err = netHttp.NewRequest(netHttp.MethodPost, client.rootUrl+urlPath, dataReader)
		if err != nil {
			return nil, err
		}

	case "post-json":

		dataByte, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		dataReader := bytes.NewBuffer(dataByte)

		request, err = netHttp.NewRequest(netHttp.MethodPost, client.rootUrl+urlPath, dataReader)
		if err != nil {
			return nil, err
		}

		request.Header.Set("Content-Type", "application/json")
	}

	request.Header.Set("PRIVATE-TOKEN", client.accessToken)

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func download(url string, path string) error {

	resp, err := netHttp.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建目录
	dir := filepath.Dir(path)
	_, err = os.ReadDir(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 文件已存在时删除
	_, err = os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			os.Remove(path)
		}
	}

	// 创建文件
	outFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 写入数据
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
