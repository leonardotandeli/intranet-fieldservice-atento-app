package modelos

// Struct dos Gerentes
type Site struct {
	IDSITE     uint64 `json:"idsite,omitempty"`
	NOME       string `json:"nome,omitempty"`
	ID_USUARIO string `json:"id_usuario,omitempty"`
}
