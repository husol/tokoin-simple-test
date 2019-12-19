package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "husol.org/tokoin-simple-test/models"
    "io/ioutil"
    "os"
    "reflect"
    "strings"
)

const ORGANIZATION_PATH = "data/organizations.json"
const TICKET_PATH = "data/tickets.json"
const USER_PATH = "data/users.json"

var users []models.User
var tickets []models.Ticket
var organizations []models.Organization

func search()  {
    for {
        fmt.Println("Select 1) Users or 2) Tickets or 3) Organization")
        reader := bufio.NewReader(os.Stdin)
        selection, _ := reader.ReadString('\n')

        switch selection {
            case "1\n"://SEARCH USER
                fmt.Println("Enter search term")
                field, _ := reader.ReadString('\n')
                field = strings.Trim(field, "\n")
                fmt.Println("Enter search value")
                value, _ := reader.ReadString('\n')
                value = strings.Trim(value, "\n")

                var result struct{
                    User models.User
                    OrganizationName string
                    AssigneeTicketSubjects []string
                    SubmittedTicketSubjects []string
                }

                hus := models.Hus{}
                noResult := true
                fmt.Println("Searching users for "+ field + " with a value of "+ value)
                for _, user := range users {
                   if hus.GetField(user, field) == value {
                       noResult = false
                       result.User = user
                       for _, organization := range organizations {
                           if user.OrganizationId == organization.ID {
                               result.OrganizationName = organization.Name
                               break;
                           }
                       }
                       for _, ticket := range tickets {
                           if user.ID == ticket.AssigneeId {
                               result.AssigneeTicketSubjects = append(result.AssigneeTicketSubjects, ticket.Subject)
                           }
                           if user.ID == ticket.SubmitterId {
                               result.SubmittedTicketSubjects = append(result.SubmittedTicketSubjects, ticket.Subject)
                           }
                       }
                       hus.PrettyPrint(result)
                   }
                }
                if noResult {
                    fmt.Println("No results found.")
                }
                fmt.Println()
                reader.ReadString('\n')
            case "2\n"://SEARCH TICKET
                fmt.Println("Enter search term")
                field, _ := reader.ReadString('\n')
                field = strings.Trim(field, "\n")
                fmt.Println("Enter search value")
                value, _ := reader.ReadString('\n')
                value = strings.Trim(value, "\n")

                var result struct{
                    Ticket models.Ticket
                    AssigneeName string
                    SubmitterName string
                    OrganizationName string
                }

                hus := models.Hus{}
                noResult := true
                fmt.Println("Searching tickets for "+ field + " with a value of "+ value)
                for _, ticket := range tickets {
                    if hus.GetField(ticket, field) == value {
                        noResult = false
                        result.Ticket = ticket
                        for _, user := range users {
                            if ticket.AssigneeId == user.ID {
                                result.AssigneeName = user.Name
                            }
                            if ticket.SubmitterId == user.ID {
                                result.SubmitterName = user.Name
                            }
                        }
                        for _, organization := range organizations {
                            if ticket.OrganizationId == organization.ID {
                                result.OrganizationName = organization.Name
                                break
                            }
                        }
                        hus.PrettyPrint(result)
                    }
                }
                if noResult {
                    fmt.Println("No results found.")
                }
                fmt.Println()
                reader.ReadString('\n')
            case "3\n"://SEARCH ORGANIZATION
                fmt.Println("Enter search term")
                field, _ := reader.ReadString('\n')
                field = strings.Trim(field, "\n")
                fmt.Println("Enter search value")
                value, _ := reader.ReadString('\n')
                value = strings.Trim(value, "\n")

                var result struct{
                    Organization models.Organization
                    Tickets []string
                    UserNames []string
                }

                hus := models.Hus{}
                noResult := true
                fmt.Println("Searching organizations for "+ field + " with a value of "+ value)
                for _, organization := range organizations {
                    if hus.GetField(organization, field) == value {
                        noResult = false
                        result.Organization = organization
                        for _, ticket := range tickets {
                            if ticket.OrganizationId == organization.ID {
                                result.Tickets = append(result.Tickets, ticket.Subject)
                            }
                        }
                        for _, user := range users {
                            if user.OrganizationId == organization.ID {
                                result.UserNames = append(result.UserNames, user.Name)
                            }
                        }
                        hus.PrettyPrint(result)
                    }
                }
                if noResult {
                    fmt.Println("No results found.")
                }
                fmt.Println()
                reader.ReadString('\n')
            case "quit\n":
                fmt.Println("Exited.")
                os.Exit(0)
            default:
                fmt.Println("Wrong selection; Choose again.")
        }
    }

}

func list()  {
    fmt.Println("Search Users with")
    var user models.User
    e := reflect.ValueOf(&user).Elem()
    for i := 0; i < e.NumField(); i++ {
        fmt.Printf("%v\n", e.Type().Field(i).Name)
    }

    fmt.Println("\n------------------------------------")
    fmt.Println("Search Tickets with")
    var ticket models.Ticket
    e = reflect.ValueOf(&ticket).Elem()
    for i := 0; i < e.NumField(); i++ {
        fmt.Printf("%v\n", e.Type().Field(i).Name)
    }

    fmt.Println("\n------------------------------------")
    fmt.Println("Search Organizations with")
    var organization models.Organization
    e = reflect.ValueOf(&organization).Elem()
    for i := 0; i < e.NumField(); i++ {
        fmt.Printf("%v\n", e.Type().Field(i).Name)
    }
}

func main() {
    file, _ := ioutil.ReadFile(USER_PATH)
    _ = json.Unmarshal([]byte(file), &users)
    file, _ = ioutil.ReadFile(TICKET_PATH)
    _ = json.Unmarshal([]byte(file), &tickets)
    file, _ = ioutil.ReadFile(ORGANIZATION_PATH)
    _ = json.Unmarshal([]byte(file), &organizations)

    for {
        fmt.Println("Type 'quit' to exit at any time, Press Enter to continue")
        fmt.Println("\n")
        fmt.Println("\tSelect search options:")
        fmt.Println("\t* Press 1 to search")
        fmt.Println("\t* Press 2 to view a list of searchable fields")
        fmt.Println("\t* Type 'quit' to exit")
        fmt.Println("\n")

        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Choose: ")
        selection, _ := reader.ReadString('\n')

        switch selection {
        case "1\n":
            search()
            fmt.Println()
            reader.ReadString('\n')
        case "2\n":
            list()
            fmt.Println()
            reader.ReadString('\n')
        case "quit\n":
            fmt.Println("Exited.")
            os.Exit(0)
        }
    }
}
