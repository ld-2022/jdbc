package jdbc

import "database/sql"

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
		case "VARCHAR":
			if val != nil {
				val = string(val.([]byte))
			}
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
			case "VARCHAR":
				if val != nil {
					val = string(val.([]byte))
				}
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
