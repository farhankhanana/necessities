package utils

import (
	"bytes"
	crand "crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"text/template"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var lettersPassword = []rune("!@#$%^&*-_=+|,.?abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// WhiteSpaceTrimmer removes white spaces.
func WhiteSpaceTrimmer(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// TimestampNow returns current timestamp.
// If format is not specified, it will return timestamp in RFC3339 format.
func TimestampNow(UTC bool, format string) (time.Time, string) {
	var timestampNow time.Time
	if UTC {
		timestampNow = time.Now().UTC()
	} else {
		timestampNow = time.Now().Local()
	}

	if format != "" {
		return timestampNow, timestampNow.Format(format)
	}

	return timestampNow, timestampNow.Format(time.RFC3339)

}

// AnyToJsonStr converts any data type to JSON string if the data type itself supported to be converted to JSON.
func AnyToJsonStr(data interface{}) string {
	switch d := data.(type) {
	case []byte:
		return string(d)
	default:
		dataByte, err := json.Marshal(d)
		if err != nil {
			return fmt.Sprintf("%v", d)
		}
		return string(dataByte)
	}
}

// AnyToMapStringInterface converts any data type to map[string]interface{}.
func AnyToMapStringInterface(data interface{}) map[string]interface{} {
	functionName := "[AnyToMapStringInterface]"
	var result map[string]interface{}
	var err error
	caseName := ""
	switch d := data.(type) {
	case []byte:
		if len(d) > 0 {
			caseName = "[]byte"
			err = json.Unmarshal(d, &result)
		}
	default:
		caseName = "default"
		dataByte, _ := json.Marshal(data)
		err = json.Unmarshal(dataByte, &result)
	}
	if err != nil {
		log.Println(functionName, fmt.Sprintf("[case: %s]", caseName), "error unmarshal dataByte", err)
		return nil
	}

	return result
}

// PrintStructValue iterates struct.
func PrintStructValue(s interface{}) error {
	functionName := "[PrintStructValue]"
	var structValue map[string]interface{}
	sByte, _ := json.Marshal(s)
	err := json.Unmarshal(sByte, &structValue)
	if err != nil {
		log.Println(functionName, "error unmarshal sByte", err)
		return err
	}

	for k, v := range structValue {
		fmt.Printf("Field: %s\tValue: %v\n", k, v)
	}
	return err
}

// DecodeJWT decodes JWT.
func DecodeJWT(accessToken string) (jwt.MapClaims, error) {
	decoded, _ := jwt.Parse(accessToken, nil)

	claims, ok := decoded.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ParseTemplateToString parses a form template and returns a string.
func ParseTemplateToString(tmp string, data interface{}) (string, error) {
	var err error
	buf := new(bytes.Buffer)
	t, _ := template.New("template").Parse(tmp)

	if err = t.Execute(buf, data); err != nil {
		return buf.String(), err
	}

	return buf.String(), err
}

// ParseTemplateEmailSMS parses a form template `s` for email and sms.
func ParseTemplateEmailSMS(subject string, s string, data interface{}) *bytes.Buffer {
	functionName := "[ParseTemplateEmailSMS]"
	t, _ := template.New("template").Parse(s)
	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+"\n%s\n\n", headers)))

	err := t.Execute(&body, data)
	if err != nil {
		log.Println(functionName, "error executing", err)
		return nil
	}

	return &body
}

// GeneratePassword generates random password.
// If length not specified or zero, default length is 12.
func GeneratePassword(length int) (string, error) {
	functionName := "[GeneratePassword]"
	rand.Seed(time.Now().UnixNano())
	if length == 0 {
		length = 12
	}
	b := make([]rune, length)
	for i := range b {
		num, err := crand.Int(crand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			log.Println(functionName, "error generating bytes", err)
			return "", err
		}
		b[i] = lettersPassword[num.Int64()]
	}
	return string(b), nil
}

// randInt returns random integer between min and max.
func randInt(min int64, max int64) int64 {
	x, _ := crand.Int(crand.Reader, big.NewInt(max-min))
	return min + x.Int64()
}

// RandomString returns random string.
// If length not specified or zero, default length is 12.
func RandomString(length int) string {
	if length == 0 {
		length = 12
	}
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = byte(randInt(65, 90))
	}
	return string(b)
}

// UUIDGenerator generates and returns UUID.
func UUIDGenerator() string {
	return uuid.New().String()
}

// B2s converts byte to string.
// func B2s(b []byte) string {
// 	return *(*string)(unsafe.Pointer(&b))
// }

// S2b converts string to byte.
// func S2b(s string) (b []byte) {
// 	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
// 	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
// 	bh.Data = sh.Data
// 	bh.Cap = sh.Len
// 	bh.Len = sh.Len
// 	return b
// }
