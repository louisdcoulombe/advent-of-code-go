package cast

import (
	"fmt"
	"strconv"
)

// ToInt will case a given arg into an int type.
// Supported types are:
//   - string
func ToInt(arg interface{}) int {
	var val int
	switch arg := arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg)
		if err != nil {
			panic("error converting string to int " + err.Error())
		}

	case uint8:
		val = int(arg - '0')

	case rune:
		val = int(arg - '0')

	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}
