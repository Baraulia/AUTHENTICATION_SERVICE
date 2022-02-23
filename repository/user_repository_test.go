package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/model"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"testing"
	"time"
)

var logger = logging.GetLogger()

type AnyTime struct {
	time.Time
}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
func (a AnyTime) Value() (driver.Value, error) {
	return driver.Value(a.Time), nil
}

func TestRepository_GetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db, logger)

	testTable := []struct {
		name          string
		mock          func(id int)
		id            int
		expectedUser  *model.ResponseUser
		expectedError bool
	}{
		{
			name: "OK",
			mock: func(id int) {
				rows := sqlmock.NewRows([]string{"id", "email", "created_at"}).
					AddRow(1, "test@yandex.ru", AnyTime{})
				mock.ExpectQuery("SELECT id, email, created_at FROM users WHERE id = (.+)").
					WithArgs(id).WillReturnRows(rows)
			},
			id: 1,
			expectedUser: &model.ResponseUser{
				ID:        1,
				Email:     "test@yandex.ru",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: false,
		},
		{
			name: "Not found",
			mock: func(id int) {
				rows := sqlmock.NewRows([]string{"id", "email", "created_at"})

				mock.ExpectQuery("SELECT id, email, created_at FROM users WHERE id = (.+)").
					WithArgs(id).WillReturnRows(rows)

			},
			id:            1,
			expectedUser:  nil,
			expectedError: true,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.id)
			got, err := r.GetUserByID(tt.id)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetUserAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db, logger)

	testTable := []struct {
		name          string
		inputPage     int
		inputLimit    int
		mock          func(page, limit int)
		expectedUser  []model.ResponseUser
		expectedPages int
		expectedError bool
	}{
		{
			name:       "Zero page and limit",
			inputPage:  0,
			inputLimit: 0,
			mock: func(page, limit int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id", "email", "created_at"}).
					AddRow(1, "test@yandex.ru", AnyTime{}).
					AddRow(2, "test1@yandex.ru", AnyTime{}).
					AddRow(3, "test2@yandex.ru", AnyTime{})
				mock.ExpectQuery("SELECT id, email, created_at FROM users").WillReturnRows(rows)
				mock.ExpectCommit()
			},

			expectedUser: []model.ResponseUser{
				{
					ID:        1,
					Email:     "test@yandex.ru",
					CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					Email:     "test1@yandex.ru",
					CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        3,
					Email:     "test2@yandex.ru",
					CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedError: false,
		},
		{
			name:       "OK",
			inputPage:  1,
			inputLimit: 10,
			mock: func(page, limit int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id", "email", "created_at"}).
					AddRow(1, "test@yandex.ru", AnyTime{}).
					AddRow(2, "test1@yandex.ru", AnyTime{}).
					AddRow(3, "test2@yandex.ru", AnyTime{})

				mock.ExpectQuery("SELECT id, email, created_at FROM users ORDER BY id LIMIT (.+) OFFSET (.+)").WithArgs(limit, (page-1)*limit).WillReturnRows(rows)
				rows2 := sqlmock.NewRows([]string{"pages"}).
					AddRow(1)
				mock.ExpectQuery("SELECT CEILING").WillReturnRows(rows2)
				mock.ExpectCommit()
			},

			expectedUser: []model.ResponseUser{
				{
					ID:        1,
					Email:     "test@yandex.ru",
					CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					Email:     "test1@yandex.ru",
					CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        3,
					Email:     "test2@yandex.ru",
					CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedError: false,
		},
		{
			name:       "db error",
			inputPage:  1,
			inputLimit: 10,
			mock: func(page, limit int) {
				mock.ExpectBegin().WillReturnError(errors.New("some error"))
			},
			expectedUser:  nil,
			expectedError: true,
		},
		{
			name:       "db error2",
			inputPage:  1,
			inputLimit: 10,
			mock: func(page, limit int) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT id, email, created_at FROM users ORDER BY id LIMIT (.+) OFFSET (.+)").WithArgs(limit, (page-1)*limit).WillReturnError(errors.New("some error"))
			},
			expectedUser:  nil,
			expectedError: true,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.inputPage, tt.inputLimit)
			got, _, err := r.GetUserAll(tt.inputPage, tt.inputLimit)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db, logger)

	testTable := []struct {
		name          string
		mock          func(email string)
		email         string
		expectedUser  *model.User
		expectedError bool
	}{
		{
			name: "OK",
			mock: func(email string) {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at"}).
					AddRow(1, "test@yandex.ru", "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy", AnyTime{})

				mock.ExpectQuery("SELECT id, email, password, created_at FROM users WHERE email = (.+)").
					WithArgs(email).WillReturnRows(rows)
			},
			email: "test@yandex.ru",
			expectedUser: &model.User{
				ID:        1,
				Email:     "test@yandex.ru",
				Password:  "$2a$10$ooCmcWnLIubagB1MqM3UWOIpJTrq58tPQO6HVraj3yTKASiXBXHqy",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: false,
		},
		{
			name: "Not found",
			mock: func(email string) {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at"})

				mock.ExpectQuery("SELECT id, email, password, created_at FROM users WHERE email = (.+)").
					WithArgs(email).WillReturnRows(rows).WillReturnError(errors.New("some error"))

			},
			email:         "test@yandex.ru",
			expectedUser:  nil,
			expectedError: true,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.email)
			got, err := r.GetUserByEmail(tt.email)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_DeleteUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db, logger)

	testTable := []struct {
		name           string
		mock           func(id int)
		id             int
		expectedUserId int
		expectedError  bool
	}{
		{
			name: "OK",
			mock: func(id int) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				mock.ExpectQuery("DELETE FROM users WHERE id=(.+) RETURNING id").
					WithArgs(id).WillReturnRows(rows)
			},
			id:             1,
			expectedUserId: 1,
			expectedError:  false,
		},
		{
			name: "Not found",
			mock: func(id int) {
				rows := sqlmock.NewRows([]string{"id"})

				mock.ExpectQuery("DELETE FROM users WHERE id=(.+) RETURNING id").
					WithArgs(id).WillReturnRows(rows)

			},
			id:             1,
			expectedUserId: 0,
			expectedError:  true,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.id)
			got, err := r.DeleteUserByID(tt.id)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUserId, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db, logger)

	testTable := []struct {
		name           string
		mock           func(user *model.CreateUser)
		InputUser      *model.CreateUser
		expectedUserId int
		expectedError  bool
	}{
		{
			name: "OK",
			mock: func(user *model.CreateUser) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				mock.ExpectQuery("INSERT INTO users").WithArgs(user.Email, user.Password, AnyTime{}).WillReturnRows(rows)
			},
			InputUser: &model.CreateUser{
				Email:    "test@yandex.ru",
				Password: "$2a$10$EpAGhm0HGkxBiPyBAB7xzuyEbZlZCjvSdcJTjamaJyxZRir1vaMmW",
			},
			expectedUserId: 1,
			expectedError:  false,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.InputUser)
			got, err := r.CreateUser(tt.InputUser)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUserId, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_UpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db, logger)

	testTable := []struct {
		name          string
		mock          func(user *model.UpdateUser, id int)
		InputUser     *model.UpdateUser
		inputId       int
		expectedError bool
	}{
		{
			name: "OK",
			mock: func(user *model.UpdateUser, id int) {
				result := driver.ResultNoRows
				mock.ExpectExec("UPDATE users").
					WithArgs(user.NewPassword, id).WillReturnResult(result)
			},
			InputUser: &model.UpdateUser{
				Email:       "test@yandex.ru",
				OldPassword: "$2a$10$EpAGhm0HGkxBiPyBAB7xzuyEbZlZCjvSdcJTjamaJyxZRir1vaMmW",
				NewPassword: "$2a$10$EpAGhm0HGkxBiPyBAB7xzuyEbZlZCjvSdcJTjamaJyxZRir1vaMmW",
			},
			inputId:       1,
			expectedError: false,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.InputUser, tt.inputId)
			err := r.UpdateUser(tt.InputUser, tt.inputId)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
