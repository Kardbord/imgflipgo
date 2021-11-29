package imgflipgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/gorilla/schema"
)

const CaptionMemeEndpoint = "https://api.imgflip.com/caption_image"

type Font string

// TextBox specifies parameters for a CaptionReqeust TextBox.
// any [optional] parameters that are not set will default
// to settings provided by imgflip.com/memegenerator.
type TextBox struct {
	// Text to be displayed
	Text string `json:"text,omitempty"`

	// [optional] X coord of of the top left corner of the TextBox
	// If specified, must also specify Y, Width, Height
	X *uint `json:"x,omitempty"`

	// [optional] Y coord of of the top left corner of the TextBox
	// If specified, must also specify X, Width, Height
	Y *uint `json:"y,omitempty"`

	// [optional] width of the TextBox
	// If specified, must also specify X, Y, Height
	Width *uint `json:"width,omitempty"`

	// [optional] height of the TextBox
	// If specified, must also specify X, Y, Width
	Height *uint `json:"height,omitempty"`

	// [optional] Hex color for Text
	Color *uint `json:"color,omitempty"`

	// [optional] Hex color for Text outline
	OutlineColor *uint `json:"outline_color,omitempty"`
}

func (t *TextBox) SetX(x uint) *TextBox {
	t.X = &x
	return t
}
func (t *TextBox) SetY(y uint) *TextBox {
	t.Y = &y
	return t
}
func (t *TextBox) SetWidth(width uint) *TextBox {
	t.Width = &width
	return t
}
func (t *TextBox) SetHeight(height uint) *TextBox {
	t.Height = &height
	return t
}
func (t *TextBox) SetColor(color uint) *TextBox {
	t.Color = &color
	return t
}
func (t *TextBox) SetOutlineColor(outlineColor uint) *TextBox {
	t.OutlineColor = &outlineColor
	return t
}
func (t *TextBox) TextJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(t), "Text")
}
func (t *TextBox) XJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(t), "X")
}
func (t *TextBox) YJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(t), "Y")
}
func (t *TextBox) WidthJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(t), "Width")
}
func (t *TextBox) HeightJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(t), "Height")
}
func (t *TextBox) ColorJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(t), "Color")
}
func (t *TextBox) OutlineColorTextJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(t), "OutlineColor")
}

const (
	FontArial  Font = "arial"
	FontImpact Font = "impact"
)

const DefaultMaxFontSizePx uint = 50

type CaptionRequest struct {
	// A template ID as returned by the get_memes response. Any ID that was ever
	// returned from the get_memes response should work for this parameter. For
	// custom template uploads, the template ID can be found in the memegenerator
	// URL, e.g. https://imgflip.com/memegenerator/14859329/Charlie-Sheen-DERP.
	TemplateID string `schema:"template_id,omitempty" json:"template_id,omitempty"`

	// Username of a valid imgflip account. This is used to track where API
	// requests are coming from.
	Username string `schema:"username,omitempty" json:"username,omitempty"`

	// Password for the imgflip account.
	Password string `schema:"password,omitempty" json:"password,omitempty"`

	// Top text for the meme. Do not use this parameter if you are using the
	// boxes parameter below.
	TopText *string `schema:"text0,omitempty" json:"text0,omitempty"`

	// Bottom text for the meme. Do not use this parameter if you are using the
	// boxes parameter below.
	BottomText *string `schema:"text1,omitempty" json:"text1,omitempty"`

	// [optional] The font family to use for the text
	Font *Font `schema:"font,omitempty" json:"font,omitempty"`

	// [optional] Maximum font size in pixels. Defaults to 50px.
	MaxFontSizePx *uint `schema:"max_font_size,omitempty" json:"max_font_size,omitempty"`

	// [optional] For creating memes with more than two text boxes, or for further
	// customization. If TextBoxes is specified, TopText and BototmText will be ignored,
	// and text will not be automatically converted to uppercase, so you'll have to
	// handle capitalization yourself if you want the standard uppercase meme text.
	// The API is currently limited to 20 text boxes per image. The first TextBox in
	// the list may be left empty so that the second box will automatically be used
	// as bottom text.
	TextBoxes []TextBox `schema:"-" json:"boxes,omitempty"`
}

func (cr *CaptionRequest) SetTopText(topText string) *CaptionRequest {
	cr.TopText = &topText
	return cr
}
func (cr *CaptionRequest) SetBottomText(bottomText string) *CaptionRequest {
	cr.BottomText = &bottomText
	return cr
}
func (cr *CaptionRequest) SetFont(font Font) *CaptionRequest {
	cr.Font = &font
	return cr
}
func (cr *CaptionRequest) SetMaxFontSize(maxFontSizePx uint) *CaptionRequest {
	cr.MaxFontSizePx = &maxFontSizePx
	return cr
}
func (cr *CaptionRequest) TemplateIDJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "TemplateID")
}
func (cr *CaptionRequest) UsernameJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "Username")
}
func (cr *CaptionRequest) PasswordJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "Password")
}
func (cr *CaptionRequest) TopTextJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "TopText")
}
func (cr *CaptionRequest) BottomTextJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "BottomText")
}
func (cr *CaptionRequest) FontJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "Font")
}
func (cr *CaptionRequest) MaxFontSizePxJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "MaxFontSizePx")
}
func (cr *CaptionRequest) TextBoxesJSONTag() (string, error) {
	return getStructFieldJSONTag(reflect.TypeOf(cr), "TextBoxes")
}

var encoder schema.Encoder

func init() {
	encoder = *schema.NewEncoder()
	encoder.RegisterEncoder(CaptionRequest{}.Font, func(value reflect.Value) string {
		return value.String()
	})
}

func (cr CaptionRequest) CreateHTTPFormBody() (url.Values, error) {
	form := url.Values{}
	err := encoder.Encode(cr, form)
	if err != nil {
		return form, err
	}

	// Have to encode TextBoxes manually because gorilla/schema doesn't handle
	// slices of custom structs (or if it does I'm too dumb to figure out how).
	textBoxesJSONTag, err := cr.TextBoxesJSONTag()
	if err != nil {
		return form, err
	}

	var textJSONTag string
	var xJSONTag string
	var yJSONTag string
	var widthJSONTag string
	var heightJSONTag string
	var colorJSONTag string
	var outlineColorJSONTag string
	for i := range cr.TextBoxes {
		if i == 0 {
			textJSONTag, err = cr.TextBoxes[i].TextJSONTag()
			if err != nil {
				return form, err
			}
			xJSONTag, err = cr.TextBoxes[i].XJSONTag()
			if err != nil {
				return form, err
			}
			yJSONTag, err = cr.TextBoxes[i].YJSONTag()
			if err != nil {
				return form, err
			}
			widthJSONTag, err = cr.TextBoxes[i].WidthJSONTag()
			if err != nil {
				return form, err
			}
			heightJSONTag, err = cr.TextBoxes[i].HeightJSONTag()
			if err != nil {
				return form, err
			}
			colorJSONTag, err = cr.TextBoxes[i].ColorJSONTag()
			if err != nil {
				return form, err
			}
			outlineColorJSONTag, err = cr.TextBoxes[i].OutlineColorTextJSONTag()
			if err != nil {
				return form, err
			}
		}

		form.Add(fmt.Sprintf("%s[%d][%s]", textBoxesJSONTag, i, textJSONTag), cr.TextBoxes[i].Text)

		if cr.TextBoxes[i].X != nil {
			form.Add(fmt.Sprintf("%s[%d][%s]", textBoxesJSONTag, i, xJSONTag), fmt.Sprint(*cr.TextBoxes[i].X))
		}

		if cr.TextBoxes[i].Y != nil {
			form.Add(fmt.Sprintf("%s[%d][%s]", textBoxesJSONTag, i, yJSONTag), fmt.Sprint(*cr.TextBoxes[i].Y))
		}

		if cr.TextBoxes[i].Width != nil {
			form.Add(fmt.Sprintf("%s[%d][%s]", textBoxesJSONTag, i, widthJSONTag), fmt.Sprint(*cr.TextBoxes[i].Width))
		}

		if cr.TextBoxes[i].Height != nil {
			form.Add(fmt.Sprintf("%s[%d][%s]", textBoxesJSONTag, i, heightJSONTag), fmt.Sprint(*cr.TextBoxes[i].Height))
		}

		if cr.TextBoxes[i].Color != nil {
			form.Add(fmt.Sprintf("%s[%d][%s]", textBoxesJSONTag, i, colorJSONTag), fmt.Sprintf("#%06x", *cr.TextBoxes[i].Color))
		}

		if cr.TextBoxes[i].OutlineColor != nil {
			form.Add(fmt.Sprintf("%s[%d][%s]", textBoxesJSONTag, i, outlineColorJSONTag), fmt.Sprintf("#%06x", *cr.TextBoxes[i].OutlineColor))
		}
	}

	return form, nil
}

type CaptionResponse struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		URL     string
		PageURL string
	} `json:"data,omitempty"`

	ErrorMsg string `json:"error_message,omitempty"`
}

// CaptionImage wraps the caption_image endpoint. It makes a request using the provided
// parameters, and will return an error if something goes wrong at any point during the
// process. The errors can come originate in Go or be from the API, depending on where
// the failure occurred. This was done so that the caller does not have to check both
// the returned error value, AND CaptionResponse.Success. If the API returns an error,
// it will be reflected in both CaptionResponse.ErrorMsg and in the returned Go error.
func CaptionImage(req *CaptionRequest) (CaptionResponse, error) {
	if req == nil {
		return CaptionResponse{Success: false, ErrorMsg: "nil request provided"}, errors.New("nil request provided")
	}

	form, err := req.CreateHTTPFormBody()
	if err != nil {
		return CaptionResponse{Success: false, ErrorMsg: fmt.Sprint(err)}, err
	}

	resp, err := http.PostForm(CaptionMemeEndpoint, form)
	if err != nil {
		return CaptionResponse{Success: false, ErrorMsg: fmt.Sprint(err)}, err
	}
	if resp == nil {
		return CaptionResponse{Success: false, ErrorMsg: "nil response received"}, errors.New("nil response received")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CaptionResponse{Success: false, ErrorMsg: fmt.Sprint(err)}, err
	}

	captionResponse := CaptionResponse{}
	err = json.Unmarshal(respBody, &captionResponse)
	if err != nil {
		return CaptionResponse{Success: false, ErrorMsg: fmt.Sprint(err)}, err
	}

	if !captionResponse.Success {
		return captionResponse, errors.New(captionResponse.ErrorMsg)
	}

	return captionResponse, nil
}
