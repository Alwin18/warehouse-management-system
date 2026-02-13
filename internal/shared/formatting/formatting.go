package formatting

import (
	"strconv"
	"time"
)

func ConvertStringIntoUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(val)
}

// format po number like this PO-YYYYMMDD-XXXX
// example: PO-20260214-0001
func GeneratePONumber(count int64) string {
	return "PO-" + time.Now().Format("20060102") + "-" + strconv.FormatInt(count+1, 10)
}
