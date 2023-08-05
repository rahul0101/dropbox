package main

import (
	"github.com/gorilla/mux"
	config2 "github.com/rahularcota/dropbox/config"
	"github.com/rahularcota/dropbox/controller"
	"github.com/rahularcota/dropbox/dal/repo/impl"
	"github.com/rahularcota/dropbox/database"
	impl3 "github.com/rahularcota/dropbox/orch/impl"
	impl2 "github.com/rahularcota/dropbox/service/impl"
	"github.com/rahularcota/dropbox/utils"
	"log"
	"net/http"
)

func main() {

	config := config2.LoadConf()
	dbConn, err := database.NewPostgresConnection(&config.DB)
	if err != nil {
		panic(err)
	}

	// Storing files in memory in different folders for each user
	// Can create another implementation of IFileRepo -> S3FileRepo and substitute here to upload files to S3 cloud storage
	fileRepo := impl.NewInMemoryFileRepo()
	// Using postgres to store file metadata
	fileMetaDataRepo := impl.NewPGFileMetaDataRepo(dbConn)

	fileService := impl2.NewFileService(fileRepo)
	fileMetaDataService := impl2.NewFileMetaDataService(fileMetaDataRepo)

	fileUploadOrch := impl3.NewFileOrch(fileMetaDataService, fileService)

	fileController := controller.NewFileController(fileUploadOrch)

	r := mux.NewRouter()
	r.HandleFunc("/file/upload", fileController.UploadFile).Methods(http.MethodPost)
	r.HandleFunc("/file/list/{userId}", fileController.ListFiles).Methods(http.MethodGet)

	http.Handle("/", utils.EnableCors(r))
	err = http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
