package models

import (
    "encoding/json"
    "strings"
    "time"
)

type HusTime struct {
    time.Time
}

func (husTime *HusTime) UnmarshalJSON(b []byte) error {
    err := json.Unmarshal(b, &husTime.Time)
    if err != nil { //Assume non tz time input
        sTime := string(b)
        sTime = strings.TrimSuffix(sTime, "\"")
        sTime = strings.TrimPrefix(sTime, "\"")
        husTime.Time, err = time.Parse("2006-01-02T15:04:05 -07:00", sTime)
        if err != nil {
            return err //TODO add more formats to try
        }
    }
    return nil
}

func (husTime HusTime) MarshalJSON() ([]byte, error) {
    return json.Marshal(husTime.Time)
}

func (husTime HusTime) ToTime() time.Time {
    return husTime.Time
}