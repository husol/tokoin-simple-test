package models

type User struct {
    ID              uint `json:"_id"`
    Url             string `json:"url"`
    ExternalId      string `json:"external_id"`
    Name            string `json:"name"`
    Alias           string `json:"alias"`
    CreatedAt       HusTime `json:"created_at"`
    Active          bool `json:"active"`
    Verified        bool `json:"verified"`
    Shared          bool `json:"shared"`
    Locale          string `json:"locale"`
    Timezone        string `json:"timezone"`
    LastLoginAt     HusTime `json:"last_login_at"`
    Email           string `json:"email"`
    Phone           string `json:"phone"`
    Signature       string `json:"signature"`
    OrganizationId    uint `json:"organization_id"`
    Tags            []string `json:"tags"`
    Suspended       bool `json:"suspended"`
    Role            string `json:"role"`
}
