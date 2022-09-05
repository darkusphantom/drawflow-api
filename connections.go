// Package main provides ...
package main

type nodesIn []NodeInput
type nodesOut []NodeOutput

type NodeInput struct {
  Node string `json:node`
  Input string `json:input`
}

type NodeOutput struct {
  Node string `json:node`
  Output string `json:output`
}

type ConnectionIn struct {
  Connections nodesIn `json:connectionIn`
}

type ConnectionOut struct {
  Connections nodesOut `json:connectionOut`
}

type Inputs struct {
  Input_1 ConnectionIn `json:input_1`
  Input_2 ConnectionIn `json:input_2`
}

type Outputs struct {
  Output_1 ConnectionOut `json:output_1`
  Output_2 ConnectionOut `json:output_2`
}
