package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/controller/menu"
	"mus_projekt/app/controller/protocol"
	"net/http"
)

/**

package: app/controller ; file menu.go

Responsible for answering to HTTP Requests concerning the /api/menu interface


@author ben

*/

func MenuApi(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method == protocol.MethodFetch {
		fmt.Println("http:FETCH:MenuAPI")
		m := menu.CreateAPIData()
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
}
