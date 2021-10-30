package imgflipgo_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/TannerKvarfordt/imgflipgo"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

const (
	ImgflipAPIUsernameEnv = "IMGFLIP_API_USERNAME"
	ImgflipAPIPasswordEnv = "IMGFLIP_API_PASSWORD"
)

var (
	ImgflipAPIUsername string
	ImgflipAPIPassword string
)

func setup() {
	godotenv.Load()
	ImgflipAPIUsername, _ = os.LookupEnv(ImgflipAPIUsernameEnv)
	ImgflipAPIPassword, _ = os.LookupEnv(ImgflipAPIPasswordEnv)
}

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
		t.Fatal("Expected request to fail")
	}
}

func expectSuccess(t *testing.T, resp *imgflipgo.CaptionResponse, err error) {
	expectNilError(t, err)
	expectNonNilResponse(t, resp)
	if !resp.Success {
		t.Fatalf("Expected request to succeed, err=%s", resp.ErrorMsg)
	}
}

const testTemplateID = "181913649"

func TestCaptionImageNilRequest(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(nil)
	expectNonNilError(t, err)
	expectNilResponse(t, resp)
}

func TestCaptionImageEmptyRequest(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{})
	expectUnsuccessfulNoErr(t, resp, err)
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
		TopText:    "FOO",
		BottomText: "BAR",
	})
	expectUnsuccessfulNoErr(t, resp, err)
}

func TestCaptionImageInvalidAuth(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   "not_a_real_user_1234u173829",
		Password:   "asdf",
		TopText:    "TOP TEXT",
		BottomText: "BOTTOM TEXT",
	})
	expectUnsuccessfulNoErr(t, resp, err)
}

func TestCaptionImageTopTextOnly(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TopText:    "Top Text",
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageBottomTextOnly(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		BottomText: "Bottom Text",
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageTopAndBottomText(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TopText:    "Top Text",
		BottomText: "Bottom Text",
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageArial(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TopText:    "Top Text",
		BottomText: "Bottom Text",
		Font:       imgflipgo.FontArial,
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageImpact(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TopText:    "Top Text",
		BottomText: "Bottom Text",
		Font:       imgflipgo.FontImpact,
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageMaxFontSize(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID:    testTemplateID,
		Username:      ImgflipAPIUsername,
		Password:      ImgflipAPIPassword,
		TopText:       "Top Text",
		BottomText:    "Bottom Text",
		MaxFontSizePx: uint(rand.Intn(int(imgflipgo.DefaultMaxFontSizePx))),
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageTopTextbox(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{
				Text: "Top Text",
			},
		},
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageBottomTextbox(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{},
			{
				Text: "Bottom Text",
			},
		},
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}
