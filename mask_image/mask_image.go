package mask_image

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/neko-neko/lgtmgen/images"
	"image"
	"io/ioutil"
)

type MaskImage struct {
	Width int
	Height int
	MaskImage image.Image
}

// constructor
func NewMaskImage() *MaskImage {
	return &MaskImage{
		Width: 0,
		Height: 0,
		MaskImage: nil,
	}
}

// Load mask image
func (m *MaskImage) LoadMaskImage(maskImage string) error {
	imageByte, err := images.Asset(maskImage)
	if err != nil {
		return err
	}

	// convert []byte to Image.image
	img, _, _ := image.Decode(bytes.NewReader(imageByte))

	// load mask image config
	size := img.Bounds().Size()

	// set self
	m.Height = size.Y
	m.Width = size.X
	m.MaskImage = img

	return nil
}

// Get target image paths from target dir
func (m *MaskImage) ReadImagePaths(target string) []string {
	files, err := ioutil.ReadDir(target)
	if err != nil {
		panic(err)
	}

	// create full path lists
	var filesPaths []string
	for _, fileInfo := range files {
		// skip directory
		if fileInfo.IsDir() {
			continue
		}

		filesPaths = append(filesPaths, target+fileInfo.Name())
	}

	return filesPaths
}

// Execute mask
func (m *MaskImage) OverlayImage(file string, maskImage image.Image, width int, height int) (*image.NRGBA, error) {
	srcImage, err := imaging.Open(file)
	if err != nil {
		return nil, err
	}

	resizedImage := imaging.Resize(srcImage, width, height, imaging.Box)
	maskedImage := imaging.OverlayCenter(resizedImage, maskImage, 1.0)
	return maskedImage, nil
}