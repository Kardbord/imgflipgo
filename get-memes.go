package imgflipgo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const GetMemesEndpoint string = "https://api.imgflip.com/get_memes"

type Meme struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	URL    string `json:"url,omitempty"`
	Width  uint   `json:"width,omitempty"`
	Height uint   `json:"height,omitempty"`

	// BoxCount describes the number of text boxes the Meme uses.
	BoxCount uint `json:"box_count,omitempty"`
}

type MemesResponse struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		Memes []Meme `json:"memes,omitempty"`
	} `json:"data,omitempty"`
}

func GetMemesWithResponse() (*MemesResponse, error) {
	resp, err := http.Get(GetMemesEndpoint)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("nil response received")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	memesResp := MemesResponse{}
	err = json.Unmarshal(body, &memesResp)
	if err != nil {
		return nil, err
	}

	return &memesResp, err
}

func GetMemes() ([]Meme, error) {
	memesResp, err := GetMemesWithResponse()
	if err != nil {
		return nil, err
	}
	if memesResp == nil {
		return nil, errors.New("nil response received")
	}
	if !memesResp.Success {
		return nil, errors.New("request was unsuccessful")
	}
	return memesResp.Data.Memes, nil
}
