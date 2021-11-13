package main

import (
        "github.com/GregNemeth/GO/magic8ball"
        "net/http"
)

func main() {
        http.Handlefunc("/", magic8ball.magic8ball)
        http.ListenAndServe("", nil)
}
