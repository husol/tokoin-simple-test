package models

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
