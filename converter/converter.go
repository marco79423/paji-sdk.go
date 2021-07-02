package converter

import (
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

func ConvertIntToIntRef(int int) *int {
	return &int
}

func ConvertInt64ToInt64Ref(int64 int64) *int64 {
	return &int64
}

func ConvertInt64ToStr(int64 int64) string {
	return strconv.FormatInt(int64, 10)
}

func ConvertIntToStr(int int) string {
	return strconv.Itoa(int)
}

func ConvertStrToStrRef(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func ConvertStrRefToStr(strRef *string) string {
	if strRef == nil {
		return ""
	}

	return *strRef
}

func ConvertStrToInt(intStr string) (int, error) {
	targetInt64, err := strconv.ParseInt(intStr, 10, 0)
	if err != nil {
		return 0, xerrors.Errorf("無法將字串 %s 轉換成 int: %w", intStr, err)
	}
	return int(targetInt64), nil
}

func ConvertStrToInt64(intStr string) (int64, error) {
	targetInt64, err := strconv.ParseInt(intStr, 10, 0)
	if err != nil {
		return 0, xerrors.Errorf("無法將字串 %s 轉換成 int: %w", intStr, err)
	}
	return targetInt64, nil
}

func ConvertStrToIntRef(intStr string) (*int, error) {
	if intStr == "" {
		return nil, nil
	}

	targetInt64, err := strconv.ParseInt(intStr, 10, 0)
	if err != nil {
		return nil, xerrors.Errorf("無法將字串 %s 轉換成 *int: %w", intStr, err)
	}

	targetInt := int(targetInt64)

	return &targetInt, nil
}

func ConvertStrToFloat64(float64Str string) (float64, error) {
	targetFloat64, err := strconv.ParseFloat(float64Str, 64)
	if err != nil {
		return 0, xerrors.Errorf("無法將字串 %s 轉換成 float64: %w", float64Str, err)
	}

	return targetFloat64, nil
}

func ConvertStrToTime(timeStr string) (time.Time, error) {
	targetTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, xerrors.Errorf("無法將字串 %s 轉換成 time.Time: %w", timeStr, err)
	}

	return targetTime, nil
}

func ConvertStrToTimeRef(timeStr string) (*time.Time, error) {
	if timeStr == "" {
		return nil, nil
	}

	targetTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return nil, xerrors.Errorf("無法將字串 %s 轉換成 *time.Time: %w", timeStr, err)
	}

	return &targetTime, nil
}

func ConvertStrToBool(boolStr string) bool {
	return boolStr == "1"
}

func ConvertTimeToTimeRef(time time.Time) *time.Time {
	return &time
}

func ConvertFloat64ToStr(float64 float64) string {
	return strconv.FormatFloat(float64, 'f', -1, 64)
}
