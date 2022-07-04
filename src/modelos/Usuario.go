package modelos

import "time"

// Usuario representa um usuário cadastrado no sistema
type Usuario struct {
	IDUSUARIO              uint64    `json:"idusuario,omitempty"`
	NOME                   string    `json:"nome,omitempty"  xlsx:"column(NOME)"`
	LOGIN_NT               string    `json:"login_nt,omitempty"  xlsx:"column(LOGIN_NT)"`
	RE                     string    `json:"re,omitempty"  xlsx:"column(RE)"`
	CARGO                  string    `json:"cargo,omitempty"  xlsx:"column(CARGO)"`
	EMAIL                  string    `json:"email,omitempty"  xlsx:"column(EMAIL)"`
	SENHA                  string    `json:"senha,omitempty"`
	V_USUARIOS             string    `json:"v_usuarios,omitempty"`
	V_BDC_POSTS            string    `json:"v_bdc_posts,omitempty"`
	V_BDC_ADM              string    `json:"v_bdc_adm,omitempty"`
	V_IMDB                 string    `json:"v_imdb,omitempty"`
	V_GSA                  string    `json:"v_gsa,omitempty"`
	V_MAPA_OPERACIONAL     string    `json:"v_mapa_operacional,omitempty"`
	V_MAPA_OPERACIONAL_ADM string    `json:"v_mapa_operacional_adm,omitempty"`
	ID_SITE                string    `json:"id_site,omitempty"   xlsx:"column(SIGLA_SITE)"`
	DATA_CRIACAO           time.Time `json:"data_criacao,omitempty"`
	Site                   Site
	STATUS                 string `json:"status,omitempty"`
}
