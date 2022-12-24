package main

type QueryOption func(*QueryOptions)
type QueryOptions struct {
	Sql         string
	Args        []any
	ElementType any
}

func loadOp(option ...QueryOption) *QueryOptions {
	options := new(QueryOptions)
	for _, e := range option {
		e(options)
	}
	return options
}

func QueryForMap(option ...QueryOption) (map[string]any, error) {
	op := loadOp(option...)
	rows, err := dataSource.Query(op.Sql, op.Args...)
	if err != nil {
		return nil, err
	}
	return GetScanMap(rows)
}
func QueryForObject[T any](option ...QueryOption) (t T, err error) {
	op := loadOp(option...)
	err = dataSource.Get(op.ElementType, op.Sql, op.Args...)
	if err == nil {
		t = op.ElementType.(T)
	}
	return
}

func QueryForMapList(option ...QueryOption) ([]map[string]any, error) {
	op := loadOp(option...)
	rows, err := dataSource.Query(op.Sql, op.Args...)
	if err != nil {
		return nil, err
	}

	return GetScanMapList(rows)
}
func QueryForObjectList[T any](option ...QueryOption) (t T, err error) {
	op := loadOp(option...)
	err = dataSource.Select(op.ElementType, op.Sql, op.Args...)
	if err == nil {
		t = op.ElementType.(T)
	}
	return
}
