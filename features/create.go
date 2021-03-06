package features

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CiaranAshton/features-go/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// CreateFeature creates a new feature and stores it in the database
func (fa FeatureAPI) CreateFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f := models.Feature{}

	json.NewDecoder(r.Body).Decode(&f)

	f.Id = bson.NewObjectId()

	err := fa.db.CreateFeature(fa.l, &f)

	if err != nil {
		fa.l.Err.Println("Unable to create feature")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Issue creating feature\n")
		return
	}

	fj, _ := json.Marshal(f)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", fj)
}
