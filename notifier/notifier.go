package notifier

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	noSkinProvided         = ""
	lebronJames            = "Lebron James"
	lebronJamesLoveMessage = "LETS GO!!!!! WAITING FOR LEBRON JAMES!!!\\n"
	otherSkinHateMessage   = "YOUR HATE FOR LEBRON JAMES SHALL BE PUNISHED, but ok looking for %s...\n"
	itemShopUrl            = "https://fnbr.co/api/shop"
	apiKeyHeader           = "x-api-key"
)

type Notifier struct {
	apiKey string
	skin   string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewNotifier(skin string, apiKey string) *Notifier {
	notifier := new(Notifier)
	notifier.apiKey = apiKey

	if skin == noSkinProvided || skin == lebronJames {
		notifier.skin = lebronJames
		fmt.Println(lebronJamesLoveMessage)
	} else {
		notifier.skin = skin
		fmt.Printf(otherSkinHateMessage, skin)
	}

	return notifier
}

func (notifier Notifier) Run() {
	req, err := http.NewRequest(http.MethodGet, itemShopUrl, nil)
	checkError(err)

	req.Header.Add(apiKeyHeader, notifier.apiKey)

	resp, err := http.DefaultClient.Do(req)
	checkError(err)

	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	checkError(err)

	err = resp.Body.Close()
	checkError(err)

	fmt.Printf("%s", j)
}
