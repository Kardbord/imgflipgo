package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/TannerKvarfordt/imgflipgo"
	"github.com/joho/godotenv"
)

const (
	ImgflipAPIUserEnv string = "IMGFLIP_API_USERNAME"
	ImgflipAPIPassEnv string = "IMGFLIP_API_PASSWORD"
)

var (
	ImgflipAPIUser string
	ImgflipAPIPass string
)

func init() {
	rand.Seed(time.Now().UnixNano())

	godotenv.Load()
	var found bool

	ImgflipAPIUser, found = os.LookupEnv(ImgflipAPIUserEnv)
	if !found || ImgflipAPIUser == "" {
		fmt.Printf("environment variable %s must be set\n", ImgflipAPIUserEnv)
		os.Exit(1)
	}

	ImgflipAPIPass, found = os.LookupEnv(ImgflipAPIPassEnv)
	if !found || ImgflipAPIPass == "" {
		fmt.Printf("environment variable %s must be set\n", ImgflipAPIPassEnv)
		os.Exit(1)
	}
}

func main() {
	template, err := randomTemplate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Caption the template \"%s\"\n", template.Name)
	captions := inputCaptions(*template)

	response, err := imgflipgo.CaptionImage(&imgflipgo.CaptionRequest{
		TemplateID: template.ID,
		Username:   ImgflipAPIUser,
		Password:   ImgflipAPIPass,
		TextBoxes:  captions,
	})
	if err != nil {
		fmt.Printf("Caption request failed, err=%v\n", err)
	}

	fmt.Printf("View your captioned image at %s\n", response.Data.URL)
}

func randomTemplate() (*imgflipgo.Meme, error) {
	memes, err := imgflipgo.GetMemes()
	if err != nil {
		return nil, err
	}

	return &memes[rand.Intn(len(memes))], nil
}

func inputCaptions(template imgflipgo.Meme) []imgflipgo.TextBox {
	captions := make([]imgflipgo.TextBox, template.BoxCount)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < int(template.BoxCount); i++ {
		fmt.Printf("Enter a caption for box #%d: ", i+1)
		input, _ := reader.ReadString('\n')
		captions[i].Text = input
		captions[i].Color = 0xFFA500 // orange
		captions[i].OutlineColor = 0xFF0000 // red
	}
	return captions
}
