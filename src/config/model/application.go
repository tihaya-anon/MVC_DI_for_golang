package model

type app struct {
    Host string
    Port int
    Uri string
}

type database struct {
    Username string
    Password string
    Host string
    Port int
    Name string
    Uri string
}

type Application struct {
    Database database
    App app
}
