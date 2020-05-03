package modelo

import (
	"database/sql"
	"time"
)

type Colaborador struct {
	IdColaborador int            `json:"id_colaborador" db:"id_colaborador"`
	Nome          string         `json:"nome" db:"nome"`
	Telefone      sql.NullString `json:"telefone" db:"telefone"`
	Email         sql.NullString `json:"email" db:"email"`
}

type ColaboradorSemNil struct {
	IdColaborador int    `json:"id_colaborador" db:"id_colaborador"`
	Nome          string `json:"nome" db:"nome"`
	Telefone      string `json:"telefone" db:"telefone"`
	Email         string `json:"email" db:"email"`
}

type Agendamento struct {
	IdAgendamento  int       `json:"id_agendamento" db:"id_agendamento"`
	IdPaciente     int       `json:"id_paciente" db:"id_paciente"`
	IdProcedimento int       `json:"id_procedimento" db:"id_procedimento"`
	DataInicio     time.Time `json:"data_inicio" db:"data_inicio"`
}

type Paciente struct {
	IdPaciente     int            `json:"id_paciente" db:"id_paciente"`
	Nome           string         `json:"nome" db:"nome"`
	DataNascimento time.Time      `json:"data_nascimento" db:"data_nascimento"`
	Email          sql.NullString `json:"email" db:"email"`
	Telefone       sql.NullString `json:"telefone" db:"telefone"`
}

type PacienteSemNil struct {
	IdPaciente     int    `json:"id_paciente" db:"id_paciente"`
	Nome           string `json:"nome" db:"nome"`
	DataNascimento string `json:"data_nascimento" db:"data_nascimento"`
	Email          string `json:"email" db:"email"`
	Telefone       string `json:"telefone" db:"telefone"`
}

type PacienteSemNil2 struct {
	IdPaciente     int       `json:"id_paciente" db:"id_paciente"`
	Nome           string    `json:"nome" db:"nome"`
	DataNascimento time.Time `json:"data_nascimento" db:"data_nascimento"`
	Email          string    `json:"email" db:"email"`
	Telefone       string    `json:"telefone" db:"telefone"`
}

type Procedimento struct {
	IdProcedimento int    `json:"id_procedimento" db:"id_procedimento"`
	Denominacao    string `json:"denominacao" db:"denominacao"`
	Descricao      string `json:"descricao" db:"descricao"`
}
