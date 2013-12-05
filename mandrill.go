package mandrill

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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

type response struct {
	Email         string `json:"email"`
	Status        string `json:"status"`
	Reject_reason string `json:"reject_reason"`
	ID            string `json:"_id"`
}

type mError struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (m *MandrillCall) Send() (mResp []*response, mErr *mError) {
	m.Args["key"] = APIkey
	args, _ := json.Marshal(m.Args)
	log.Print(string(args))
	url := baseURL + "/" + m.Version + "/" + m.Category + "/" + m.Method + "." + m.Format
	client := &http.Client{}
	r, _ := http.NewRequest("POST", url, bytes.NewReader(args))
	resp, _ := client.Do(r)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if resp.StatusCode >= 400 {
		json.Unmarshal(buf.Bytes(), &mErr)
	} else {
		json.Unmarshal(buf.Bytes(), &mResp)
	}
	return mResp, mErr
}
