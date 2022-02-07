package szamlazzhu

import "encoding/base64"

func (x *Xmlszamlavalasz) GetPDF() ([]byte, error) {
	pdf := make([]byte, base64.StdEncoding.DecodedLen(len(x.Pdf)))
	_, err := base64.StdEncoding.Decode(pdf, x.Pdf)
	return pdf, err
}

func (x *Szamla) GetPDF() ([]byte, error) {
	pdf := make([]byte, base64.StdEncoding.DecodedLen(len(x.Pdf)))
	_, err := base64.StdEncoding.Decode(pdf, x.Pdf)
	return pdf, err
}

func (x *Xmlnyugtavalasz) GetPDF() ([]byte, error) {
	pdf := make([]byte, base64.StdEncoding.DecodedLen(len(x.NyugtaPdf)))
	_, err := base64.StdEncoding.Decode(pdf, x.NyugtaPdf)
	return pdf, err
}
