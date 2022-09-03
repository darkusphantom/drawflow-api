// Package main provides ...
package main

import (
  "fmt" // Principal
  "net/http" // Para el manejo de HTTP
  "log" /// Para arrancar el servidor y mostrarlo
  "encoding/json" // Para encoding los archivos json
  "io/ioutil" //Para manejar las entradas y salidas de los datos del servidor
  "strconv" // Convierte el string en un entero

  "github.com/gorilla/mux"
)

type node struct {
  ID int `json:ID`
  Name string `json:Name`
  Data string `json:Data`
}

type allNodes []node

var nodes = allNodes {
  {
    ID: 1,
    Name: "Value",
    Data: "",
  },
}

func getNodes(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "application/json"); //Indica a la tarea que es un formato json
  json.NewEncoder(w).Encode(nodes)
}

func getNode(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r) //Mux vars trae las variables de la peticion que le envies
  nodeID, err := strconv.Atoi(vars["id"]) //Atoi convierte una cadena en entero

  if err != nil {
    fmt.Fprintf(w, "ID invalido")
    return
  }

  for _, node := range nodes {
    if node.ID == nodeID {
      w.Header().Set("Content-type", "application/json"); //Indica a la tarea que es un formato json
      json.NewEncoder(w).Encode(nodes)
    }
  }
}

func createNode(w http.ResponseWriter, r *http.Request) {
  var newNode node
  reqBody, err := ioutil.ReadAll(r.Body) //Guarda los datos del request del body

  if err != nil {
    fmt.Fprintf(w, "Insert a valid node")
  }

  json.Unmarshal(reqBody, &newNode)

  newNode.ID = len(nodes) + 1
  nodes = append(nodes, newNode)

  w.Header().Set("Content-type", "application/json");
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(newNode)
}

func deleteNode(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  nodeID, err := strconv.Atoi(vars["id"])

  if err != nil {
    fmt.Fprintf(w, "Invalid ID")
    return
  }

  for i, node := range nodes {
    if node.ID == nodeID {
      nodes = append(nodes[:i], nodes[i+1:]...)
      fmt.Fprintf(w, "The node with ID %v has been remove succesuflly", nodeID)
    }
  }
}

func updateNode(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  nodeID, err := strconv.Atoi(vars["id"])
  var updateNode node

  if err != nil {
    fmt.Fprintf(w, "Invalid ID")
  }

  reqBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    fmt.Fprintf(w, "Please, enter valid data")
  }

  json.Unmarshal(reqBody, &updateNode)

  for i, node := range nodes {
    if node.ID == nodeID {
      nodes = append(nodes[:i], nodes[i+1:]...) //Elimina la tarea...
      updateNode.ID = nodeID
      nodes = append(nodes, updateNode)

      fmt.Fprintf(w, "The node with ID %v has been updated", nodeID)
    }
  }
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to mmy epyai")
}


func main() {
  //Se utiliza para indicar la url. Si usas StrictSlash, se encarga de escribir la url correcta
  router := mux.NewRouter().StrictSlash(true)

  // Cuando el usuario visita /, se ejecuta la accion
  router.HandleFunc("/", indexRoute)
  router.HandleFunc("/nodes", getNodes).Methods("GET")
  router.HandleFunc("/nodes/{id}", getNode).Methods("GET")
  router.HandleFunc("/nodes", createNode).Methods("POST")
  router.HandleFunc("/nodes/{id}", updateNode).Methods("PUT")
  router.HandleFunc("/nodes/{id}", deleteNode).Methods("DELETE")
  
  //http.ListenAndServe se encarga de abrir un puerto y la data a utilzar
  //log.Fatal se encarga de mostrar la información
  log.Fatal(http.ListenAndServe(":3000", router))
}
