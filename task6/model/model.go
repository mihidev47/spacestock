package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"../logger"
)

var log = logger.Get()

// parseJSON is a generic function to parse json
func parseJSON(src interface{}, target interface{}) error {
	// If source is nil, set target to nil
	if src == nil {
		target = nil
		return nil
	}
	// Assert source to byte
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to byte failed")
	}
	// Unmarshal to target
	err := json.Unmarshal(source, target)
	if err != nil {
		log.Errorf("Failed to unmarshal to %s", reflect.TypeOf(target))
		return err
	}
	return nil
}

func StringPtrToString(input *string) string {
	if input != nil {
		return *input
	}
	return ""
}

func StringPtrToNull(input *string) sql.NullString {
	if input != nil {
		return sql.NullString{*input, true}
	}
	return sql.NullString{}
}

func StringNullToPtr(input sql.NullString) *string {
	if input.Valid {
		return &input.String
	}
	return nil
}

func Float64PtrToNull(input *float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: *input,
		Valid:   input != nil,
	}
}

func Float64NullToPtr(input sql.NullFloat64) *float64 {
	if input.Valid {
		return &input.Float64
	}
	return nil
}

// setString ignores a valid NullString to be set null, otherwise it will update value and validation
func setString(val *sql.NullString, input string) {
	// Prevent string to be set null
	if val.Valid && val.String == "" {
		return
	}
	// Set email
	val.String = input
	// Set validation
	val.Valid = input != ""
}

type TypesItem struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type AppConfig struct {
	Id        string    `db:"id"`
	AppId     string    `db:"app_id"`
	Key       string    `db:"key"`
	Value     string    `db:"value"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ChangeLog struct {
	Init    bool `json:"init,omitempty"`
	Status  bool `json:"status,omitempty"`
	Notes   bool `json:"notes,omitempty"`
	Payment bool `json:"payment,omitempty"`
}

func (t *ChangeLog) Scan(src interface{}) error {
	return parseJSON(src, t)
}

func (t *ChangeLog) Value() (driver.Value, error) {
	return json.Marshal(t)
}
