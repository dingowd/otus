package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkGetUsers(b *testing.B) {
	r, _ := zip.OpenReader("testdata/users.dat.zip")
	defer r.Close()

	data, _ := r.File[0].Open()
	getUsers(data)
}
