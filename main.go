package main

import(
	"fmt"
	"net/http"
	ConnectionDb "github.com/andamorim2206/api-poligonos/infra"
)


func main() {
	  // Inicialização do banco de dados
	  ConnectionDb.InitDB()
	  defer ConnectionDb.CloseDB()
  
	  // Seu código de roteamento e manipulação aqui
	  http.HandleFunc("/", homeHandler)
  
	  // Inicie o servidor
	  fmt.Println("Servidor rodando na porta :8080")
	  http.ListenAndServe(":3306", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Bem-vindo à minha API!")
}