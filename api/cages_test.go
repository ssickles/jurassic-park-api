package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"jurassic-park-api/data"
	"jurassic-park-api/data/mock"
	"jurassic-park-api/models"
	"jurassic-park-api/park"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testCages = []models.Cage{
	{
		Name:        "west-1",
		Capacity:    1,
		PowerStatus: models.PowerStatusActive,
	},
	{
		Name:        "east-1",
		Capacity:    2,
		PowerStatus: models.PowerStatusActive,
	},
	{
		Name:        "east-2",
		Capacity:    3,
		PowerStatus: models.PowerStatusDown,
	},
}

func TestCagesController_List(t *testing.T) {
	t.Run("no results", func(t *testing.T) {
		// set up the mock store and router
		store, err := mock.NewStore()
		assert.NoError(t, err)
		router := SetupRouter(store)

		// make the request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/cages", nil)
		router.ServeHTTP(w, req)

		// assert the response with no data
		assert.Equal(t, http.StatusOK, w.Code)
		var result map[string][]models.Cage
		err = json.Unmarshal(w.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Equal(t, map[string][]models.Cage{"data": nil}, result)
	})

	t.Run("3 results", func(t *testing.T) {
		// set up the mock store and router
		store, err := mock.NewStore()
		assert.NoError(t, err)
		router := SetupRouter(store)

		// load the store with 2 cages
		for _, cage := range testCages {
			_, err = store.Cages.Create(cage)
			assert.NoError(t, err)
		}

		// make the request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/cages", nil)
		router.ServeHTTP(w, req)

		// assert the response with 2 dinosaurs
		assert.Equal(t, http.StatusOK, w.Code)
		var result map[string][]models.Cage
		err = json.Unmarshal(w.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Equal(t, 3, len(result["data"]))
	})
}

func TestCagesController_Create(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(store data.Store)
		payload       park.CreateCagePayload
		expectedCode  int
		expectedError string
	}{
		{
			name: "successful creation of a cage",
			payload: park.CreateCagePayload{
				Name:     "east-1",
				Capacity: 1,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "attempt to create a cage with an invalid power status",
			payload: park.CreateCagePayload{
				Name:        "east-1",
				Capacity:    1,
				PowerStatus: "INVALID",
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: "Invalid cage power status: INVALID, must be ACTIVE or DOWN",
		},
		{
			name: "attempt to create a cage with a name that already exists",
			setup: func(store data.Store) {
				_, _ = store.Cages.Create(models.Cage{
					Name: "east-1",
				})
			},
			payload: park.CreateCagePayload{
				Name:     "east-1",
				Capacity: 1,
			},
			expectedCode:  http.StatusConflict,
			expectedError: "A cage with this name already exists: east-1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up the mock store and router
			store, err := mock.NewStore()
			assert.NoError(t, err)
			if test.setup != nil {
				test.setup(store)
			}
			router := SetupRouter(store)

			// make the request
			w := httptest.NewRecorder()
			body, err := json.Marshal(test.payload)
			assert.NoError(t, err)
			req, _ := http.NewRequest(http.MethodPost, "/cages", bytes.NewReader(body))
			router.ServeHTTP(w, req)

			// assert the response with the created cage
			assert.Equal(t, test.expectedCode, w.Code)
			if test.expectedCode == http.StatusCreated {
				var result map[string]models.Cage
				err = json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				expected := map[string]models.Cage{
					"data": {
						Id:       result["data"].Id,
						Name:     test.payload.Name,
						Capacity: test.payload.Capacity,
					},
				}
				assert.Equal(t, expected, result)
			}
			if test.expectedError != "" {
				var result map[string][]string
				err = json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				expected := map[string][]string{
					"errors": {test.expectedError},
				}
				assert.Equal(t, expected, result)
			}
		})
	}
}
