package handlers

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/Fedorova199/shorturl/internal/config"
)

func (md *Models) FileSave(idURL string, cfg config.Config) {

	file, err := os.OpenFile(cfg.FileStoragePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()

	fileModel := md.model[idURL]
	js, err := json.Marshal(fileModel)
	if err != nil {
		log.Fatalln(err)
		return
	}
	w := bufio.NewWriter(file)
	w.Write(js)
	w.WriteByte('\n')
	w.Flush()
}

func (md *Models) FileSet(idURL string, cfg config.Config) map[string]string {
	file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	model := md.model
	for scan.Scan() {
		body := []byte(scan.Text())
		err := json.Unmarshal(body, &model)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return model
}
