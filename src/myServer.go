package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "++++=====Hello world!=====20150130=="
  })
  m.Run()
}