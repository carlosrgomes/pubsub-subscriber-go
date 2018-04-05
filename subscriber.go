package subscriber

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//funcao principal do appengine
func init() {
	http.HandleFunc("/_ah/push-handlers/subscriber1", subscriber)

}

func subscriber(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	body, _ := ioutil.ReadAll(r.Body)

	log.Infof(ctx, "user %v ", string(body))

}
