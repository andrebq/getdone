package web

type Json map[string]interface{}

func (j Json) Put(key string, value interface{}) Json {
	j[key] = value
	return j
}

func (j Json) Get(key string) interface{} {
	if v, has := j[key]; has {
		return v
	} else {
		return nil
	}
	return j
}

func (j Json) Push(key string, value interface{}) Json {
	arr := j.Get(key)
	if arr == nil {
		arr = make([]interface{}, 0)
	}
	arr = append(arr.([]interface{}), value)
	j.Put(key, arr)
	return j
}
