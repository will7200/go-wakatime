package internal

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-openapi/strfmt"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

func init() {
	d := Date{}
	// register this format in the default registry
	strfmt.Default.Add("date(us)", &d, IsDate)
}

// IsDate returns true when the string is a valid date
func IsDate(str string) bool {
	_, err := time.Parse(DateUSRepresentation, str)
	return err == nil
}

const (
	// DateUSRepresentation represents a full-date as specified by RFC3339
	// See: http://goo.gl/xXOvVd
	DateUSRepresentation = "01/02/2006"
	jsonNull             = "null"
)

// Date represents a date from the API
//
// swagger:type date(us)
type Date time.Time

// String converts this date into a string
func (d Date) String() string {
	return time.Time(d).Format(DateUSRepresentation)
}

// UnmarshalText parses a text representation into a date type
func (d *Date) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}
	dd, err := time.Parse(DateUSRepresentation, string(text))
	if err != nil {
		return err
	}
	*d = Date(dd)
	return nil
}

// MarshalText serializes this date type to string
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// Scan scans a Date value from database driver type.
func (d *Date) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case []byte:
		return d.UnmarshalText(v)
	case string:
		return d.UnmarshalText([]byte(v))
	case time.Time:
		*d = Date(v)
		return nil
	case nil:
		*d = Date{}
		return nil
	default:
		return fmt.Errorf("cannot sql.Scan() strfmt.Date from: %#v", v)
	}
}

// Value converts Date to a primitive value ready to written to a database.
func (d Date) Value() (driver.Value, error) {
	return driver.Value(d.String()), nil
}

// MarshalJSON returns the Date as JSON
func (d Date) MarshalJSON() ([]byte, error) {
	var w jwriter.Writer
	d.MarshalEasyJSON(&w)
	return w.BuildBytes()
}

// MarshalEasyJSON writes the Date to a easyjson.Writer
func (d Date) MarshalEasyJSON(w *jwriter.Writer) {
	w.String(time.Time(d).Format(DateUSRepresentation))
}

// UnmarshalJSON sets the Date from JSON
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}
	l := jlexer.Lexer{Data: data}
	d.UnmarshalEasyJSON(&l)
	return l.Error()
}

// UnmarshalEasyJSON sets the Date from a easyjson.Lexer
func (d *Date) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if data := in.String(); in.Ok() {
		tt, err := time.Parse(DateUSRepresentation, data)
		if err != nil {
			in.AddError(err)
			return
		}
		*d = Date(tt)
	}
}

// GetBSON returns the Date as a bson.M{} map.
func (d *Date) GetBSON() (interface{}, error) {
	return bson.M{"data": d.String()}, nil
}

// SetBSON sets the Date from raw bson data
func (d *Date) SetBSON(raw bson.Raw) error {
	var m bson.M
	if err := raw.Unmarshal(&m); err != nil {
		return err
	}

	if data, ok := m["data"].(string); ok {
		rd, err := time.Parse(DateUSRepresentation, data)
		*d = Date(rd)
		return err
	}

	return errors.New("couldn't unmarshal bson raw value as Date")
}

// DeepCopyInto copies the reciever and writes its value into out.
func (in *Date) DeepCopyInto(out *Date) {
	*out = *in
	return
}

// DeepCopy copies the receiver into a new Date.
func (in *Date) DeepCopy() *Date {
	if in == nil {
		return nil
	}
	out := new(Date)
	in.DeepCopyInto(out)
	return out
}
