package lowlevel

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/mittwald/spacectl/service/auth"
	"bytes"
	"time"
	"log"
)

type SpacesLowlevelClient struct {
	token string
	endpoint string
	version string

	client *http.Client
	logger *log.Logger
}

func NewSpacesLowlevelClient(token string, endpoint string, logger *log.Logger) (*SpacesLowlevelClient, error) {
	client := &http.Client{
	}

	return &SpacesLowlevelClient{
		token,
		endpoint,
		"v1",
		client,
		logger,
	}, nil
}

func (c *SpacesLowlevelClient) Get(path string, target interface{}) error {
	url := c.endpoint + "/" + c.version + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Access-Token", c.token)

	c.logger.Printf("executing GET on %s", url)

	client := http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 403 {
		return auth.InvalidCredentialsErr{}
	}

	if res.StatusCode >= 400 {
		msg := Message{}

		// The error here can safely be ignored since it does not matter much, anyway.
		// Either the response body contains a "msg" or it doesn't.
		_ = json.NewDecoder(res.Body).Decode(&msg)

		return ErrUnexpectedStatusCode{res.StatusCode, msg.String()}
	}

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("could not JSON-decode response body: %s", err)
	}

	return nil
}

func (c *SpacesLowlevelClient) Post(path string, body interface{}, target interface{}) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	url := c.endpoint + "/" + c.version + path
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Access-Token", c.token)

	c.logger.Printf("executing POST on %s: %s", url, string(reqBody))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	c.logger.Printf("response code: %d", res.StatusCode)

	if res.StatusCode == 403 {
		return auth.InvalidCredentialsErr{}
	}

	if res.StatusCode >= 400 {
		msg := Message{}

		// The error here can safely be ignored since it does not matter much, anyway.
		// Either the response body contains a "msg" or it doesn't.
		_ = json.NewDecoder(res.Body).Decode(&msg)

		return ErrUnexpectedStatusCode{res.StatusCode, msg.String()}
	}

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("could not JSON-decode response body: %s", err)
	}

	c.logger.Printf("response: %s", target)

	return nil
}