package models

import (
    "encoding/json"
    "fmt"
    "husol.org/vome_kids/app/models"
    "reflect"
    "strconv"
)

type Hus struct {

}

func (hus *Hus) GetField(v interface{}, field string) string {
    r := reflect.ValueOf(v)
    f := reflect.Indirect(r).FieldByName(field)

    if !f.IsValid() {
        return ""
    }
    fieldValue := f.Interface()

    switch v := fieldValue.(type) {
    case int64:
        return strconv.FormatInt(v, 10)
    case int32:
        return strconv.FormatInt(int64(v), 10)
    case int:
        return strconv.FormatInt(int64(v), 10)
    case uint:
        return strconv.FormatUint(uint64(v), 10)
    case string:
        return v
    case bool:
        if v {
            return "true"
        }
        return "false"
    case models.HusTime:
        return v.ToTime().String()
    default:
        return ""
    }
}

func (hus *Hus) PrettyPrint(v interface{}) (err error) {
    b, err := json.MarshalIndent(v, "", "  ")
    if err == nil {
        fmt.Println(string(b))
    }
    return
}