package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"ptapp.cn/util/dlog.v1"
	"ptapp.cn/util/filter/mtype"
)

type Api struct {
	headers map[string]string
}

const (
	ContentType       string = "Content-Type"
	ApplicationJson   string = "application/json"
	WwwFormUrlencoded string = "application/x-www-form-urlencoded"
)

func (a *Api) Header(key string, val string) *Api {
	if a.headers == nil {
		a.headers = map[string]string{}
	}
	a.headers[key] = val
	return a
}

func (a *Api) getBytes(ctx context.Context, fields interface{}) ([]byte, error) {
	var fieldsByte []byte
	var err error
	switch a.headers[ContentType] {
	case ApplicationJson:
		if fieldString, ok := fields.(string); ok {
			if fieldString != "" {
				mapp := make(map[string]interface{})
				err = json.Unmarshal([]byte(fieldString), &mapp)
				if err != nil {
					return nil, err
				}
				fieldsByte, err = json.Marshal(mapp)
				if err != nil {
					return nil, err
				}
			}
		} else {
			fieldsByte, err = json.Marshal(fields)
			if err != nil {
				return nil, err
			}
		}
	case WwwFormUrlencoded:
		if fieldString, ok := fields.(string); ok {
			return []byte(fieldString), nil
		}
		if fieldsByte, ok := fields.([]byte); ok {
			return fieldsByte, nil
		}
		fieldsByte, err = json.Marshal(fields)
		if err != nil {
			return nil, err
		}
		mapp := make(map[string]interface{})
		err = json.Unmarshal(fieldsByte, &mapp)
		if err != nil {
			return nil, err
		}
		params := url.Values{}
		for k, v := range mapp {
			params.Add(k, mtype.GetString(v))
		}
		return []byte(params.Encode()), nil
	}

	return fieldsByte, nil
}
func (a *Api) sendRequest(ctx context.Context, method string, url string, fields interface{}) ([]byte, error) {
	dl := dlog.FromContext(ctx)
	fieldsByte, err := a.getBytes(ctx, fields)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	request, err := a.createRequest(ctx, method, url, fieldsByte)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	} else {

		defer func() {
			resp.Body.Close()
		}()

		dl.V(5).Infof("resp %+v", resp)
		if resp.StatusCode != 200 {
			dl.Errorf("status not 200:%+v,%v", resp.StatusCode, resp.Status)
			bd, err := ioutil.ReadAll(resp.Body)
			dl.Errorf("status not 200:%+v,%v", string(bd), err)
			if err == nil {
				return nil, errors.New(string(bd))
			}
		}
		bd, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bd, nil
	}
}

func (a *Api) createRequest(ctx context.Context, method string, url string, fields []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(fields))
	if err != nil {
		return nil, err
	}
	for k, v := range a.headers {
		req.Header.Add(k, v)
	}
	a.log(ctx, method, url, fields)
	return req, nil
}
func (a *Api) log(ctx context.Context, method string, url string, fields []byte) {
	dl := dlog.FromContext(ctx)
	headers := ""
	for k, v := range a.headers {
		headers += fmt.Sprintf("-H \"%v:%v\"", k, v)
	}
	dl.V(5).Infof("curl %v -X %s %s -d '%v'", headers, method, url, string(fields))
	dl.V(5).Infof("curl %v", headers)
}
func (a *Api) Get(ctx context.Context, url string, fields interface{}) ([]byte, error) {
	return a.sendRequest(ctx, "GET", url, fields)
}

func (a *Api) Post(ctx context.Context, url string, fields interface{}) ([]byte, error) {
	return a.sendRequest(ctx, "POST", url, fields)
}
func (a *Api) ApplicationJson() *Api {
	return a.Header(ContentType, ApplicationJson)
}

func (a *Api) WwwFormUrlencoded() *Api {
	return a.Header(ContentType, WwwFormUrlencoded)
}
