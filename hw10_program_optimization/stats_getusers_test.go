package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkGetDomains(b *testing.B) {
	r, _ := zip.OpenReader("testdata/users.dat.zip")
	defer r.Close()

	data, _ := r.File[0].Open()
	getResult(data, "com")
}
