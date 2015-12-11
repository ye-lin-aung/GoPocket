package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"os"
	"strings"
)

type Action_add struct {
	Action string `json:"action"`
	ItemId string `json:"item_id"`
	Time   string `json:"time"`
}

type Modify_item struct {
	Consumer_key string   `json:"consumer_key"`
	Token        string   `json:"access_token"`
	Action_add   []string `json:"action"`
}

type Item struct {
	ItemId         string `json:"item_id,omitempty"`
	NormalUrl      string `json:"normal_url,omitempty"`
	ResolvedId     string `json:"resolved_id,omitempty"`
	ResolvedUrl    string `json:"resolved_url,omitempty"`
	DomainId       string `json:"domain_id,omitempty"`
	OriginDomainId string `json:"origin_domain_id,omitempty"`
	ResponseCode   string `json:"response_code,omitempty"`
	MimeType       string `json:"mime_type,omitempty"`
	ContentLength  string `json:"content_length,omitempty"`
	Encoding       string `json:"encoding,omitempty"`
	DateResolved   string `json:"date_resolved,omitempty"`
	DatePublished  string `json:"date_published,omitempty"`
	Title          string `json:"title,omitempty"`
	Excerpt        string `json:"excerpt,omitempty"`
	WordCount      string `json:"word_count,omitempty"`
	HasImage       string `json:"has_image,omitempty"`
	HasVideo       string `json:"has_video,omitempty"`
	IsIndex        string `json:"is_index,omitempty"`
	IsArticle      string `json:"is_article,omitempty"`
}

type Link struct {
	Url          string `json:"url"`
	Title        string `json:"title"`
	Time         string `json:"time"`
	Consumer_key string `json:"consumer_key"`
	Access_token string `json:"access_token"`
}

type RequestToken struct {
	Code string `json:"code"`
}
type AccessToken struct {
	Token string `json:"access_token"`
}
type StatusCode struct {
	Item          Item `json:"item,omitempty"`
	Status_result int  `json:"status"`
}

func main() {

	res := &RequestToken{}
	type AccessRe struct {
		Consumer_key string `json:"consumer_key"`
		Code         string `json:"code"`
	}
	type Re struct {
		Consumer_key string `json:"consumer_key"`
		Redirect_uri string `json:"redirect_uri"`
	}
	codet := Re{
		"48923-3169028ad25207cdae061c09",
		"www.google.com",
	}

	err := postJson("POST", codet, res, "https://getpocket.com/v3/oauth/request")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Code)
	s := "https://getpocket.com/auth/authorize?request_token=" + res.Code + "&&redirect_uri=www.google.com"

	open.Run(s)
	rs := &AccessToken{}
	access := AccessRe{
		"48923-3169028ad25207cdae061c09",
		res.Code,
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something")
	_, _ = reader.ReadString('\n')
	fmt.Println("Kwi Kwi")
	er := postJson("POST", access, rs, "https://getpocket.com/v3/oauth/authorize")

	if er != nil {
		panic(er)
	}
	fmt.Println(rs.Token)
	fmt.Println("Enter the file path")

	s, err = reader.ReadString('\n')
	if err != nil {

		panic(err)
	}
	openFile(s, rs.Token)
}
func openFile(path string, Token string) {

	fmt.Println(path)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addUrl("48923-3169028ad25207cdae061c09", Token, scanner.Text())

		lines = append(lines, scanner.Text())
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
func modify(Ckey string, Akey string) {
	//	data = &
	//	add := Modify_item{Ckey,Akey,}
	//	err := postJson("POST", add,data, "https://getpocket.com/v3/send")
	//	if err != nil{
	//	panic(err)
	//	}
	//	fmt.Println(data.)
}
func addUrl(Ckey string, Akey string, url string) {
	var res = &StatusCode{}

	l := Link{url, "Testing", "111111231", Ckey, Akey}
	err := postJson("POST", l, res, "https://getpocket.com/v3/add")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status_result)

}
func doJson(req *http.Request, res interface{}) error {
	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return (err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	return json.NewDecoder(resp.Body).Decode(res)
}

func postJson(action string, data, res interface{}, url string) error {

	body, err := json.Marshal(data)
	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(action, url, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	return doJson(req, res)

}
