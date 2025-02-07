package model

type app struct {
    Host string
    Port int
}

type database struct {
    Host string
    Port int
    Name string
    Uri string
}

type Application struct {
    Database database
    App app
}
