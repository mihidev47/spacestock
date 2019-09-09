package repository

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"database/sql"

	"github.com/jmoiron/sqlx"

	"../logger"
)

var log = logger.Get()

// releaseTx clean db transaction by commit if no error, or rollback if an error occurred
func releaseTx(tx *sql.Tx, err *error) {
	if *err != nil {
		// If an error occurred, rollback transaction
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Errorf("Rollback error: %s", errRollback)
		} else {
			log.Info("Rolled Back")
		}
		return
	}
	// Else, commit transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		log.Errorf("Commit error: %s", errCommit)
	}
}

// txStmt determines whether to use transaction or not
//func txStmt(s *sqlx.Stmt, tx *sqlx.Tx) *sqlx.Stmt {
//	if tx != nil {
//		return tx.Stmtx(s)
//	}
//	return s
//}

// txNamedStmt determines whether to use transaction or not
func txNamedStmt(s *sqlx.NamedStmt, tx *sqlx.Tx) *sqlx.NamedStmt {
	if tx != nil {
		return tx.NamedStmt(s)
	}
	return s
}

// ReleaseTx clean db sqlx.Tx by commit if no error, or rollback if an error occurred
func ReleaseTx(tx *sqlx.Tx, err *error) {
	if *err != nil {
		// If an error occurred, rollback transaction
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Errorf("Rollback error: %s", errRollback)
		} else {
			log.Info("Rolled Back")
		}
		return
	}
	// Else, commit transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		log.Errorf("Commit error: %s", errCommit)
	}
}

// MustPrepare prepare sql statements or exit api if fails or error
func MustPrepare(db *sqlx.DB, query string) *sqlx.Stmt {
	s, err := db.Preparex(query)
	if err != nil {
		log.Errorf("Error preparing statement. Query: %s", query)
		log.Error(err)
		os.Exit(21)
	}
	return s
}

// MustPrepareNamed prepare sql statements with named bindvars or exit api if fails or error
func MustPrepareNamed(db *sqlx.DB, query string) *sqlx.NamedStmt {
	s, err := db.PrepareNamed(query)
	if err != nil {
		log.Errorf("Error preparing named statement. Query: %s", query)
		log.Error(err)
		os.Exit(22)
	}
	return s
}

// InsertBulkQuery prepares query, args for bulk insert
func InsertBulkQuery(tableName string, bindvar string, input interface{}) (query string,
	args []interface{}, err error) {
	// Get reflection value
	rv := reflect.ValueOf(input)
	// If input is not a slice, return error
	if rv.Kind() != reflect.Slice {
		return query, args, errors.New("input is not an array")
	}
	// if value is less than 1, return error
	if rv.Len() < 1 {
		return query, args, errors.New("values cannot be less than 1")
	}
	// Get columns reflection from sample
	cols, fields := getColumns(rv)
	// Get base query
	colQuery, bindvarQuery := getColumnQuery(cols, bindvar)
	// Generate base query
	query = "INSERT INTO " + tableName + colQuery + " VALUES "
	// Get values query and args
	valuesQuery, args := getValuesQueryArgs(fields, rv, bindvarQuery)
	// Append values query
	query += valuesQuery
	// Return
	return query, args, nil
}

func getColumns(rv reflect.Value) (columns []string, fields []string) {
	// Get sample
	var sample reflect.Value
	if k := rv.Kind(); k == reflect.Array || k == reflect.Slice {
		sample = rv.Index(0)
	} else {
		sample = rv
	}
	// Get field count
	var fieldCount int
	var fieldType reflect.Type
	if sample.Kind() == reflect.Ptr {
		elem := sample.Elem()
		fieldCount = elem.NumField()
		fieldType = elem.Type()
	} else {
		fieldCount = sample.NumField()
		fieldType = sample.Type()
	}
	// Filter fields for tagging
	for i := 0; i < fieldCount; i++ {
		f := fieldType.Field(i)
		// If field is an embedded
		if f.Anonymous && f.Type.Kind() == reflect.Ptr {
			embedColumns, embedFields := getColumns(sample.Field(i))
			// Merge columns and fields
			columns = append(columns, embedColumns...)
			fields = append(fields, embedFields...)
		} else if column := f.Tag.Get("db"); column != "-" && column != "" {
			columns = append(columns, column)
			fields = append(fields, f.Name)
		}
	}
	return columns, fields
}

func getColumnQuery(columns []string, bindvar string) (columnQuery string, bindvarQuery string) {
	// Init opener
	columnQuery = "("
	bindvarQuery = "("
	// index
	i := 1
	// Iterate columns
	for _, v := range columns {
		// Append column name
		columnQuery += v + ", "
		// Append bindvar query
		switch bindvar {
		case "$":
			bindvarQuery += "$" + string(i)
		case ":":
			bindvarQuery += ":" + v
		default:
			bindvarQuery += "?"
		}
		bindvarQuery += ", "
		// Increment index
		i++
	}
	// Trim extra separator and close group
	columnQuery = strings.TrimRight(columnQuery, ", ") + ")"
	bindvarQuery = strings.TrimRight(bindvarQuery, ", ") + ")"
	// Return queries
	return columnQuery, bindvarQuery
}

func getValuesQueryArgs(fields []string, values reflect.Value, bindvarQuery string) (query string, args []interface{}) {
	// Get value length
	valueCount := values.Len()
	// Init query
	q := make([]string, valueCount)
	// Iterate cols
	for i := 0; i < valueCount; i++ {
		// Append args query
		q[i] = bindvarQuery
		// Convert args to array
		a := getArgs(fields, values.Index(i))
		// Merge args
		args = append(args, a...)
	}
	// Generate full query
	query += strings.Join(q, ", ")
	query = strings.TrimRight(query, ", ")
	return query, args
}

func getArgs(fields []string, v reflect.Value) (args []interface{}) {
	// Get args by column name
	for _, fieldName := range fields {
		val := v.FieldByName(fieldName).Interface()
		args = append(args, val)
	}
	return args
}
