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

func setMeta(opt Options) string {
	fmt.Println(opt)
	meta := "\\author{"

	if opt.Author != "" {
		fmt.Println(opt.Author)
		meta += opt.Author + "}"
	} else {
		fmt.Println(opt.Author)
		us, err := user.Current()
		if err != nil {
			meta += "unknown}"
		} else {
			meta += us.Username + "}"
		}
	}
	meta += "\n\\title{no-title}\n"
	return meta
}

func usage() {
	fmt.Println("Usage texgen [options] filename")
}

type Options struct {
	Filename string
	Author   string
	Title    string
}

func main() {
	var opt Options
	if len(os.Args) < 2 {
		usage()
		return
	}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-u" {
			if i+1 >= len(os.Args) {
				usage()
				return
			}
			opt.Author = os.Args[i+1]
			i++
		} else {
			opt.Filename = os.Args[i]
		}
	}
	fmt.Println(opt)
	if opt.Filename == "" {
		usage()
		return
	}
	wf, err := os.OpenFile(opt.Filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer wf.Close()
	writer := bufio.NewWriter(wf)
	writer.Write([]byte(setHeader()))
	writer.Write([]byte(setMeta(opt)))
	writer.Write([]byte(setPackages()))
	writer.Write([]byte(setBody()))
	writer.Flush()
}
