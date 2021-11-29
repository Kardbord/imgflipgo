package imgflipgo_test

import (
	"fmt"
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

const testTemplateID = "181913649"

func expectSuccess(t *testing.T, resp imgflipgo.CaptionResponse, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Success {
		t.Fatal("Expected successful query, err=", resp.ErrorMsg)
	}
	if resp.ErrorMsg != "" {
		t.Fatal("Did not expecte an error. ErrorMsg=", resp.ErrorMsg)
	}
	if resp.Data.URL == "" && resp.Data.PageURL == "" {
		t.Fatal("Response URLs should not both be empty")
	}
}

func expectFailure(t *testing.T, resp imgflipgo.CaptionResponse, err error) {
	if err == nil {
		t.Fatal("Expected an error")
	}
	if resp.Success {
		t.Fatal("Expected an error")
	}
	if resp.ErrorMsg == "" {
		t.Fatal("Expected a non-empty ErrorMsg")
	}
	if fmt.Sprint(err) == "" {
		t.Fatal("Expected a non-empty ErrorMsg")
	}
	if resp.ErrorMsg != fmt.Sprint(err) {
		t.Fatalf("returned error and ErrorMsg strings should be equivalent, but %s != %s", resp.ErrorMsg, err)
	}
	if resp.Data.PageURL != "" || resp.Data.URL != "" {
		t.Fatalf("Did not expect response URLs to be populated; PageURL=%s URL=%s", resp.Data.PageURL, resp.Data.URL)
	}
}

func TestCaptionImageNilRequest(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(nil)
	expectFailure(t, resp, err)
}

func TestCaptionImageEmptyRequest(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{})
	expectFailure(t, resp, err)
}

func TestCaptionImageNoUser(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Password:   "asdf",
	})
	expectFailure(t, resp, err)
}

func TestCaptionImageNoPW(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   "asdf",
	})
	expectFailure(t, resp, err)
}

func TestCaptionImageNoUserOrPW(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
	}).SetTopText("FOO").SetBottomText("BAR"))
	expectFailure(t, resp, err)
}

func TestCaptionImageInvalidAuth(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   "not_a_real_user_1234u173829",
		Password:   "asdf",
	}).SetTopText("TOP TEXT").SetBottomText("BOTTOM TEXT"))
	expectFailure(t, resp, err)
}

func TestCaptionImageTopTextOnly(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
	}).SetTopText("Top Text"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageBottomTextOnly(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
	}).SetBottomText("Bottom Text"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageTopAndBottomText(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
	}).SetTopText("Top Text").SetBottomText("Bottom Text"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageArial(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
	}).SetTopText("Top Text").SetBottomText("Bottom Text").SetFont(imgflipgo.FontArial))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageImpact(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
	}).SetTopText("Top Text").SetBottomText("Bottom Text").SetFont(imgflipgo.FontImpact))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageMaxFontSize(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
	}).SetTopText("Top Text").SetBottomText("Bottom Text").SetMaxFontSize(uint(rand.Intn(int(imgflipgo.DefaultMaxFontSizePx)))))
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

func TestCaptionImageTextBoxes(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{
				Text: "Top Text",
			},
			{
				Text: "Bottom Text",
			},
		},
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageTopBottomAndBoxes(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{
				Text: "Top Text",
			},
			{
				Text: "Bottom Text",
			},
		},
	}).SetTopText("This text should not be displayed (top)").SetBottomText("This text should not be displayed (bottom)"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageTopAndBox(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{
				Text: "Top Text",
			},
		},
	}).SetTopText("This text should not be displayed (top)"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageBottomAndBox(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{},
			{
				Text: "Bottom Text",
			},
		},
	}).SetBottomText("This text should not be displayed (bottom)"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageTopAndBottomBox(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{},
			{
				Text: "Bottom Text",
			},
		},
	}).SetTopText("This text should not be displayed (top)"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageBottomAndTopBox(t *testing.T) {
	resp, err := imgflipgo.CaptionImage((&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			{
				Text: "Top Text",
			},
		},
	}).SetBottomText("This text should not be displayed (bottom)"))
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}

func TestCaptionImageColor(t *testing.T) {
	resp, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: testTemplateID,
		Username:   ImgflipAPIUsername,
		Password:   ImgflipAPIPassword,
		TextBoxes: []imgflipgo.TextBox{
			*((&imgflipgo.TextBox{
				Text: "Orange text",
			}).SetColor(0xFFA500)),
			*((&imgflipgo.TextBox{
				Text: "Red Outline",
			}).SetOutlineColor(0xFF0000)),
		},
	})
	expectSuccess(t, resp, err)
	t.Logf("See test image at %s", resp.Data.URL)
}
