package models

type Ticket struct {
    ID              string `json:"_id"`
    Url             string `json:"url"`
    ExternalId      string `json:"external_id"`
    CreatedAt       HusTime `json:"created_at"`
    Type            string `json:"type"`
    Subject         string `json:"subject"`
    Description     string `json:"description"`
    Priority        string `json:"priority"`
    Status          string `json:"status"`
    SubmitterId     uint `json:"submitter_id"`
    AssigneeId      uint `json:"assignee_id"`
    OrganizationId  uint `json:"organization_id"`
    Tags            []string `json:"tags"`
    HasIncidents    bool `json:"has_incidents"`
    DueAt           HusTime `json:"due_at"`
    Via             string `json:"via"`
}
