package imgflipgo

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
	X uint `json:"x,omitempty"`

	// [optional] Y coord of of the top left corner of the TextBox
	// If specified, must also specify X, Width, Height
	Y uint `json:"y,omitempty"`

	// [optional] width of the TextBox
	// If specified, must also specify X, Y, Height
	Width uint `json:"width,omitempty"`

	// [optional] height of the TextBox
	// If specified, must also specify X, Y, Width
	Height uint `json:"height,omitempty"`

	// [optional] Hex color for Text
	Color uint `json:"color,omitempty"`

	// [optional] Hex color for Text outline
	OutlineColor uint `json:"outline_color,omitempty"`
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
	TemplateID string `json:"template_id,omitempty"`

	// Username of a valid imgflip account. This is used to track where API
	// requests are coming from.
	Username string `json:"username,omitempty"`

	// Password for the imgflip account.
	Password string `json:"password,omitempty"`

	// Top text for the meme. Do not use this parameter if you are using the
	// boxes parameter below.
	TopText string `json:"text0,omitempty"`

	// Bottom text for the meme. Do not use this parameter if you are using the
	// boxes parameter below.
	BottomText string `json:"text1,omitempty"`

	// [optional] The font family to use for the text
	Font Font `json:"font,omitempty"`

	// [optional] Maximum font size in pixels. Defaults to 50px.
	MaxFontSizePx uint `json:"max_font_size,omitempty"`

	// [optional] For creating memes with more than two text boxes, or for further
	// customization. If TextBoxes is specified, TopText and BototmText will be ignored,
	// and text will not be automatically converted to uppercase, so you'll have to
	// handle capitalization yourself if you want the standard uppercase meme text.
	// The API is currently limited to 20 text boxes per image. The first TextBox in
	// the list may be left empty so that the second box will automatically be used
	// as bottom text.
	TextBoxes []TextBox `json:"boxes,omitempty"`
}

type CaptionResponse struct {
	Success bool `json:"success,omitempty"`
	Data    struct {
		URL     string
		PageURL string
	} `json:"data,omitempty"`

	ErrorMsg string `json:"error_message,omitempty"`
}
