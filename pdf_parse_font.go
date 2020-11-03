package pdft

import (
	"io"
)

//PDFParseFont  parse font
func PDFParseFont(path string, name string) (*PDFFontData, error) {

	var fontData PDFFontData
	fontData.init()
	err := pdfParseFont(path, name, &fontData)
	if err != nil {
		return nil, err
	}
	return &fontData, nil
}

func pdfParseFont(path string, name string, outFontData *PDFFontData) error {
	err := outFontData.setTTFPath(path)
	if err != nil {
		return err
	}
	outFontData.setFontName(name)
	return nil
}

//PDFParseFont  parse font
func PDFParseFontByReader(r io.Reader, name string) (*PDFFontData, error) {

	var fontData PDFFontData
	fontData.init()
	err := pdfParseFontByReader(r, name, &fontData)
	if err != nil {
		return nil, err
	}
	return &fontData, nil
}

func pdfParseFontByReader(r io.Reader, name string, outFontData *PDFFontData) error {
	err := outFontData.setTTFByReader(r)
	if err != nil {
		return err
	}
	outFontData.setFontName(name)
	return nil
}
