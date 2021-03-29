package main

import (
	"fmt"
	"dropbox"

	log "github.com/joaosoft/logger"
)

func main() {
	d, err := dropbox.NewDropbox()
	if err != nil {
		panic(err)
	}

	//get user information
	log.Info("get user information")
	if user, err := d.User().Get(); err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("\n\nUSER: %+v \n\n", user)
	}

	// upload a file
	log.Info("upload a file")
	if response, err := d.File().Upload("/teste.txt", []byte("teste")); err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("\n\nUPLOADED: %+v \n\n", response)
	}

	// download the uploaded file
	log.Info("download the uploaded file")
	if response, err := d.File().Download("/teste.txt"); err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("\n\nDOWNLOADED: %s \n\n", string(response))
	}

	// create folder
	log.Info("listing folder")
	if response, err := d.Folder().Create("/bananas"); err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("\n\nCREATED FOLDER: %+v \n\n", response)
	}

	// listing folder
	log.Info("listing folder")
	if response, err := d.Folder().List("/"); err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("\n\nLIST FOLDER: %+v \n\n", response)
	}

	// deleting the uploaded file
	log.Info("deleting the uploaded file")
	if response, err := d.File().Delete("/teste.txt"); err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("\n\nDELETED FILE: %+v \n\n", response)
	}

	// deleting the created folder
	log.Info("deleting the created folder")
	if response, err := d.Folder().DeleteFolder("/bananas"); err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("\n\nDELETED FOLDER: %+v \n\n", response)
	}
}
