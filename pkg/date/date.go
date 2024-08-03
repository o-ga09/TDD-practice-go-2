package date

import (
	"time"
)

func TimeToString(t string) (time.Time, error) {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		return time.Now(), err
	}
	return parsedTime, nil
}
