package main

import (
	"fmt"
	"net/http"
	"tokopedia.se.training/Project2/vwa/helper/middleware"
	"tokopedia.se.training/Project2/vwa/modules/product/komentar"
	"tokopedia.se.training/Project2/vwa/modules/user"
	"tokopedia.se.training/Project2/vwa/modules/user/profile"
	"tokopedia.se.training/Project2/vwa/util"
	"tokopedia.se.training/Project2/vwa/util/render"

	"github.com/julienschmidt/httprouter"
)


func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	data := make(map[string]interface{})
	data["title"] = "Home"
	render.HTMLRender(w,r, "template.index", data)
}

func main(){
	mw := middleware.New()
	router := httprouter.New()


	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	router.GET("/", mw.LoggingMiddleware(indexHandler))
	router.GET("/index", mw.LoggingMiddleware(indexHandler))

	user := user.New()
	user.SetRouter(router)
	
	komentar := komentar.New()
	komentar.SetRouter(router)

	profile := profile.New()
	profile.SetRouter(router)

	s := http.Server{
		Addr : ":"+util.Cfg.Webport,
		Handler : router,
	}

	fmt.Printf("Server running at port %s\n", s.Addr)
	fmt.Printf("Open this url %s on your browser to access VWA",util.Fullurl)
	s.ListenAndServe()
}