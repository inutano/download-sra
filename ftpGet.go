package main

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/jlaffaye/ftp"
)

// FtpGet : Download file from FTP Server
func FtpGet(ftpURL *url.URL) {
	h := ftpURL.Host
	p := ftpURL.Path
	d := filepath.Dir(p)
	b := filepath.Base(p)

	out, err := os.Create(b + ".tmp")
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

	err = os.Rename(b+".tmp", b)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
