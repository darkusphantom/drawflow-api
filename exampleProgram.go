// Package main provides ...
package main

// Types
type allPrograms []Program

// Example
var nodeInput = nodesIn {
  {
    Node: "19",
    Input: "output_1",
  },
}

var nodeInput2 = nodesIn {
  {
    Node: "19",
    Input: "output_2",
  },
}

var nodeOutput = nodesOut  {
  {
    Node: "20",
    Output: "input_1",
  },
  {
    Node: "20",
    Output: "input_2",
  },
}

var connectionsIn = ConnectionIn {
  Connections: nodeInput,
}

var connectionsIn2 = ConnectionIn {
  Connections: nodeInput2,
}

var connectionsOut = ConnectionOut {
  Connections: nodeOutput,
}

var inputs = Inputs {
  Input_1: connectionsIn,
  Input_2: connectionsIn2,
}

var outputs = Outputs {
  Output_1: connectionsOut,
}

var dataProgram = DataProgram {
  script: "Dorime",
}

var dataProgram2 = DataProgram {
  Total: 12,
  Value_1: 6,
  Value_2: 6,
  script: "Dorime2",
}

var nodes = allPrograms {
  {
    ID: 0,
    Name: "Value",
    Data: dataProgram,
    Class: "Value",
    Html: "Value",
    Typenode: "vue",
    Outputs: outputs,
    Pos_x: 25,
    Pos_y: 243,
    Input: 6,
  },
  {
    ID: 1,
    Name: "Add",
    Data: dataProgram2,
    Class: "Add",
    Html: "Add",
    Typenode: "vue",
    Inputs: inputs,
    Pos_x: 360,
    Pos_y: 198,
    Input: 6,
  },
}
