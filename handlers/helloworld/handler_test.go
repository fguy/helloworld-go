package helloworld

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/fguy/helloworld-go/entities"
	mock_repo "github.com/fguy/helloworld-go/mocks/repositories/helloworld"
	"github.com/golang/mock/gomock"
)

func TestServeHTTP_Success(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock_repo.NewMockInterface(mockCtrl)
	mockRepo.EXPECT().GetPage(gomock.Any(), "hello").Return(&entities.Page{
		Title: "hello",
		Body:  "world",
	}, nil).Times(1)

	handler := NewHandler(zap.NewNop(), mockRepo)
	server := httptest.NewServer(handler)
	defer server.Close()

	res, err := server.Client().Get(fmt.Sprint(server.URL, "/hello"))
	assert.NoError(t, err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var page entities.Page
	json.Unmarshal(body, &page)
	assert.Equal(t, "hello", page.Title)
	assert.Equal(t, "world", page.Body)
}

func TestServeHTTP_Error(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock_repo.NewMockInterface(mockCtrl)
	mockRepo.EXPECT().GetPage(gomock.Any(), "hello").Return(nil, errors.New("")).Times(1)

	handler := NewHandler(zap.NewNop(), mockRepo)
	server := httptest.NewServer(handler)
	defer server.Close()

	res, err := server.Client().Get(fmt.Sprint(server.URL, "/hello"))
	assert.NoError(t, err)
	assert.Equal(t, 500, res.StatusCode)
}

func TestServeHTTP_NotFound(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock_repo.NewMockInterface(mockCtrl)
	mockRepo.EXPECT().GetPage(gomock.Any(), "hello").Return(nil, nil).Times(1)

	handler := NewHandler(zap.NewNop(), mockRepo)
	server := httptest.NewServer(handler)
	defer server.Close()

	res, err := server.Client().Get(fmt.Sprint(server.URL, "/hello"))
	assert.NoError(t, err)
	assert.Equal(t, 404, res.StatusCode)
}
