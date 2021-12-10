package clients

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (v *vectorizer) MetaInfo() (map[string]interface{}, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", v.url("/meta"), nil)
	if err != nil {
		return nil, errors.Wrap(err, "create GET meta request")
	}

	res, err := v.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send GET meta request")
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read meta response body")
	}

	var resBody map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &resBody); err != nil {
		return nil, errors.Wrap(err, "unmarshal meta response body")
	}
	return resBody, nil
}