package imgflipgo_test

import (
	"testing"

	"github.com/TannerKvarfordt/imgflipgo"
)

const testTemplateID = "181913649"

func expectNilError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
}

func expectNonNilError(t *testing.T, err error) {
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
}

func expectNilResponse(t *testing.T, resp *imgflipgo.CaptionResponse) {
	if resp != nil {
		t.Fatalf("Expected nil response, got %v", resp)
	}
}

func expectNonNilResponse(t *testing.T, resp *imgflipgo.CaptionResponse) {
	if resp == nil {
		t.Fatal("Expected non-nil response")
	}
}

func expectUnsuccessfulNoErr(t *testing.T, resp *imgflipgo.CaptionResponse, err error) {
	expectNilError(t, err)
	expectNonNilResponse(t, resp)
	if resp.Success {
		t.Fatalf("Expected request to fail")
	}
}

func TestCaptionImageNilRequest(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(nil)
	expectNonNilError(t, err)
	expectNilResponse(t, resp)
}

func TestCaptionImageNoUser(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Password:   "asdf",
	})
	expectUnsuccessfulNoErr(t, resp, err)
}

func TestCaptionImageNoPW(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   "asdf",
	})
	expectUnsuccessfulNoErr(t, resp, err)
}

func TestCaptionImageNoUserOrPW(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
	})
	expectUnsuccessfulNoErr(t, resp, err)
}
