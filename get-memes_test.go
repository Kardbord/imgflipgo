package imgflipgo_test

import (
	"testing"

	"github.com/Kardbord/imgflipgo"
)

func TestGetMemesResponse(t *testing.T) {
	memesResp, err := imgflipgo.GetMemesWithResponse()
	if err != nil {
		t.Fatal(err)
	}
	if !memesResp.Success {
		t.Fatal("GET was not successful")
	}
	if len(memesResp.Data.Memes) < 1 {
		t.Fatal("No memes were retrieved despite successful request")
	}
	t.Logf("Retrieved %d memes", len(memesResp.Data.Memes))
}

func TestGetMemes(t *testing.T) {
	memes, err := imgflipgo.GetMemes()
	if err != nil {
		t.Fatal(err)
	}
	if len(memes) < 1 {
		t.Fatal("No memes were retrieved despite successful request")
	}
	t.Logf("Retrieved %d memes", len(memes))
}
