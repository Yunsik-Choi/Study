package network

import (
	"fmt"
	"net/http"
	"strconv"
)

func init() {

}

func Home(w http.ResponseWriter, r *http.Request) {
	body, _ := r.Body.Read([]byte{})
	fmt.Fprintf(w, strconv.Itoa(body))
}
