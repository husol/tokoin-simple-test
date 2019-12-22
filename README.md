# USAGE

> This is a simple project which does loading JSON data from file, then helps to search data by existed fields.

- Build command:

`go build -o ./challenge /path/to/challenge.go`

- Edit `config/default.toml` with correct paths.
- Put config directory in somewhere, then specify its path in running command:

`./challenge -conf=/path/to/config/default.toml`

# CODE CHALLENGE

Using the provided data (tickets.json and users.json and organization.json)
Write a simple command line application to search the data and return the results
in a human readable format.
* Feel free to use libraries or roll your own code as you see fit.
* Where the data exists, values from any related entities MUST be included in
the results, i.e.
    * Searching organization MUST return its ticket subject and users
name
    * Searching users MUST return his/her assignee ticket subject and
submitted ticket subject and his/her organization name
    * Searching tickets MUST return its assignee name, submitter name,
and organization name.
    * The user should be able to search on any field, full value matching is fine
(e.g. "mar" won't return "mary").
* The user should also be able to search for empty values, e.g. where
description is empty.
* Search can get pretty complicated pretty easily, we just want to see that you
can code a basic search application.

# EVALUATION CRITERIA (IMPORTANT)

We will look at your project and assess it for:
1. Extensibility - separation of concerns.
2. Simplicity - aim for the simplest solution that gets the job done whilst
remaining readable, extensible and testable.
3. Test Coverage - breaking changes should break your tests. We prefer Unit
Test without testing the json file, please mock them.
4. Performance - should gracefully handle a significant increase in amount of
data provided (e.g 10000+ users).
5. Robustness - should handle and report errors. If you have any questions
about this criteria please ask.

If you have any questions about this criteria please ask.