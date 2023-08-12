package jdbc

func QueryForMap(query string, args ...any) (map[string]any, error) {
	rows, err := dataSource.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return GetScanMap(rows)
}

func QueryForMapList(query string, args ...any) ([]map[string]any, error) {
	rows, err := dataSource.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return GetScanMapList(rows)
}
