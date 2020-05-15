package log

import (
	"fmt"
	"strconv"
	"time"
)

func boolToStr(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func appendFields(buf []byte, args []interface{}) []byte {
	for ; len(args) > 0; args = args[2:] {
		buf = append(buf, ", "...)
		buf = append(buf, args[0].(string)...)
		buf = append(buf, '=')

		switch val := args[1].(type) {
		case bool:
			buf = append(buf, boolToStr(val)...)
		case string:
			buf = appendQuote(buf, val)
		case []string:
			buf = append(buf, '[')
			for i, item := range val {
				if i > 0 {
					buf = append(buf, ", "...)
				}
				buf = appendQuote(buf, item)
			}
			buf = append(buf, ']')
		case int:
			buf = strconv.AppendInt(buf, int64(val), 10)
		case int8:
			buf = strconv.AppendInt(buf, int64(val), 10)
		case int16:
			buf = strconv.AppendInt(buf, int64(val), 10)
		case int32:
			buf = strconv.AppendInt(buf, int64(val), 10)
		case int64:
			buf = strconv.AppendInt(buf, val, 10)
		case uint:
			buf = strconv.AppendUint(buf, uint64(val), 10)
		case uint8:
			buf = strconv.AppendUint(buf, uint64(val), 10)
		case uint16:
			buf = strconv.AppendUint(buf, uint64(val), 10)
		case uint32:
			buf = strconv.AppendUint(buf, uint64(val), 10)
		case uint64:
			buf = strconv.AppendUint(buf, val, 10)
		case time.Duration:
			buf = append(buf, val.String()...)
		case error:
			buf = appendQuote(buf, val.Error())
		case fmt.Stringer:
			buf = appendQuote(buf, val.String())
		default:
			panic(val)
		}
	}

	return buf
}
