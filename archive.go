package main

import (
	"archive/zip"
	"bytes"
	"errors"
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

func (a *Archive) Compress() error {
	if a.Buffer == nil {
		return errors.New("Buffer nil pointer error")
	}

	zw := zip.NewWriter(a.Buffer)
	for k, v := range a.targets {
		f, e := zw.Create(k)
		if e != nil {
			return e
		}
		if _, e := f.Write(v); e != nil {
			return e
		}
	}

	return zw.Close()
}
