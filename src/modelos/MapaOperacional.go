package modelos

import (
	"time"
)

// MapaOperacional representa uma operação cadastrada no sistema
type MapaOperacional struct {
	IDMAPA            uint64 `json:"idmapa,omitempty"`
	OPERACAO          string `json:"operacao,omitempty" xlsx:"column(OPERACAO)"`
	VLAN_DADOS        string `json:"vlan_dados,omitempty" xlsx:"column(VLAN_DADOS)"`
	VLAN_VOZ          string `json:"vlan_voz,omitempty" xlsx:"column(VLAN_VOZ)"`
	CONFIG_CONTRATUAL string `json:"config_contratual,omitempty" xlsx:"column(CONFIG_CPU)"`
	VERSAO_WINDOWS    string `json:"versao_windows,omitempty" xlsx:"column(VERSAO_WINDOWS)"`
	IMAGEM            string `json:"imagem,omitempty" xlsx:"column(NOME_IMAGEM_DEPLOY)"`
	TEMPLATE          string `json:"template,omitempty" xlsx:"column(TEMPLATE)"`
	GRUPO_IMDB        string `json:"grupo_imdb,omitempty" xlsx:"column(GRUPO_IMDB)"`
	GRAVADOR          string `json:"gravador,omitempty" xlsx:"column(GRAVADOR)"`
	OBSERVACOES       string `json:"observacoes,omitempty" xlsx:"column(OBSERVACOES)"`
	ID_SITE           string `json:"id_site,omitempty" xlsx:"column(SITE)"`
	ID_CLIENTE        string `json:"id_cliente,omitempty" xlsx:"column(CLIENTE)"`
	ID_DOMINIO        string `json:"id_dominio,omitempty" xlsx:"column(DOMINIO)"`
	ID_DAC            string `json:"id_dac,omitempty" xlsx:"column(DAC)"`
	Pagination        Pagination
	DATA_CRIACAO      time.Time `json:"data_criacao,omitempty"`
	Site              Site
	Cliente           Cliente
	Dominio           Dominio
	Dac               Dac
}

//
type Pagination struct {
	Total        int     `json:"total,omitempty"`
	Pagina       int     `json:"pagina,omitempty"`
	UltimaPagina float64 `json:"ultima_pagina,omitempty"`
}
