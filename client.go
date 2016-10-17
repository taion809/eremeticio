package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

type HttpClient struct {
	url string
	*http.Client
}

func NewClient(url string) *HttpClient {
	return &HttpClient{
		url,
		&http.Client{
			Timeout: 500 * time.Millisecond,
		},
	}
}

func (h *HttpClient) FetchTasks(ctx context.Context) ([]*Task, error) {
	req, _ := http.NewRequest("GET", h.url+"/task", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := h.send(ctx, req)
	if err != nil {
		return nil, err
	}

	return deserializeTaskList(resp.Body)
}

func (h *HttpClient) send(ctx context.Context, req *http.Request) (*http.Response, error) {
	return ctxhttp.Do(ctx, h.Client, req)
}

func deserializeTaskList(body io.ReadCloser) ([]*Task, error) {
	var taskList []*Task

	err := json.NewDecoder(body).Decode(&taskList)
	if err != nil {
		return nil, err
	}

	body.Close()

	return taskList, nil
}
