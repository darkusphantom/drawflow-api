// Package main provides ...
package main

type allDataProgram []allPrograms

type DataProgram struct {
  Total int `json:id`
  Value_1 int `json:value_1`
  Value_2 int `json:value_2`
  condition bool `json:condition`
  script string `json:script`
}

type Program struct {
  ID int `json:id`
  Name string `json:name`
  Data DataProgram `json:data`
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
