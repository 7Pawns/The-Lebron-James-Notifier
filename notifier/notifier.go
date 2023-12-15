package notifier

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/go-toast/toast"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	noSkinProvided = ""
	lebronJames    = "Lebron James"
	itemShopUrl    = "https://fnbr.co/api/shop"
	apiKeyHeader   = "x-api-key"
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

	if skin == noSkinProvided {
		notifier.skin = lebronJames
	} else {
		notifier.skin = skin
	}

	return notifier
}

func (notifier Notifier) Run() {
	ticker := time.NewTicker(24 * time.Hour)
	shouldStop := make(chan bool)

	defer func() {
		firstTime := true
		for {
			if firstTime {
				firstTime = false
				notifier.checkItemShop()
			}
			select {
			case <-shouldStop:
				return
			case <-ticker.C:
				notifier.checkItemShop()
			}
		}
	}()
}

func (notifier Notifier) checkItemShop() {
	if notifier.isSkinInJson() {
		notifier.pushNotification()
	}
}

func makeRequest(notifier Notifier) []byte {
	req, err := http.NewRequest(http.MethodGet, itemShopUrl, nil)
	checkError(err)

	req.Header.Add(apiKeyHeader, notifier.apiKey)

	resp, err := http.DefaultClient.Do(req)
	checkError(err)

	body, err := io.ReadAll(resp.Body)
	checkError(err)

	err = resp.Body.Close()

	return body
}

func (notifier Notifier) isSkinInJson() bool {
	// You can use `ObjectEach` helper to iterate objects { "key1":object1, "key2":object2, .... "keyN":objectN }
	resp := makeRequest(notifier)
	var skinFound bool

	jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if curr, _, _, _ := jsonparser.Get(value, "name"); string(curr) == notifier.skin {
			skinFound = true
		}
	}, "data", "featured")

	return skinFound
}

func (notifier Notifier) pushNotification() {
	toast := toast.Notification{
		AppID:   "Lebron James Notifier",
		Message: fmt.Sprintf("%s is in the item shop!", notifier.skin),
	}

	err := toast.Push()
	checkError(err)
}
