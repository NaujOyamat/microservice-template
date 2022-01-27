package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Github.com/NaujOyamat/microservice-template/internal/infrastructure/storage/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(mocks.ICourseRepository)
	courseRepository.On("Save", mock.Anything, mock.Anything).
		Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(courseRepository))

	t.Run("BadRequest", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:   "dd94dd49-0e92-4e84-ad77-74d0896e4963",
			Name: "Demo Course",
		}

		body, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(body))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("StatusCreated", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "dd94dd49-0e92-4e84-ad77-74d0896e4963",
			Name:     "Demo Course",
			Duration: "10 Months",
		}

		body, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(body))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
