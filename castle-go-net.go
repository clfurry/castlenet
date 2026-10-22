package castlenet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// endpoint of castle graphql

const qlendpoint string = "https://api.castle.xyz/graphql"

// IMPORTANT FOR THE THING
var auth = ""
var jsongetss = `{"variables":{},"query":"{\n  sharedDeckStorage(deckId: \"hereputdeckid\", keys: [\"hereputvar\"])\n}"}`
var jsonpostss = `{"variables":{},"query":"mutation {\n  updateSharedDeckStorage(deckId: \"hereputdeckid\", updates: [{key: \"hereputvar\", value: \"hereputvalue\"}])\n}"}`

// Pointer to net/http client, i use so it prevents 0 timeout (wait inf time) of default client

var cl = &http.Client{
	// 20 cuz castle isnt fast
	Timeout: time.Second * 20,
}

func LoadAuth(xauth string) {
	auth = xauth
}

// Internal Post HTTP for castle
func InternalPostHttp(to string, payl []byte) ([]byte, error) {
	// Best way is making the io.Reader in the func
	// New request better for auth cuz it makes headers
	// post method, to, io.reader
	req, err := http.NewRequest(http.MethodPost, to, bytes.NewReader(payl))
	// if err is found (not nil)
	if err != nil {
		fmt.Println("error: ", req, err)
		// returns nil and error so i can make things
		return nil, err
	}
	// header set to the request
	req.Header.Set("Content-Type", "application/json")
	// required
	req.Header.Set("X-Auth-Token", auth)

	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println("error: ", resp, err)
		// returns nil and error so i can make things
		return nil, err
	}
	// closes the response when func ends
	defer resp.Body.Close()

	//read all is recomended
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", string(body), err)
		// returns nil and error so i can make things
		return nil, err
	}
	// errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return body, fmt.Errorf("http err code %d: %s", resp.StatusCode, body)
	}
	// final return
	return body, nil
}

// As an alias and smth mor
func DecodeJson(body []byte, to any) error {
	return json.NewDecoder(bytes.NewReader(body)).Decode(to)
}

func GetSharedData(deckid string, tovar string) (any, error) {
	// temp to more easy
	tempcopy := jsongetss
	tempcopy = strings.ReplaceAll(tempcopy, "hereputdeckid", deckid)
	tempcopy = strings.ReplaceAll(tempcopy, "hereputvar", tovar)

	// yep, i copy the code
	body, err := InternalPostHttp(qlendpoint, []byte(tempcopy))
	if err != nil {
		fmt.Println("error: ", string(body), err)
		// returns nil and error so i can make things
		return nil, err
	}

	return body, err
}

func UpdateSharedData(deckid string, tovar string, val string) (any, error) {
	// temp to more easy
	tempcopy := jsonpostss
	tempcopy = strings.ReplaceAll(tempcopy, "hereputdeckid", deckid)
	tempcopy = strings.ReplaceAll(tempcopy, "hereputvar", tovar)
	tempcopy = strings.ReplaceAll(tempcopy, "hereputvalue", val)

	// yep, i copy the code
	body, err := InternalPostHttp(qlendpoint, []byte(tempcopy))
	if err != nil {
		fmt.Println("error: ", string(body), err)
		// returns nil and error so i can make things
		return nil, err
	}

	return body, err
}
