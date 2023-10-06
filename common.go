package jdbc

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"math"
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
		switch column.DatabaseTypeName() {
		case "VARCHAR", "CHAR", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT", "ENUM", "SET":
			if val != nil {
				val = string(val.([]byte))
			}
		case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT",
			"BIT", "BOOL":
			//if val != nil {
			//	intVal, err := parseInt(val)
			//	if err != nil {
			//		// Handle the error
			//	}
			//	val = intVal
			//}
		case "FLOAT", "DOUBLE", "DECIMAL":
			//if val != nil {
			//	floatVal, err := parseFloat(val)
			//	if err != nil {
			//		// Handle the error
			//	}
			//	val = floatVal
			//}
		case "DATE", "DATETIME", "TIMESTAMP", "YEAR", "TIME":
			if val != nil {
				switch val.(type) {
				case time.Time:
					val = val.(time.Time).Format("2006-01-02 15:04:05")
				default:
					val = string(val.([]byte))
				}
			}
		case "BINARY", "VARBINARY", "BLOB", "TINYBLOB", "MEDIUMBLOB", "LONGBLOB":
			if val != nil {
				val = val.([]byte)
			}
		default:
			// No action for other types
		}
		result[column.Name()] = val
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
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
			switch column.DatabaseTypeName() {
			case "VARCHAR", "CHAR", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT", "ENUM", "SET":
				if val != nil {
					val = string(val.([]byte))
				}
			case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT",
				"BIT", "BOOL":
				//if val != nil {
				//	intVal, err := parseInt(val)
				//	if err != nil {
				//		// Handle the error
				//	}
				//	val = intVal
				//}
			case "FLOAT", "DOUBLE", "DECIMAL":
				/*if val != nil {
					floatVal, err := parseFloat(val)
					if err != nil {
						// Handle the error
					}
					val = floatVal
				}*/
			case "DATE", "DATETIME", "TIMESTAMP", "YEAR", "TIME":
				if val != nil {
					switch val.(type) {
					case time.Time:
						val = val.(time.Time).Format("2006-01-02 15:04:05")
					default:
						val = string(val.([]byte))
					}
				}
			case "BINARY", "VARBINARY", "BLOB", "TINYBLOB", "MEDIUMBLOB", "LONGBLOB":
				if val != nil {
					val = val.([]byte)
				}
			default:
				// No action for other types
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
