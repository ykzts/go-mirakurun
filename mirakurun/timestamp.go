package mirakurun

import (
	"strconv"
	"time"
)

// Timestamp ...
type Timestamp struct {
	time.Time
}

// UnmarshalJSON ...
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	str := string(data)

	mSec, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	(*t).Time = time.Unix(mSec/1000, 0)

	return nil
}
