package models

import "reflect"

type Organization struct {
    ID              uint `json:"_id"`
    Url             string `json:"url"`
    ExternalId      string `json:"external_id"`
    Name            string `json:"name"`
    DomainNames     []string `json:"domain_names"`
    CreatedAt       HusTime `json:"created_at"`
    Details         string `json:"details"`
    SharedTickets   bool `json:"shared_tickets"`
    Tags            []string `json:"tags"`
}

func (organization *Organization) GetField(field string) string {
    r := reflect.ValueOf(organization)
    f := reflect.Indirect(r).FieldByName(field)
    return string(f.String())
}