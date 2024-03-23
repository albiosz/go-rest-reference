package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/albiosz/honeycombs"
	"github.com/albiosz/honeycombs/internal/util"
	"github.com/albiosz/honeycombs/internal/util/test"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	db := test.SetupDB()
	mux := getRouter(db)
	server := httptest.NewServer(mux)

	client := http.Client{Timeout: time.Second * 10}

	t.Run("GET /users/{id}", func(t *testing.T) {
		tests := []struct {
			title               string
			userID              string
			expectedHTTPResCode int
		}{
			{
				title:               "200",
				userID:              "1",
				expectedHTTPResCode: http.StatusOK,
			}, {
				title:               "400 - incorrect parameter type",
				userID:              "string",
				expectedHTTPResCode: http.StatusBadRequest,
			}, {
				title:               "404 - user not found",
				userID:              "9999",
				expectedHTTPResCode: http.StatusNotFound,
			},
		}

		for _, test := range tests {
			t.Run(test.title, func(t *testing.T) {
				url := server.URL + "/users/" + test.userID
				req, err := http.NewRequest(http.MethodGet, url, nil)
				assert.NoError(t, err)

				req.Header.Set("Accept", "application/json")

				res, err := client.Do(req)
				assert.NoError(t, err)
				assert.Equal(t, test.expectedHTTPResCode, res.StatusCode)

				if res.StatusCode == http.StatusOK {
					var user honeycombs.User
					err = json.NewDecoder(res.Body).Decode(&user)
					assert.NoError(t, err)
					assert.NotEmpty(t, user)
				}
			})
		}
	})

	t.Run("POST /users", func(t *testing.T) {
		url := server.URL + "/users"
		userToCreate := honeycombs.User{
			Email:    "new@user.de",
			Password: "safePassword",
			Nickname: "NewUser",
		}

		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(&userToCreate)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, url, reqBody)
		assert.NoError(t, err)

		req.Header.Set("Accept", "application/json")

		res, err := client.Do(req)
		assert.NoError(t, err)

		var user honeycombs.User
		err = json.NewDecoder(res.Body).Decode(&user)
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
	})

	t.Run("PATCH /users/{id}", func(t *testing.T) {
		url := server.URL + "/users/1"
		updates := honeycombs.UserUpdate{
			Nickname: util.NewPtr("new-nick"),
		}

		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(&updates)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, url, reqBody)
		assert.NoError(t, err)

		req.Header.Set("Accept", "application/json")

		res, err := client.Do(req)
		assert.NoError(t, err)

		var updatedUser honeycombs.User
		err = json.NewDecoder(res.Body).Decode(&updatedUser)
		assert.NoError(t, err)
		assert.Equal(t, *updates.Nickname, updatedUser.Nickname)
	})

	t.Run("DELETE /users/{id}", func(t *testing.T) {
		url := server.URL + "/users/1"

		req, err := http.NewRequest(http.MethodDelete, url, nil)
		assert.NoError(t, err)

		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
