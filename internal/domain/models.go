package domain

import "mime/multipart"

type File struct {
	Name    string
	Content multipart.File
}
