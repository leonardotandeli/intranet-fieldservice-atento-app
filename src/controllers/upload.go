package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const maxUploadSize = 40 * 1024 * 1024 // 10 mb
const uploadPath = "./assets/uploads"

func FazerUpload(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Printf("Não foi possível receber o upload: %v\n", err)
		renderError(w, "Erro 01", http.StatusInternalServerError)
		return
	}

	// Erro 01 - Não foi possível receber os dados formulário
	// Erro 02 - Arquivo inválido
	// Erro 03 - Arquivo muito grande
	// Erro 04 - Arquivo inválido
	// Erro 05 - Tipo de Arquivo inválido
	// Erro 06 - Não foi possível ler o arquivo
	// Erro 07 - Não foi possivel salvar o arquivo

	file, fileHeader, err := r.FormFile("uploadFile")
	if err != nil {
		renderError(w, "Erro 02", http.StatusBadRequest)
		return
	}

	defer file.Close()
	fileSize := fileHeader.Size
	fmt.Printf("Tamanho do arquivo em bytes: %v\n", fileSize)

	if fileSize > maxUploadSize {
		renderError(w, "Arquivo não suportado", http.StatusBadRequest)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		renderError(w, "Erro 04", http.StatusBadRequest)
		return
	}

	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/x-rar-compressed":
	case "text/plain; charset=utf-8":
	case "application/pdf":
	case "application/zip":
		break
	default:
		renderError(w, "Erro 05", http.StatusBadRequest)
		return
	}
	fileName := randToken(12)
	fileEndings, err := mime.ExtensionsByType(detectedFileType)
	if err != nil {
		renderError(w, "Erro 06", http.StatusInternalServerError)
		return
	}
	newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
	imagemupada := fmt.Sprintf("/assets/uploads/%s\n", fileName+fileEndings[0])

	fmt.Printf("/assets/uploads/%s\n", fileName+fileEndings[0])

	newFile, err := os.Create(newPath)
	if err != nil {
		renderError(w, "Erro 07", http.StatusInternalServerError)
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		renderError(w, "Erro 07", http.StatusInternalServerError)
		return
	}
	//w.Write([]byte(fmt.Sprintf("URL: localhost:3000/assets/upload/%s\n", fileName+fileEndings[0])))
	var usuarios []modelos.Usuario

	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	busca := strings.ToLower(r.URL.Query().Get("a"))
	aSite, _ := cookie["asite"]
	fmt.Println(usuarioID)
	utils.ExecutarTemplate(w, "uploadok.html", struct {
		Usuarios  []modelos.Usuario
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Busca     string
		Pagina    string
		ASite     string
	}{
		Usuarios:  usuarios,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Nome:      nomeUsuario,
		Busca:     busca,
		Pagina:    imagemupada,
		ASite:     aSite,
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
