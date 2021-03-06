package features

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// DeleteFeature removes function from database
func (fa FeatureAPI) DeleteFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		fa.l.Err.Printf("Id %s is not a valid Id \n", id)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	err := fa.db.DeleteFeature(fa, oid)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Deleted Feature:", id)
}
