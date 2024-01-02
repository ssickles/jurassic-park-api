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

var testSpecies = []models.Species{
	{
		Name:     "Velociraptor",
		FoodType: "carnivore",
	},
	{
		Name:     "Tyrannosaurus",
		FoodType: "carnivore",
	},
	{
		Name:     "Brachiosaurus",
		FoodType: "herbivore",
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
	tests := []struct {
		name          string
		setup         func(store data.Store)
		payload       park.CreateDinosaurPayload
		expectedCode  int
		expectedError string
	}{
		{
			name: "successful creation of a dinosaur",
			payload: park.CreateDinosaurPayload{
				Name:        "Valentino",
				SpeciesName: "Velociraptor",
				CageName:    "east-1",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "attempt to create a dinosaur with a name that already exists",
			setup: func(store data.Store) {
				_, _ = store.Dinosaurs.Create(models.Dinosaur{
					Name:        "Valentino",
					SpeciesName: "Velociraptor",
				})
			},
			payload: park.CreateDinosaurPayload{
				Name:        "Valentino",
				SpeciesName: "Velociraptor",
			},
			expectedCode:  http.StatusConflict,
			expectedError: "A dinosaur with this name already exists: Valentino",
		},
		{
			name: "attempt to create a dinosaur with a species that doesn't exist",
			payload: park.CreateDinosaurPayload{
				Name:        "Valentino",
				SpeciesName: "BadSpeciesName",
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: "Invalid species name: BadSpeciesName",
		},
		{
			name: "attempt to create a dinosaur with a cage that doesn't exist",
			payload: park.CreateDinosaurPayload{
				Name:        "Valentino",
				SpeciesName: "Velociraptor",
				CageName:    "BadCageName",
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: "Invalid cage name: BadCageName",
		},
		{
			name: "attempt to create a dinosaur with a cage that is at capacity",
			setup: func(store data.Store) {
				cage, _ := store.Cages.FindByName("west-1")
				testDinosaurs[0].CageId = cage.Id
				_, _ = store.Dinosaurs.Create(testDinosaurs[0])
			},
			payload: park.CreateDinosaurPayload{
				Name:        "Rex",
				SpeciesName: "Tyrannosaurus",
				CageName:    "west-1",
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: "The cage is already at capacity (1), can't add another dinosaur",
		},
		{
			name: "attempt to create a dinosaur with a food type that doesn't match the cage's food type",
			setup: func(store data.Store) {
				cage, _ := store.Cages.FindByName("east-1")
				_, _ = store.Dinosaurs.Create(models.Dinosaur{
					Name:        "Brett",
					SpeciesName: "Brachiosaurus",
					CageId:      cage.Id,
				})
			},
			payload: park.CreateDinosaurPayload{
				Name:        "Rex",
				SpeciesName: "Tyrannosaurus",
				CageName:    "east-1",
			},
			expectedCode:  http.StatusBadRequest,
			expectedError: "The dinosaur's food type (carnivore) does not match the cage's food type (herbivore)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up the mock store and router
			store, err := mock.NewStore()
			// make sure we have our species and cages loaded up in our data store
			for _, species := range testSpecies {
				_, err = store.Species.Create(species)
				assert.NoError(t, err)
			}
			for _, cage := range testCages {
				_, err = store.Cages.Create(cage)
				assert.NoError(t, err)
			}
			assert.NoError(t, err)
			if test.setup != nil {
				test.setup(store)
			}
			router := SetupRouter(store)

			// make the request
			w := httptest.NewRecorder()
			body, err := json.Marshal(test.payload)
			assert.NoError(t, err)
			req, _ := http.NewRequest(http.MethodPost, "/dinosaurs", bytes.NewReader(body))
			router.ServeHTTP(w, req)

			// assert the response with the created dinosaur
			assert.Equal(t, test.expectedCode, w.Code)
			if test.expectedCode == http.StatusCreated {
				var result map[string]models.Dinosaur
				err = json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				expected := map[string]models.Dinosaur{
					"data": {
						Id:          result["data"].Id,
						Name:        test.payload.Name,
						SpeciesName: test.payload.SpeciesName,
						CageId:      result["data"].CageId,
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
