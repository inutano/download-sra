package main

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jlaffaye/ftp"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading, %s complete", humanize.Bytes(wc.Total))
}

// FtpGet : Download file from FTP Server
func FtpGet(ftpURL *url.URL) {
	h := ftpURL.Host
	p := ftpURL.Path
	d := filepath.Dir(p)
	b := filepath.Base(p)
	sraFileName := strings.TrimSuffix(b, filepath.Ext(b)) + ".sra"

	out, err := os.Create(sraFileName + ".tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	c, err := ftp.Dial(h+":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	err = c.ChangeDir(d)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Retr(b)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(res, counter))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n")

	err = os.Rename(sraFileName+".tmp", sraFileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
