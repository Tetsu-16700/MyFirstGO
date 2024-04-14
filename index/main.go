package main

import (
	"bytes"
	"fmt"
	"io"
	"net/mail"
	"os"
	"path/filepath"

	// "prueba/zinc"
	"sync"
	"time"

	"com.githubetsu/MyFirstGO/zinc"
)

const (
	emailsDir         = "F:/enron_mail_20110402"
	cConcEmailProcess = 10
	cConcZincProcess  = 10
	cConcEmails       = 500
	cConcZincData     = 500
	batchEmail        = 50
	batchEmailMaxSize = 600000
)

var (
	concEmail  = make(chan struct{}, cConcEmailProcess)
	concZinc   = make(chan struct{}, cConcZincProcess)
	emails     = make(chan string, cConcEmails)
	zincData   = make(chan string, cConcZincData)
	wge, wgd   sync.WaitGroup
	inserted   int
	rejected   int
	formaterr  int
	emailtobig int
	mus        sync.Mutex
)

func main() {
	defer close(emails)
	defer close(zincData)

	_ = zinc.DeleteIndex()
	if err := zinc.CreateIndex(); err != nil {
		panic(err)
	}

	start := time.Now()

	wgd.Add(1)
	go processZincData()

	wgd.Add(1)
	go processEmails()

	findEmails(emailsDir)

	wge.Wait()
	close(zincData)
	wgd.Wait()

	fmt.Println("Duracion:", time.Since(start))
	fmt.Println("Inserciones:", inserted)
	fmt.Println("Errores de formatos:", formaterr)
	fmt.Println("Rechazados por tamaño de lote:", rejected)
	fmt.Println("Correos demasiado grandes:", emailtobig)
}

func findEmails(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error al acceder al directorio:", err)
			return nil
		}
		if !info.IsDir() && info.Size() <= batchEmailMaxSize {
			wge.Add(1)
			emails <- path
		}
		return nil
	})
}

func processEmails() {
	defer wgd.Done()
	for email := range emails {
		concEmail <- struct{}{}
		go func(email string) {
			defer func() {
				<-concEmail
				wge.Done()

			}()

			fileContent, err := os.ReadFile(email)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return
			}
			r := bytes.NewReader(fileContent)
			m, err := mail.ReadMessage(r)
			if err != nil {
				fmt.Println("Error al leer el mansaje del email:", err)
				mus.Lock()
				formaterr++
				mus.Unlock()
				return
			}
			body, err := io.ReadAll(m.Body)
			if err != nil {
				fmt.Println("Error al leer el contenido del email:", err)
				mus.Lock()
				formaterr++
				mus.Unlock()
				return
			}
			zincData <- fmt.Sprintf(`{"_id": "%s", "from": "%s", "to": "%s", "subject": "%s", "content": "%s"}`,
				email, m.Header.Get("From"), m.Header.Get("To"), m.Header.Get("Subject"), string(body))
		}(email)
	}
}
func processZincData() {
	defer wgd.Done()
	for data := range zincData {
		concZinc <- struct{}{}
		go func(data string) {
			defer func() {
				<-concZinc
			}()
			count, err := zinc.CreateData(data)
			if err != nil {
				fmt.Println("Error al crear datos en zinc:", err)
				return
			}
			mus.Lock()
			inserted += count
			rejected += batchEmail - count
			mus.Unlock()
		}(data)
	}
}
