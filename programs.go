// Package main provides ...
package main

type Program struct {
  ID int `json:id`
  Name string `json:name`
  Data string `json:data`
  Class string `json:class`
  Html string `json:html`
  Typenode string `json:typenode`
  Inputs Inputs "json:inputs"
  Outputs Outputs "json:outputs"
  Pos_x int `json:pos_x`
  Pos_y int `json:pos_y`
  Input int `json:input`
  script string "json:script"
}
