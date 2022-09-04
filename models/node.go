// Package main provides ...
package node

type Output struct {
  Node string `json:node`
  Output string `json:output`
}

type Outputs struct {
  Output_1 string `json:output_1`
  Output_2 string `json:output_2`
}

type Program struct {
  ID int `json:id`
  Name string `json:name`
  Data string `json:data`
  Class string `json:class`
  Html string `json:html`
  Typenode string `json:typenode`
  Inputs string "json:inputs"
  Outputs Outputs "json:outputs"
  Pos_x int `json:pos_x`
  Pos_y int `json:pos_y`
  Input int `json:input`
}


