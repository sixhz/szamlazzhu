package szamlazzhu

import (
	"encoding/xml"
	"time"
)

// Date implements date fields as go time.Time
type Date struct {
	time.Time
}

// UnmarshalText unmarshals from 2006-01-02 format
func (d *Date) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02", string(text))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalText marshals a Date from time.Time to 2006-01-02 format
func (d *Date) MarshalText() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte(""), nil
	}
	s := d.Time.Format("2006-01-02")
	return []byte(s), nil
}

// UnmarshalXML is a wrapper to override time.Time.UnmarshalXML with UnmarshalText
func (d *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	err := decoder.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	return d.UnmarshalText([]byte(s))
}

// MarshalXML is a wrapper to override time.Time.MarshalXML with MarshalText
func (d *Date) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	s, err := d.MarshalText()
	if err != nil {
		return err
	}
	if len(s) == 0 {
		return nil // ,omitempty-like behavior for dates.
	}
	return encoder.EncodeElement(&s, start)
}
