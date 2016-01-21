package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

func setHeader(opt Options) string {
	header := "\\documentclass["
	if opt.PaperType == "beamer" {
		header += "17pt, dvipdfmx]{beamer}\n"
		header += "\\ifnum 42146=\\euc\"A4A2\n"
		header += "    \\AtBeginShipoutFirst{\\special{pdf:tounicode EUC-UCS2}}\n"
		header += "\\else\n"
		header += "    \\AtBeginShipoutFirst{\\special{pdf:tounicode 90ms-RKSJ-UCS2}}\n"
		header += "\\fi\n"
		return header
	} else if opt.PaperSize != "" {
		header += opt.PaperSize
	} else {
		header += "a4j"
	}

	if opt.PaperType != "" {
		header += "]{" + opt.PaperType + "}"
	} else {
		header += "]{jarticle}\n"
	}
	return header
}

func setPackages(opt Options) string {
	pkg := ""
	if opt.PaperType == "beamer" {
		pkg = "\\usepackage{graphicx}\n"
	} else {
		pkg = "\\usepackage[dvipdfmx]{graphicx}\n"
		pkg += "\\usepackage{verbatim}\n"
	}
	return pkg
}

func setBody() string {
	body := "\\begin{document}\n\n"
	body += "\\end{document}\n"
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
	Filename  string
	Author    string
	Title     string
	PaperSize string
	PaperType string
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
		} else if os.Args[i] == "-s" {
			if i+1 >= len(os.Args) {
				usage()
				return
			}
			opt.PaperSize = os.Args[i+1]
			i++
		} else if os.Args[i] == "-t" {
			if i+1 >= len(os.Args) {
				usage()
				return
			}
			opt.PaperType = os.Args[i+1]
			i++
		} else {
			opt.Filename = os.Args[i]
		}
	}

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
	writer.Write([]byte(setHeader(opt)))
	writer.Write([]byte(setMeta(opt)))
	writer.Write([]byte(setPackages(opt)))
	writer.Write([]byte(setBody()))
	writer.Flush()
}
