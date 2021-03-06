package features

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CiaranAshton/features-go/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// GetFeatures returns all features in the db
func (fa FeatureAPI) GetFeatures(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fs := []models.Feature{}

	err := fa.db.GetAllFeatures(fa.l, &fs)

	if err != nil {
		fa.l.Err.Println("Unable to find features")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Features not found\n")
		return
	}

	fsj, _ := json.Marshal(fs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", fsj)
}

// GetFeature returns status 200 and JSON of the desired feature
func (fa FeatureAPI) GetFeature(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		fa.l.Err.Printf("Id %s is not a valid Id \n", id)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid Id: %s", id)
		return
	}

	f := models.Feature{}

	err := fa.db.GetFeature(fa.l, id, &f)

	if err != nil {
		fa.l.Err.Println("Unable to find feature:", id)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Feature not found: %s", id)
		return
	}

	fj, _ := json.Marshal(f)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", fj)
}
