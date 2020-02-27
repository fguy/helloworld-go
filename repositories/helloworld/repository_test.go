package helloworld

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mock_fx "github.com/fguy/helloworld-go/mocks/go.uber.org/fx"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetPage_Success(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		rows := sqlmock.NewRows([]string{
			"body",
		})
		rows.AddRow("world")
		mock.ExpectQuery(query).WillReturnRows(rows).RowsWillBeClosed()

		return db, nil
	})
	assert.NoError(t, err)

	result, err := instance.GetPage(
		context.Background(),
		"hello",
	)
	assert.NoError(t, err)
	assert.Equal(t, "hello", result.Title)
	assert.Equal(t, "world", result.Body)
}

func TestGetPage_NoRow(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		rows := sqlmock.NewRows([]string{
			"body",
		})
		mock.ExpectQuery(query).WillReturnRows(rows).RowsWillBeClosed()

		return db, nil
	})
	assert.NoError(t, err)

	result, err := instance.GetPage(
		context.Background(),
		"hello",
	)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetPage_Error(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		mock.ExpectQuery(query).WillReturnError(errors.New(""))

		return db, nil
	})
	assert.NoError(t, err)

	_, err = instance.GetPage(
		context.Background(),
		"hello",
	)
	assert.Error(t, err)
}

func TestGetPage_Error_DB(t *testing.T) {
	t.Parallel()

	_, err := New(nil, func() (*sql.DB, error) {
		return nil, errors.New("")
	})

	assert.Error(t, err)
}
