package extractPages

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"image/jpeg"
	"os"
	"path/filepath"
)

func ExtractImages(pdf string) error {

	if err := removeContents("./staticPages/images/"); err != nil {
		return err
	}

	doc, err := fitz.New(pdf)
	if err != nil {
		return err
	}

	defer doc.Close()

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			return err
		}

		f, err := os.Create(filepath.Join("./staticPages/images", fmt.Sprintf("img%03d.jpg", n)))
		if err != nil {
			return err
		}

		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			return  err
		}

		f.Close()
	}
	return nil
}

// Remove contents of image folder
func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
