package models

type UserProfile struct {
    Id string
    Email string
    Username string
    Firstname string
    Lastname string
}

// for user nav bar require id and Username
type UserNav struct {
    Id string
    Username string
}
