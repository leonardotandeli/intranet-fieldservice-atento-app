package modelos

// DadosAutenticacao contém o token e o id do usuário autenticado
type DadosAutenticacao struct {
	ID                                   string `json:"id"`
	Token                                string `json:"token"`
	NOME                                 string `json:"nome,omitempty"`
	LOGIN_NT                             string `json:"login_nt,omitempty"`
	RE                                   string `json:"re,omitempty"`
	CARGO                                string `json:"cargo,omitempty"`
	EMAIL                                string `json:"email,omitempty"`
	SENHA                                string `json:"senha,omitempty"`
	V_HOMEOFFICE                         string `json:"v_homeoffice,omitempty"`
	V_HOMEOFFICE_CHAMADOS                string `json:"v_homeoffice_chamados,omitempty"`
	V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string `json:"v_homeoffice_chamados_mudar_analista,omitempty"`
	V_USUARIOS                           string `json:"v_usuarios,omitempty"`
	V_IMPRESSORAS                        string `json:"v_impressoras,omitempty"`
	V_BDC_POSTS                          string `json:"v_bdc_posts,omitempty"`
	V_BDC_ADM                            string `json:"v_bdc_adm,omitempty"`
	V_IMDB                               string `json:"v_imdb,omitempty"`
	V_GSA                                string `json:"v_gsa,omitempty"`
	V_CATRACA                            string `json:"v_catraca,omitempty"`
	V_BH                                 string `json:"v_bh,omitempty"`
	V_CEP                                string `json:"v_cep,omitempty"`
	V_MAPA_OPERACIONAL                   string `json:"v_mapa_operacional,omitempty"`
	V_MAPA_OPERACIONAL_ADM               string `json:"v_mapa_operacional_adm,omitempty"`
	Site                                 Site
}
