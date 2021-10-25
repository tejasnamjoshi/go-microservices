package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func (p *Products) DownloadPDF(rw http.ResponseWriter, r *http.Request) {
	url := "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	rw.Header().Set("Content-Disposition", "attachment; filename=test.pdf")
	rw.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	_, err = io.Copy(rw, resp.Body)
	if err != nil {
		rw.Header().Del("Content-Disposition")
		rw.Header().Del("Content-Type")
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
