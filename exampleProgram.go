// Package main provides ...
package main

// Types
type allPrograms []Program

// Example
var nodeInput = nodesIn {
  {
    Node: "2",
    Output: "input_1",
  },
}

var nodeOutput = nodesOut  {
  {
    Node: "2",
    Input: "out_1",
  },
}

var connectionsIn = ConnectionIn {
  Connections: nodeInput,
}

var connectionsOut = ConnectionOut {
  Connections: nodeOutput,
}

var inputs = Inputs {
  Input_1: connectionsIn,
  Input_2: connectionsIn,
}

var outputs = Outputs {
  Output_1: connectionsOut,
  Output_2: connectionsOut,
}

var nodes = allPrograms {
  {
    ID: 1,
    Name: "Value",
    Data: "",
    Class: "Value",
    Html: "Value",
    Typenode: "vue",
    Inputs: inputs,
    Outputs: outputs,
    Pos_x: 329,
    Pos_y: 257,
    Input: 5,
  },
}
