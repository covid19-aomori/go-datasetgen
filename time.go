package datasetgen

import "time"

func ParseJapaneseDate(d string) (time.Time, error) {
	t, err := time.ParseInLocation("2006年1月2日", d, time.Local)
	return t, err
}

func TransformRFC3339Z(d string) (string, error) {
	t, err := ParseJapaneseDate(d)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02T15:04:05.000Z"), nil
}

func Transform20060102(d string) (string, error) {
	t, err := ParseJapaneseDate(d)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02"), nil
}

func Transform0102(d string) (string, error) {
	t, err := ParseJapaneseDate(d)
	if err != nil {
		return "", err
	}

	return t.Format("01/02"), nil
}

func TransformWeekday(d string) (int, error) {
	t, err := ParseJapaneseDate(d)
	if err != nil {
		return 0, err
	}

	return int(t.Weekday()), nil
}

func Transform200601021504(t time.Time) string {
	return t.Format("2006/01/02 15:04")
}
