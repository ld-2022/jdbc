package jdbc

import "github.com/ld-2022/jsonx"

func QueryForJSON(query string, args ...any) (*jsonx.JSONObject, error) {
	forMap, err := QueryForMap(query, args...)
	if err != nil {
		return nil, err
	}
	return jsonx.NewJSONObjectMap(forMap), nil
}

func QueryForJSONArray(query string, args ...any) (*jsonx.JSONArray, error) {
	forMapList, err := QueryForMapList(query, args...)
	if err != nil {
		return nil, err
	}
	var interfaceList []interface{}
	for _, forMap := range forMapList {
		interfaceList = append(interfaceList, jsonx.NewJSONObjectMap(forMap))
	}
	return jsonx.NewJSONArrayFromList(interfaceList), nil
}
