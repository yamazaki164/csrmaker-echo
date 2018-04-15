package main

import (
	"archive/zip"
	"bytes"
)

type Archive struct {
	targets map[string][]byte
	Buffer  *bytes.Buffer
}

func NewArchive(targets map[string][]byte) *Archive {
	archive := &Archive{
		targets: targets,
		Buffer:  new(bytes.Buffer),
	}

	return archive
}

func (a *Archive) Compress() {
	zw := zip.NewWriter(a.Buffer)
	for k, v := range a.targets {
		f, e := zw.Create(k)
		if e != nil {
			panic(e)
		}
		f.Write(v)
	}

	zw.Close()
}
