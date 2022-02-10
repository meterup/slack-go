package slack

import (
	"encoding/json"
	"math"
	"strconv"
	"time"
)

type SlackTime struct {
	Time      time.Time
	Timestamp string
}

func (s *SlackTime) UnmarshalJSON(data []byte) error {
	var f float64
	var ss string
	err := json.Unmarshal(data, &ss)
	if err == nil {
		f, err = strconv.ParseFloat(ss, 64)
		if err != nil {
			return err
		}
	} else if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	intpart, divpart := math.Modf(f)
	*s = SlackTime{
		Time:      time.Unix(int64(intpart), int64(divpart*100*1000*1000)),
		Timestamp: ss,
	}
	return nil
}

func (s SlackTime) MarshalJSON() ([]byte, error) {
	ts := float64(time.Time(s.Time).Unix()) + float64(time.Time(s.Time).UnixNano())/100*1000*1000
	return json.Marshal(ts)
}
