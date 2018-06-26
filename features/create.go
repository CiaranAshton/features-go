package features

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CiaranAshton/features/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// CreateFeature creates a new feature and stores it in the database
func (fa FeatureAPI) CreateFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f := models.Feature{}

	json.NewDecoder(r.Body).Decode(&f)

	f.Id = bson.NewObjectId()

	fa.session.DB("cjla").C("features").Insert(f)

	fj, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", fj)
}