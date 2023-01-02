package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"sort"
	"strings"
)

// blockZeroFormatter is the default logger format.
type blockZeroFormatter struct {
	prefixed.TextFormatter
}

// imported from the formatter.
func (f *blockZeroFormatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.') {
			return true
		}
	}
	return false
}

// imported from the formatter.
func (f *blockZeroFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	switch value := value.(type) {
	case string:
		if !f.needsQuoting(value) {
			b.WriteString(value)
		} else {
			fmt.Fprintf(b, "%s%v%s", f.QuoteCharacter, value, f.QuoteCharacter)
		}
	case error:
		errmsg := value.Error()
		if !f.needsQuoting(errmsg) {
			b.WriteString(errmsg)
		} else {
			fmt.Fprintf(b, "%s%v%s", f.QuoteCharacter, errmsg, f.QuoteCharacter)
		}
	default:
		fmt.Fprint(b, value)
	}
}

func (f *blockZeroFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Generate the message buffer.
	buf := &bytes.Buffer{}
	// Get the keys.
	var keys []string = make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}
	lastKeyIdx := len(keys) - 1
	sort.Strings(keys)

	// Set message.
	buf.WriteString(fmt.Sprintf(
		"[%s] [%s] %s",
		entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message,
	))

	if len(keys) > 0 {
		buf.WriteByte(' ')
		buf.WriteByte('|')
		buf.WriteByte(' ')
	}

	// Add keys to string.
	for i, key := range keys {
		buf.WriteString(key)
		buf.WriteByte('=')
		f.appendValue(buf, entry.Data[key])

		if lastKeyIdx != i {
			buf.WriteByte(' ')
		}
	}

	// Write the EOS byte and return.
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}

// Log is the custom logger.
var Log = &logrus.Logger{
	Out:   os.Stdout,
	Level: logrus.InfoLevel,
	Formatter: &blockZeroFormatter{
		prefixed.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	},
}
