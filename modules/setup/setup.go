package setup

import(
	"net/http"
	"github.com/julienschmidt/httprouter"
	"tokopedia.se.training/Project2/vwa/helper/middleware"
)

type Self struct{}

func New() *Self {
	return &Self{}
}
func (self *Self) SetRouter(r *httprouter.Router) {
	/* register all router */

	mw := middleware.New() //implement middleware

	r.GET("/setup", mw.LoggingMiddleware(mw.CapturePanic(SetupHandler)))

}

func SetupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}
