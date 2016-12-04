package shared

import (
	"strings"
	"time"
)

type CCPTime struct {
	time.Time
}

func (ct CCPTime) String() string {
	s := ct.Time.Format(time.RFC3339)
	return s[:len(s)-1]
}

func (ct *CCPTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		s += "Z"
		t, err = time.Parse(time.RFC3339, s)
		if err != nil {
			t = time.Time{}
		}
	}
	ct.Time = t
	return nil
}

var nilTime = (time.Time{}).UnixNano()
