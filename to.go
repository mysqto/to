package to

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strconv"
)

// String change val type to string
func String(val interface{}) string {
	if val == nil {
		return "nil"
	}

	switch t := val.(type) {
	case bool:
		return strconv.FormatBool(t)
	case int:
		return strconv.FormatInt(int64(t), 10)
	case int8:
		return strconv.FormatInt(int64(t), 10)
	case int16:
		return strconv.FormatInt(int64(t), 10)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case int64:
		return strconv.FormatInt(t, 10)
	case uint:
		return strconv.FormatUint(uint64(t), 10)
	case uint8:
		return strconv.FormatUint(uint64(t), 10)
	case uint16:
		return strconv.FormatUint(uint64(t), 10)
	case uint32:
		return strconv.FormatUint(uint64(t), 10)
	case uint64:
		return strconv.FormatUint(t, 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case []byte:
		return string(t)
	case string:
		return t
	default:
		data, err := json.Marshal(val)
		if err != nil {
			return fmt.Sprintf("%v", val)
		}
		return string(data)
	}
}

//Atoi change val type to int64
func Atoi(val interface{}) int64 {
	if val == nil {
		return 0
	}

	switch t := val.(type) {
	case bool:
		if t {
			return int64(1)
		}
		return int64(0)
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return int64(t)
	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case []byte:
		i, _ := strconv.Atoi(string(t))
		return int64(i)
	case string:
		b, err := strconv.ParseBool(t)
		if err == nil {
			if b {
				return int64(1)
			}

			return int64(0)
		}
		i, _ := strconv.ParseFloat(t, 64)
		return int64(i)
	default:
		i, _ := strconv.ParseFloat(fmt.Sprintf("%v", t), 64)
		return int64(i)
	}
}

const hexTable = "0123456789ABCDEF"

// encode encodes src into EncodedLen(len(src))
// bytes of dst. As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// encode implements hexadecimal encoding.
func encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*5+0] = '0'
		dst[i*5+1] = 'x'
		dst[i*5+2] = hexTable[v>>4]
		dst[i*5+3] = hexTable[v&0x0f]
		if i != len(src)-1 {
			dst[i*5+4] = ' '
		}
	}

	return len(src)*5 - 1
}

// encodedLen returns the length of an encoding of n source bytes.
// Specifically, it returns n * 5.
func encodedLen(n int) int { return n*5 - 1 }

func getBytes(data interface{}) ([]byte, error) {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// HexString returns the hexadecimal encoding of src.
func HexString(v interface{}) string {
	var hex string
	src, err := getBytes(v)
	if err != nil {
		return hex
	}

	dst := make([]byte, encodedLen(len(src)))
	encode(dst, src)
	return string(dst)
}
