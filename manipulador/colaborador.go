package manipulador

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/rntpassos/consulta/repo"

	"github.com/rntpassos/consulta/modelo"
)

func RotearColaborador(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		buscaColaborador(w, r)
	case "POST":
		cadastraColaborador(w, r)
	case "DELETE":
		deletaColaborador(w, r)
	case "PATCH":
		atualizaColaborador(w, r)
	}

}

func atualizaColaborador(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	idColaborador, err := strconv.Atoi(partes[2])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao recuperar ID. Id digitado: '", partes[2], "'")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	colaborador := modelo.ColaboradorSemNil{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal(body, &colaborador)
	fmt.Printf("%+v\r\n", colaborador)
	if colaborador.Nome != "" {
		atualizaNome := "UPDATE colaborador SET nome = $2 WHERE id_colaborador = $1"
		if _, err := repo.Db.Exec(atualizaNome, idColaborador, colaborador.Nome); err != nil {
			fmt.Println("Não foi possível atualizar o colaborador ", idColaborador, "Erro: ", err.Error())
		}
	}
	if colaborador.Email != "" {
		atualizaEmail := "UPDATE colaborador SET email = $2 WHERE id_colaborador = $1"
		if _, err := repo.Db.Exec(atualizaEmail, idColaborador, colaborador.Email); err != nil {
			fmt.Println("Não foi possível atualizar o colaborador ", idColaborador, "Erro: ", err.Error())
		}
	}
	if colaborador.Telefone != "" {
		atualizaTelefone := "UPDATE colaborador SET telefone = $2 WHERE id_colaborador = $1"
		if _, err := repo.Db.Exec(atualizaTelefone, idColaborador, colaborador.Telefone); err != nil {
			fmt.Println("Não foi possível atualizar o colaborador ", idColaborador, "Erro: ", err.Error())
		}
	}
	ListarColaboradores(w, r)

}

func deletaColaborador(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	idColaborador, err := strconv.Atoi(partes[2])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao recuperar ID. Id digitado: '", partes[2], "'")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	deletaColaborador := "DELETE FROM colaborador WHERE id_colaborador = $1"

	if _, err := repo.Db.Exec(deletaColaborador, idColaborador); err != nil {
		fmt.Println("Não foi possível deletar o colaborador ", idColaborador, "Erro: ", err.Error())
	}
	ListarColaboradores(w, r)

}

func buscaColaborador(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	partes := strings.Split(r.URL.Path, "/")
	idColaborador, err := strconv.Atoi(partes[2])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao recuperar ID. Id digitado: '", partes[2], "'")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	colaboradores := modelo.Colaborador{}

	consultaColaborador := "SELECT c.id_colaborador, c.nome, c.telefone, c.email FROM colaborador c WHERE c.id_colaborador = $1"

	consulta := repo.Db.QueryRow(consultaColaborador, idColaborador)
	switch err := consulta.Scan(&colaboradores.IdColaborador, &colaboradores.Nome, &colaboradores.Telefone.String, &colaboradores.Email.String); err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Colaborador não encontrado.\r\nIdcolaborador: ", idColaborador)
	case nil:
		encoder := json.NewEncoder(w)
		encoder.Encode(colaboradores)
	default:
		panic(err)
	}
}

func ListarColaboradores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	colaboradores := make([]modelo.Colaborador, 0)

	consultaColaborador := "SELECT c.id_colaborador, c.nome, c.telefone, c.email FROM colaborador c WHERE c.registro_ativo = 'S'"

	consulta, err := repo.Db.Query(consultaColaborador)
	if err != nil {
		fmt.Println("[manipulador/colaborador] Erro ao executar query:", consultaColaborador, " Erro: ", err.Error())
		return
	}

	for consulta.Next() {
		linhaColaborador := modelo.Colaborador{}
		err := consulta.Scan(&linhaColaborador.IdColaborador, &linhaColaborador.Nome, &linhaColaborador.Telefone, &linhaColaborador.Email)
		if err != nil {
			fmt.Println("[manipulador/colaborador] Não foi possivel fazer o binding dos dados do banco na struct. Erro: ", err.Error())
		}
		colaboradores = append(colaboradores, linhaColaborador)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(colaboradores)
}

func cadastraColaborador(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	colaborador := modelo.ColaboradorSemNil{}
	colaborador.Nome = r.FormValue("nome")
	colaborador.Telefone = r.FormValue("telefone")
	colaborador.Email = r.FormValue("email")
	insereColaborador := "INSERT INTO colaborador (nome,telefone,email) VALUES ($1,$2,$3)"
	if _, err := repo.Db.Exec(insereColaborador, colaborador.Nome, colaborador.Telefone, colaborador.Email); err != nil {
		fmt.Println("[manipulador/colaborador] Erro ao executar query:", insereColaborador, " Erro: ", err.Error())
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(colaborador)
}
