package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"jurassic-park-api/data/mock"
	"jurassic-park-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testSpecies = []models.Species{
	{
		Name:     "Velociraptor",
		FoodType: "carnivore",
	},
	{
		Name:     "Tyrannosaurus",
		FoodType: "carnivore",
	},
}
var testDinosaurs = []models.Dinosaur{
	{
		Name:        "Valentino",
		SpeciesName: testSpecies[0].Name,
	},
	{
		Name:        "Rex",
		SpeciesName: testSpecies[1].Name,
	},
}

func TestDinosaursController_List(t *testing.T) {
	t.Run("there are no dinosaurs yet", func(t *testing.T) {
		// set up the mock store and router
		store, err := mock.NewStore()
		assert.NoError(t, err)
		router := SetupRouter(store)

		// make the request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/dinosaurs", nil)
		router.ServeHTTP(w, req)

		// assert the response with no data
		assert.Equal(t, http.StatusOK, w.Code)
		var result map[string][]models.Dinosaur
		err = json.Unmarshal(w.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Equal(t, map[string][]models.Dinosaur{"data": nil}, result)
	})

	t.Run("2 dinosaurs have been created", func(t *testing.T) {
		// set up the mock store and router
		store, err := mock.NewStore()
		assert.NoError(t, err)
		router := SetupRouter(store)

		// load the store with 2 dinosaurs
		for _, dinosaur := range testDinosaurs {
			_, err = store.Dinosaurs.Create(dinosaur)
			assert.NoError(t, err)
		}

		// make the request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/dinosaurs", nil)
		router.ServeHTTP(w, req)

		// assert the response with 2 dinosaurs
		assert.Equal(t, http.StatusOK, w.Code)
		var result map[string][]models.Dinosaur
		err = json.Unmarshal(w.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(result["data"]))
	})
}

func TestDinosaursController_Create(t *testing.T) {
	t.Run("successful creation of a dinosaur", func(t *testing.T) {
		// set up the mock store and router
		store, err := mock.NewStore()
		// make sure we have our species loaded up in our data store
		for _, species := range testSpecies {
			_, err = store.Species.Create(species)
			assert.NoError(t, err)
		}
		assert.NoError(t, err)
		router := SetupRouter(store)

		// make the request
		w := httptest.NewRecorder()
		body, err := json.Marshal(testDinosaurs[0])
		assert.NoError(t, err)
		req, _ := http.NewRequest(http.MethodPost, "/dinosaurs", bytes.NewReader(body))
		router.ServeHTTP(w, req)

		// assert the response with the created dinosaur
		assert.Equal(t, http.StatusCreated, w.Code)
		var result map[string]models.Dinosaur
		err = json.Unmarshal(w.Body.Bytes(), &result)
		assert.NoError(t, err)
		expected := map[string]models.Dinosaur{
			"data": {
				Id:          result["data"].Id,
				Name:        testDinosaurs[0].Name,
				SpeciesName: testDinosaurs[0].SpeciesName,
			},
		}
		assert.Equal(t, expected, result)
	})
}
