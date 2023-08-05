package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rahularcota/dropbox/orch"
	"mime/multipart"
	"net/http"
)

const (
	userIdHeader = "userId"
)

type FileController struct {
	FileOrch orch.IFileOrch
}

func NewFileController(fileUploadOrch orch.IFileOrch) *FileController {
	return &FileController{FileOrch: fileUploadOrch}
}

func (controller *FileController) UploadFile(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get(userIdHeader)
	if userId == "" {
		http.Error(w, "Empty userId", http.StatusBadRequest)
		return
	}
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(file)

	err = controller.FileOrch.UploadFile(file, handler.Filename, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (controller *FileController) ListFiles(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)[userIdHeader]
	if userId == "" {
		http.Error(w, "Empty userId", http.StatusBadRequest)
		return
	}

	filesInfo, err := controller.FileOrch.GetFilesInfoForUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(filesInfo)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
