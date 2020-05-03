package manipulador

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rntpassos/consulta/repo"

	"github.com/rntpassos/consulta/modelo"
)

func RotearPaciente(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bucarPaciente(w, r)
	case "POST":
		cadastrarPaciente(w, r)
	case "DELETE":
		deletaPaciente(w, r)
	case "PATCH":
		atualizaPaciente(w, r)
	}
}

func bucarPaciente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paciente := modelo.PacienteSemNil2{}
	var err error
	paciente.DataNascimento, err = time.Parse("2006-01-02 MST", r.FormValue("data_nascimento")+" +0000")
	if err != nil {
		panic(err)
	}

	pacientes := make([]modelo.Paciente, 0)

	consultaPacientes := "SELECT p.id_paciente, p.nome, p.email, p.telefone, p.data_nascimento FROM paciente p WHERE p.registro_ativo = 'S'"

	consulta, err := repo.Db.Query(consultaPacientes)
	if err != nil {
		fmt.Println("[manipulador/paciente] Erro ao executar query:", consultaPacientes, " Erro: ", err.Error())
		return
	}

	for consulta.Next() {
		linhaPaciente := modelo.Paciente{}
		err := consulta.Scan(&linhaPaciente.IdPaciente, &linhaPaciente.Nome, &linhaPaciente.Email, &linhaPaciente.Telefone, &linhaPaciente.DataNascimento)
		if err != nil {
			fmt.Println("[manipulador/paciente] Não foi possivel fazer o binding dos dados do banco na struct. Erro: ", err.Error())
		}
		pacientes = append(pacientes, linhaPaciente)
		//		linhaPaciente.DataNascimento, _ = time.Parse("2006-01-02", linhaPaciente.DataNascimento)
		if linhaPaciente.DataNascimento == paciente.DataNascimento {
			fmt.Fprintf(w, "%+v\r\n", linhaPaciente.DataNascimento)
			fmt.Fprintf(w, "%+v\r\n", paciente.DataNascimento)
			fmt.Fprintf(w, " Data Igual\r\n")
		} else {
			fmt.Fprintf(w, "%+v\r\n", linhaPaciente.DataNascimento)
			fmt.Fprintf(w, "%+v\r\n", paciente.DataNascimento)
			fmt.Fprintf(w, " Data Diferente\r\n")
		}

	}

}

func deletaPaciente(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	idPaciente, err := strconv.Atoi(partes[2])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao recuperar ID. Id digitado: '", partes[2], "'")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	deletaPaciente := "DELETE FROM paciente WHERE id_paciente = $1"

	if _, err := repo.Db.Exec(deletaPaciente, idPaciente); err != nil {
		fmt.Println("Não foi possível deletar o paciente ", idPaciente, "Erro: ", err.Error())
	}
	ListarColaboradores(w, r)
}

func atualizaPaciente(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	idPaciente, err := strconv.Atoi(partes[2])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao recuperar ID. Id digitado: '", partes[2], "'")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	paciente := modelo.PacienteSemNil{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal(body, &paciente)
	if paciente.Nome != "" {
		if _, err := repo.Db.Exec("UPDATE paciente SET nome = $2 WHERE id_paciente = $1", idPaciente, paciente.Nome); err != nil {
			fmt.Println("Não foi possível atualizar o paciente ", idPaciente, "Erro: ", err.Error())
		}
	}
	if paciente.DataNascimento != "" {
		if _, err := repo.Db.Exec("UPDATE paciente SET data_nascimento nome = $2 WHERE id_paciente = $1", idPaciente, paciente.DataNascimento); err != nil {
			fmt.Println("Não foi possível atualizar o paciente ", idPaciente, "Erro: ", err.Error())
		}
	}
	if paciente.Telefone != "" {
		if _, err := repo.Db.Exec("UPDATE paciente SET telefone nome = $2 WHERE id_paciente = $1", idPaciente, paciente.Telefone); err != nil {
			fmt.Println("Não foi possível atualizar o paciente ", idPaciente, "Erro: ", err.Error())
		}
	}
	if paciente.Email != "" {
		if _, err := repo.Db.Exec("UPDATE paciente SET email nome = $2 WHERE id_paciente = $1", idPaciente, paciente.Email); err != nil {
			fmt.Println("Não foi possível atualizar o paciente ", idPaciente, "Erro: ", err.Error())
		}
	}
	ListarColaboradores(w, r)
}

func cadastrarPaciente(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	paciente := modelo.PacienteSemNil{
		IdPaciente:     0,
		Nome:           r.FormValue("nome"),
		DataNascimento: r.FormValue("data_nascimento"),
		Email:          r.FormValue("email"),
		Telefone:       r.FormValue("telefone")}
	inserePaciente := "INSERT INTO paciente (nome,data_nascimento,telefone,email) VALUES ($1, $2,$3,$4);"
	if _, err := repo.Db.Exec(inserePaciente, paciente.Nome, paciente.DataNascimento, paciente.Telefone, paciente.Email); err != nil {
		fmt.Println("[manipulador/paciente] Erro ao executar query:", inserePaciente, " Erro: ", err.Error())
	}
	ListarPacientes(w, r)
}

func ListarPacientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pacientes := make([]modelo.Paciente, 0)

	consultaPacientes := "SELECT p.id_paciente, p.nome, p.email, p.telefone, p.data_nascimento FROM paciente p WHERE p.registro_ativo = 'S'"

	consulta, err := repo.Db.Query(consultaPacientes)
	if err != nil {
		fmt.Println("[manipulador/paciente] Erro ao executar query:", consultaPacientes, " Erro: ", err.Error())
		return
	}

	for consulta.Next() {
		linhaPaciente := modelo.Paciente{}
		err := consulta.Scan(&linhaPaciente.IdPaciente, &linhaPaciente.Nome, &linhaPaciente.Email, &linhaPaciente.Telefone, &linhaPaciente.DataNascimento)
		if err != nil {
			fmt.Println("[manipulador/paciente] Não foi possivel fazer o binding dos dados do banco na struct. Erro: ", err.Error())
		}
		pacientes = append(pacientes, linhaPaciente)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(pacientes)
}
