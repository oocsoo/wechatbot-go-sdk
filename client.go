package wechatbot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	http    *http.Client

	Auth     *AuthModule
	Message  *MessageModule
	Group    *GroupModule
	Contact  *ContactModule
	Personal *PersonalModule
	Favorite *FavoriteModule
	Label    *LabelModule
	Download *DownModule
	Sns      *SnsModule
	Finder   *FinderModule
}

type Option func(*Client)

func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = url }
}

func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.http.Timeout = d }
}

func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL: "http://124.221.45.58",
		http:    &http.Client{Timeout: 30 * time.Second},
	}
	for _, o := range opts {
		o(c)
	}
	c.Auth = &AuthModule{client: c}
	c.Message = &MessageModule{client: c}
	c.Group = &GroupModule{client: c}
	c.Contact = &ContactModule{client: c}
	c.Personal = &PersonalModule{client: c}
	c.Favorite = &FavoriteModule{client: c}
	c.Label = &LabelModule{client: c}
	c.Download = &DownModule{client: c}
	c.Sns = &SnsModule{client: c}
	c.Finder = &FinderModule{client: c}
	return c
}

func (c *Client) Request(endpoint, token string, payload any) (*Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.baseURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("AUTHORIZATION", token)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r Response
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
