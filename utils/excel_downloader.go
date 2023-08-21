package utils

import (
	"crypto/tls"
	parser "github.com/elias-gill/poliapi-parser"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	listaUrls string = "/tmp/lista.html"
	pagPoli          = "https://www.pol.una.py/academico/horarios-de-clases-y-examenes"
)

func downloadFile(url string) (*io.ReadCloser, error) {
	// Disable SSL certificate verification (grande la poli que no tiene XD)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

// Funcion que orquesta la busqueda, la descarga y posterior parseo y cargado en la base
// de datos de los archivos excel disponibles en le pagina de la POlitecnica
func SearchForNewSheets() {
	list := searchSheets()
	for _, e := range list {
		// aislar el nombre
		aux := strings.Split(e, "/")
		baseName := aux[len(aux)-1]

		// si no es repetido descargar
		if isRepeatedFile(baseName) {
			continue
		}
		r, err := downloadFile(e)
		if err != nil {
			log.Println("NO se pudo descargar el archivo: ", baseName)
			continue
		}

		// TODO: guardar el archivo parseado en la DB
		_, err = parser.ParseFileFromIo(*r)
	}
}

// Busca en la web de la poli los archivos excel disponibles
// y retorna una lista con los links de descarga
func searchSheets() []string {
	// search for new excel files
	bin := "search_links.sh"
	cmd1 := exec.Command(bin)
	_, err := cmd1.Output()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(os.Getenv("CACHED_SHEETS") + "disponibles")
	if err != nil {
		panic(err)
	}

	r, _ := io.ReadAll(file)
	return strings.Split(string(r), "\n")
}

// Determina si el archivo excel ya se descargo en algun otro momento
// TODO: enlazar con la db
func isRepeatedFile(file string) bool {
	return false
}
