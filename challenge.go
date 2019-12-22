package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "husol.org/tokoin-simple-test/models"
    "io/ioutil"
    "os"
    "strings"
)

var users []models.User
var tickets []models.Ticket
var organizations []models.Organization

func loadData(s interface{}, pathFile string)  {
    file, _ := ioutil.ReadFile(pathFile)
    _ = json.Unmarshal([]byte(file), s)
}

func initData(configPath string) {
    config := models.Config{}
    config.ReadConfig(configPath)

    loadData(&users, config.Data.User)
    loadData(&tickets, config.Data.Ticket)
    loadData(&organizations, config.Data.Organization)
}

func inputData() (string, string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter search term")
    field, _ := reader.ReadString('\n')
    field = strings.Trim(field, "\n")

    fmt.Println("Enter search value")
    value, _ := reader.ReadString('\n')
    value = strings.Trim(value, "\n")

    return field, value
}

func searchUsers(field string, value string)  {
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
        sval := hus.GetField(user, field)
        for _, val := range sval {
            if val == value {
                noResult = false
                result.User = user

                for _, organization := range organizations {
                    if user.OrganizationId == organization.ID {
                        result.OrganizationName = organization.Name
                        break
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
    }
    if noResult {
        fmt.Println("No results found.")
    }
}

func searchTickets(field string, value string)  {
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
        sval := hus.GetField(ticket, field)
        for _, val := range sval {
            if val == value {
                noResult = false
                result.Ticket = ticket

                for _, user := range users {
                    if ticket.AssigneeId == user.ID{
                        result.AssigneeName = user.Name
                    }
                    if ticket.SubmitterId == user.ID{
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
    }
    if noResult {
        fmt.Println("No results found.")
    }
}

func searchOrganizations(field string, value string)  {
    var result struct {
        Organization models.Organization
        Tickets []string
        UserNames []string
    }

    hus := models.Hus{}
    noResult := true
    fmt.Println("Searching organizations for "+ field + " with a value of "+ value)
    for _, organization := range organizations {
        sval := hus.GetField(organization, field)
        for _, val := range sval {
            if val == value {
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
    }
    if noResult {
        fmt.Println("No results found.")
    }
}

func search()  {
    for {
        fmt.Println("Select 1) Users or 2) Tickets or 3) Organization")
        reader := bufio.NewReader(os.Stdin)
        selection, _ := reader.ReadString('\n')

        switch selection {
            case "1\n"://SEARCH USER
                field, value := inputData()
                searchUsers(field, value)

            case "2\n"://SEARCH TICKET
                field, value := inputData()
                searchTickets(field, value)

            case "3\n"://SEARCH ORGANIZATION
                field, value := inputData()
                searchOrganizations(field, value)

            case "quit\n":
                fmt.Println("Exited.")
                os.Exit(0)
            default:
                fmt.Println("Wrong selection.")
        }
        fmt.Println()
        reader.ReadString('\n')
    }
}

func list()  {
    hus := models.Hus{}
    fmt.Println("Search Users with")
    var user models.User
    hus.PrintFields(user)

    fmt.Println("\n------------------------------------")
    fmt.Println("Search Tickets with")
    var ticket models.Ticket
    hus.PrintFields(ticket)

    fmt.Println("\n------------------------------------")
    fmt.Println("Search Organizations with")
    var organization models.Organization
    hus.PrintFields(organization)
}

func main() {
    var configPath string
    flag.StringVar(&configPath, "conf", "config/default.toml", "path/to/config/default.toml")
    flag.Parse()

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        fmt.Println("Error: `-conf` param is missing or invalid.")
        flag.PrintDefaults()
        os.Exit(1)
    }

    initData(configPath)
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
