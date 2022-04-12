package modelos

import "time"

// Struct dos chamados
type Post struct {
	IDPOST       uint64    `json:"idpost,omitempty"`
	TITULO       string    `json:"titulo,omitempty"`
	CONTEUDO     string    `json:"conteudo,omitempty"`
	ID_CATEGORIA string    `json:"id_categoria,omitempty"`
	ID_USUARIO   string    `json:"id_usuario,omitempty"`
	ID_SITE      string    `json:"id_site,omitempty"`
	DATA_CRIACAO time.Time `json:"data_criacao,omitempty"`
	Usuario      Usuario
	Categoria    Post_Categoria
	Site         Site
	Cliente      Cliente
}
