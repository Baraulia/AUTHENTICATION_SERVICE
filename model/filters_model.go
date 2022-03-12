package model

import (
	"database/sql/driver"
	"fmt"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"strings"
	"time"
)

var logger logging.Logger

const Layout = "2006-01-02"

type MyTime struct {
	time.Time
}

func (c *MyTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" || string(data) == "" {
		logger.Error("date  is not specified")
		return fmt.Errorf("date is not specified")
	} else {
		s := strings.Trim(string(data), "\"")
		// Fractional seconds are handled implicitly by Parse.
		tt, err := time.Parse(Layout, s)
		*c = MyTime{tt}
		return err
	}
}

func (c MyTime) Value() (driver.Value, error) {
	return driver.Value(c.Time), nil
}

func (c *MyTime) Scan(src interface{}) error {
	switch t := src.(type) {
	case time.Time:
		c.Time = t
		return nil
	default:
		return fmt.Errorf("column type not supported")
	}
}
func (c MyTime) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(Layout))), nil
}

type RequestFilters struct {
	ShowDeleted bool   `json:"show_deleted"`
	FilterData  bool   `json:"filter_data"`
	StartTime   MyTime `json:"start_time"`
	EndTime     MyTime `json:"end_time"`
	FilterRole  string `json:"filter_role"`
}

type SwaggerRequestFilters struct {
	ShowDeleted bool   `json:"show_deleted"`
	FilterData  bool   `json:"filter_data"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	FilterRole  string `json:"filter_role"`
}
