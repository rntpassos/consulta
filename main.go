package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rntpassos/consulta/manipulador"
	"github.com/rntpassos/consulta/repo"
)

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao abrir obanco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bem vindo!")
	})
	http.HandleFunc("/colaborador/", manipulador.RotearColaborador)
	http.HandleFunc("/colaborador", manipulador.ListarColaboradores)
	http.HandleFunc("/paciente", manipulador.ListarPacientes)
	http.HandleFunc("/paciente/", manipulador.RotearPaciente)
	fmt.Println("Executando serviço...")
	log.Fatal(http.ListenAndServe(":8585", nil))
}
