package main

import(
	ConnectionDb "github.com/andamorim2206/api-poligonos/infra"
)


func main() {
	  // Inicialização do banco de dados
	  ConnectionDb.InitDB()
	  defer ConnectionDb.CloseDB()
}

