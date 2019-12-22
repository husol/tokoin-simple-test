package models

import (
    "encoding/json"
    "fmt"
    "husol.org/vome_kids/app/models"
    "reflect"
    "strconv"
    "strings"
)

type Hus struct {

}

func (hus *Hus) PrintFields(obj interface{}) {
    val := reflect.ValueOf(obj)
    for i := 0; i < val.Type().NumField(); i++ {
        fmt.Println(val.Type().Field(i).Tag.Get("json"))
    }
}

func (hus *Hus) GetField(s interface{}, tag string) []string {
    //Get field name from json tag
    rt := reflect.TypeOf(s)
    if rt.Kind() != reflect.Struct {
        panic("Not struct")
    }
    name := ""
    for i := 0; i < rt.NumField(); i++ {
        f := rt.Field(i)
        v := strings.Split(f.Tag.Get("json"), ",")[0]
        if v == tag {
            name = f.Name
            break
        }
    }

    var result []string

    //Get value from field name
    rv := reflect.ValueOf(s)
    f := reflect.Indirect(rv).FieldByName(name)
    if !f.IsValid() {
        return result
    }
    fieldValue := f.Interface()

    switch v := fieldValue.(type) {
    case int64:
        return append(result, strconv.FormatInt(v, 10))
    case int32:
        return append(result, strconv.FormatInt(int64(v), 10))
    case int:
        return append(result, strconv.FormatInt(int64(v), 10))
    case uint:
        return append(result, strconv.FormatUint(uint64(v), 10))
    case string:
        return append(result, v)
    case []string:
        return v
    case bool:
        if v {
            return append(result, "true")
        }
        return append(result, "false")
    case models.HusTime:
        return append(result, v.ToTime().String())
    default:
        return result
    }
}

func (hus *Hus) PrettyPrint(v interface{}) (err error) {
    b, err := json.MarshalIndent(v, "", "  ")
    if err == nil {
        fmt.Println(string(b))
    }
    return
}