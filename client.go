package balabola

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BalabolaRequest struct {
	Query  string `json:"query"`
	Intro  int    `json:"intro"`
	Filter int    `json:"filter"`
}

type BalabolaResponse struct {
	BadQuery int    `json:"bad_query"`
	Error    int    `json:"error"`
	Query    string `json:"query"`
	Text     string `json:"text"`
}

var DefaultClient = &Client{
	ctx: context.TODO(),
	c:   http.DefaultClient,
}

type Client struct {
	ctx context.Context

	c *http.Client
}

func (c *Client) TalkNonsense(query string, style TextStyle, filter int) (Nonsense string, err error) {
	reqBody := &BalabolaRequest{
		Intro:  int(style),
		Filter: filter,
		Query:  query,
	}

	// insert user uuid
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// create new req
	req, err := http.NewRequest(http.MethodPost, "https://yandex.ru/lab/api/yalm/text3", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return "", err
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "https://yandex.ru")
	req.Header.Set("Host", "yandex.ru")
	req.Header.Set("Accept-Language", "ru")

	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resp := &BalabolaResponse{}
	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(bodyBytes, resp)
	if err != nil {
		return "", err
	}

	if resp.BadQuery == 1 {
		return "", ErrSensitiveTopics
	}

	return resp.Text, nil
}

func New(ctx context.Context) *Client {
	c := &Client{
		ctx: ctx,
	}

	return c
}

func (c *Client) WithTransport(client *http.Client) {
	c.c = client
}

func (c *Client) WithContext(ctx context.Context) {
	c.ctx = ctx
}
