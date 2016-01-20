package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

func setHeader() string {
	return "\\documentclass[a4j]{jarticle}\n"
}

func setPackages() string {
	pkg := "\\usepackage[dvipdfmx]{graphicx}\n"
	pkg += "\\usepackage{verbatim}\n"
	return pkg
}

func setBody() string {
	body := "\\begin{document}\n\n"
	body += "\\end{document}"
	return body
}

func setMeta() string {
	meta := "\\author{"
	us, err := user.Current()
	if err != nil {
		meta += "unknown}"
	} else {
		meta += us.Username + "}"
	}
	meta += "\n\\title{no-title}\n"
	return meta

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage texgen [options] filename")
		return
	}
	wf, err := os.OpenFile(os.Args[1], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer wf.Close()
	writer := bufio.NewWriter(wf)
	writer.Write([]byte(setHeader()))
	writer.Write([]byte(setMeta()))
	writer.Write([]byte(setPackages()))
	writer.Write([]byte(setBody()))
	writer.Flush()
}
