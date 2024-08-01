package repositories

import "strconv"

type IdMapper interface {
	ToInt64(id interface{}) int64
	ToInterface(id int64) interface{}
}

type Mapper struct {
}

func (m *Mapper) ToInt64(id interface{}) int64 {
	switch v := id.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return i
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return 0
		}
		return i
	default:
		return 0
	}
}
func (m *Mapper) ToInterface(id int64) (interface{}, error) {
	return id, nil
}
