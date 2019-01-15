package httphelper

import (
	"fmt"
	"net/http"
	"tokopedia.se.training/Project2/vwa/util"
)

func Redirect(w http.ResponseWriter, r *http.Request, location string, code int){
	redirect := fmt.Sprintf("%s%s", util.Fullurl,location)
	http.Redirect(w,r,redirect,code)
}