package mandrill

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
)

var baseURL string = "https://mandrillapp.com/api/"
var Version string = "1.0"
var Format string = "json"
var APIkey string

type MandrillCall struct {
	Version  string
	Format   string
	APIkey   string
	Method   string
	Category string
	Args     map[string]interface{}
}

func NewCall() *MandrillCall {
	return &MandrillCall{
		Version: Version,
		Format:  Format,
		APIkey:  APIkey,
		Args:    make(map[string]interface{}),
	}
}

func (m *MandrillCall) Send() (*http.Response, error) {
	m.Args["key"] = APIkey
	args, _ := json.Marshal(m.Args)
	fmt.Println(string(args))
	url := baseURL + "/" + m.Version + "/" + m.Category + "/" + m.Method + "." + m.Format
	client := &http.Client{}
	r, _ := http.NewRequest("POST", url, bytes.NewReader(args))
	resp, err := client.Do(r)
	return resp, err
}
