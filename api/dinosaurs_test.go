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

var testDinosaurs = []models.Dinosaur{
	{
		Name:    "Valentino",
		Species: "velociraptor",
	},
	{
		Name:    "Rex",
		Species: "tyrannosaurus",
	},
}

func TestDinosaursController_List(t *testing.T) {
	t.Run("no results", func(t *testing.T) {
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

	t.Run("2 results", func(t *testing.T) {
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
	t.Run("successful create", func(t *testing.T) {
		// set up the mock store and router
		store, err := mock.NewStore()
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
				Id:      result["data"].Id,
				Name:    testDinosaurs[0].Name,
				Species: testDinosaurs[0].Species,
			},
		}
		assert.Equal(t, expected, result)
	})
}
