package modelos

import "time"

// Struct dos chamados
type Chamado struct {
	ID                  uint64    `json:"id,omitempty"`
	Nome                string    `json:"nome,omitempty"`
	Chamado             string    `json:"chamado,omitempty"`
	AtivoCPU            string    `json:"ativocpu,omitempty"`
	AtivoMonitor        string    `json:"ativomonitor,omitempty"`
	Endereco            string    `json:"endereco,omitempty"`
	Numero              string    `json:"numero,omitempty"`
	Cep                 string    `json:"cep,omitempty"`
	Senha               string    `json:"senha,omitempty"`
	Transporte          string    `json:"transporte,omitempty"`
	Acionamento         string    `json:"acionamento,omitempty"`
	Status              string    `json:"status,omitempty"`
	Bairro              string    `json:"bairro,omitempty"`
	Obs                 string    `json:"obs,omitempty"`
	Office              string    `json:"office,omitempty"`
	Ramal               string    `json:"ramal,omitempty"`
	LoginDAC            string    `json:"logindac,omitempty"`
	DataEntrada         time.Time `json:"dataentrada,omitempty"`
	DataUpdate          time.Time `json:"dataupdate,omitempty"`
	Re                  string    `json:"re,omitempty"`
	AtivoRetornoMonitor string    `json:"ativoretornomonitor,omitempty"`
	AtivoRetornoCPU     string    `json:"ativoretornocpu,omitempty"`
	PerifericoMouse     string    `json:"perifericomouse,omitempty"`
	PerifericoTeclado   string    `json:"perifericoteclado,omitempty"`
	PerifericoHead      string    `json:"perifericohead,omitempty"`
	PerifericoRede      string    `json:"perifericorede,omitempty"`
	AnalistaField       string    `json:"analistafield,omitempty"`
	GerenteOperador     string    `json:"gerenteoperador,omitempty"`
	ASite               string    `json:"asite,omitempty"`
}
