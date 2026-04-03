package handlers

import (
	"fmt"
	"net/http"
)

func Login( w http.ResponseWriter , r *http.Request){
	s :="qsqdqsdqsd"
	fmt.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	w.Write([]byte(s))
}
