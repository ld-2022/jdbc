package jdbc

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"
)

func parseAsBytes(val interface{}) ([]byte, bool) {
	b, ok := val.([]byte)
	return b, ok
}

// Parse as Int
func parseInt(val interface{}) (int64, error) {
	bytes, ok := parseAsBytes(val)
	if !ok {
		return 0, fmt.Errorf("can't be parsed as int, type is %T", val)
	}
	return int64(binary.BigEndian.Uint64(bytes)), nil
}

// Parse as Float
func parseFloat(val interface{}) (float64, error) {
	bytes, ok := parseAsBytes(val)
	if !ok {
		return 0, fmt.Errorf("can't be parsed as float, type is %T", val)
	}
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits), nil
}

func GetScan(rows *sql.Rows) (map[string]interface{}, error) {
	defer rows.Close()

	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	numColumns := len(columns)

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	values := make([]interface{}, numColumns)
	for i := range values {
		values[i] = new(interface{})
	}

	if err := rows.Scan(values...); err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, numColumns)
	for i, column := range columns {
		//result[column] = *(values[i].(*interface{}))
		val := *(values[i].(*interface{}))
		switch column.DatabaseTypeName() {
		case "VARCHAR":
			if val != nil {
				val = string(val.([]byte))
			}
		}
		result[column.Name()] = val
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func GetScanMap(rows *sql.Rows) (map[string]interface{}, error) {
	defer rows.Close()
	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	numColumns := len(columns)
	if !rows.Next() {
		return nil, sql.ErrNoRows
	}
	values := make([]interface{}, numColumns)
	for i := range values {
		values[i] = new(interface{})
	}
	if err := rows.Scan(values...); err != nil {
		return nil, err
	}
	result := make(map[string]interface{}, numColumns)
	for i, column := range columns {
		val := *(values[i].(*interface{}))
		if val != nil {
			val = toVal(column, val)
		}
		result[column.Name()] = val
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func toVal(column *sql.ColumnType, val interface{}) interface{} {
	switch column.DatabaseTypeName() {
	case "VARCHAR", "CHAR", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT", "ENUM", "SET":
		if v, ok := val.([]byte); ok {
			val = string(v)
		}
	case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT", "BIT", "BOOL":
		if v, ok := val.([]byte); ok {
			if v, err := toUint64(v); err == nil {
				val = v
			}
		}
	case "FLOAT", "DOUBLE", "DECIMAL":
		if v, ok := val.([]uint8); ok {
			if v, err := toFloat64(v); err == nil {
				val = v
			}
		}
	case "DATE", "DATETIME", "TIMESTAMP", "YEAR", "TIME":
		switch v := val.(type) {
		case time.Time:
			val = v.Format("2006-01-02 15:04:05")
		case []byte:
			val = string(v)
		}
	case "BINARY", "VARBINARY", "BLOB", "TINYBLOB", "MEDIUMBLOB", "LONGBLOB":
		if v, ok := val.([]byte); ok {
			val = v
		}
	}
	return val
}

func toUint64(val []byte) (int64, error) {
	valStr := string(val)
	return strconv.ParseInt(valStr, 10, 64)
}

func toFloat64(val []uint8) (float64, error) {
	valStr := string(val)
	return strconv.ParseFloat(valStr, 64)
}

func GetScanMapList(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()

	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	numColumns := len(columns)

	resultList := make([]map[string]interface{}, 0)
	for rows.Next() {
		values := make([]interface{}, numColumns)
		for i := range values {
			values[i] = new(interface{})
		}

		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		result := make(map[string]interface{}, numColumns)
		for i, column := range columns {
			val := *(values[i].(*interface{}))
			if val != nil {
				val = toVal(column, val)
			}
			result[column.Name()] = val
		}
		resultList = append(resultList, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return resultList, nil
}
