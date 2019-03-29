package merge

import "reflect"

func getType(data interface{}) reflect.Kind {
	return reflect.TypeOf(data).Kind()
}

func mapEqual(m, m1 map[string]interface{}) bool {
	firstLength := 0
	for range m {
		firstLength += 1
	}
	count := 0
	for key, value := range m1 {
		if val, ok := m[key]; !ok || val == nil {
			return false
		} else {
			if value == nil {
				continue
			}
			if getType(val) == getType(value) {
				if getType(val) == reflect.Map {
					if !mapEqual(val.(map[string]interface{}), value.(map[string]interface{})) {
						return false
					}
				}
			}
		}
		count += 1
	}
	return count == firstLength
}

func Merge(origin, addition map[string]interface{}) map[string]interface{} {
	for key, value := range addition {
		if val, ok := origin[key]; ok && val != nil {
			if value == nil {
				continue
			}
			if getType(val) == getType(value) {
				if getType(val) == reflect.Map {
					origin[key] = Merge(value.(map[string]interface{}), val.(map[string]interface{}))
				} else if getType(val) == reflect.Slice {
					tmp := make([]interface{}, 0)
					for _, additionValue := range value.([]interface{}) {
						if additionValue == nil {
							continue
						}

						found := false
						for _, originValue := range val.([]interface{}) {
							if originValue == nil {
								continue
							}
							if getType(additionValue) == getType(originValue) {
								if getType(originValue) != reflect.Map ||
									getType(originValue) == reflect.Map &&
										mapEqual(originValue.(map[string]interface{}), additionValue.(map[string]interface{})) {
									found = true
									break
								}
							}

						}
						if !found {
							tmp = append(tmp, additionValue)
						}

					}
					origin[key] = append(origin[key].([]interface{}), tmp...)
				}
			} else {
				origin[key] = make([]interface{}, 0)
				origin[key] = append(origin[key].([]interface{}), val, value)
			}
		} else {
			origin[key] = value
		}
	}
	return origin
}
