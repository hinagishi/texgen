package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

type Options struct {
	Filename  string
	BaseName  string
	Author    string
	Title     string
	PaperSize string
	PaperType string
	Inst      string
	InstAbbr  string
}

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

func setBody(opt Options) string {
	body := "\\begin{document}\n\n"
	if opt.PaperType == "beamer" {
		body += "\\frame{\\titlepage}\n\n"
	} else {
		body += "\\bibliographystyle{plain}\n"
		body += "\\bibliography{" + opt.BaseName + "}\n"
	}
	body += "\\end{document}\n"
	return body
}

func createBibtex(opt Options) {
	wf, err := os.OpenFile(opt.BaseName+".bib", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer wf.Close()
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
	if opt.PaperType == "beamer" && opt.Inst != "" {
		meta += "\n\\institute{" + opt.Inst + "}\n"
	}
	meta += "\n\\title{no-title}\n"
	return meta
}

func setTheme(opt Options) string {
	theme := "\\usetheme{Madrid}\n"
	theme += "\\usefonttheme{professionalfonts}\n"
	theme += "\\useinnertheme{circles}\n"
	return theme
}

func usage() {
	fmt.Println("Usage texgen [options] filename")
}

func help() {
	fmt.Println("texgen is a tool to create a latex template file\n")
	usage()
	fmt.Println("")
	fmt.Println("The commands are:")
	fmt.Println("\t-t\tspecify paper type (ex:article, jarticle, beamer)")
	fmt.Println("\t-u\tspecify paper author")
	fmt.Println("\t-i\tspecify institute")
	fmt.Println("\t-s\tspecify paper size (ex:a4, a4j)")
	fmt.Println("\t-h\tshow help")
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
		} else if os.Args[i] == "-i" {
			if i+1 >= len(os.Args) {
				usage()
				return
			}
			opt.Inst = os.Args[i+1]
			i++
		} else if os.Args[i] == "-t" {
			if i+1 >= len(os.Args) {
				usage()
				return
			}
			opt.PaperType = os.Args[i+1]
			i++
		} else if os.Args[i] == "-h" {
			help()
			return
		} else {
			opt.Filename = os.Args[i]
		}
	}

	if opt.Filename == "" {
		usage()
		return
	}

	if opt.PaperType != "beamer" {
		elm := strings.Split(opt.Filename, ".")
		for i := 0; i < len(elm)-1; i++ {
			if i != 0 {
				opt.BaseName += "."
			}
			opt.BaseName += elm[i]
		}
		createBibtex(opt)
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
	if opt.PaperType == "beamer" {
		writer.Write([]byte(setTheme(opt)))
	}
	writer.Write([]byte(setBody(opt)))
	writer.Flush()
}
