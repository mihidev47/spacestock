package datasource

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
)

// The FileUploader interface represents a datasources that able to upload file to storage
type FileUploader interface {
	Upload(folder string, file multipart.File, fileName string) error
	UploadBulk(params ...UploadParam) error
	GetFolderPath(folder string) string
}

// FileStorage represent structure for storage datasource that use local filesystem
type FileStorage struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// UploadParam represents structure for file upload configuration
type UploadParam struct {
	Folder   string
	File     multipart.File
	FileName string
}

var (
	// Storage-related errors
	ErrUnsupported = errors.New("file-storage: unsupported function")
)

// Init implement FileUploader on FileStorage datasource
func (f *FileStorage) Init() FileUploader {
	// @todo Validate path is exist and accessible
	// @todo Fix path string if malformed, path must end with "/"
	return f
}

// Upload store file from multipart/form request into filesystem
func (f *FileStorage) Upload(folder string, file multipart.File, fileName string) error {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err)
		return err
	}
	// Generate full path for destination file
	p := f.GetFolderPath(folder) + fileName
	// Write file
	err = ioutil.WriteFile(p, data, 0666)
	if err != nil {
		log.Error(err)
		return err
	}
	// Success
	return nil
}

// UploadBulk isn't supported in FileStorage right now
func (f *FileStorage) UploadBulk(_ ...UploadParam) error {
	return ErrUnsupported
}

// GetFolderPath returns folder path in filesystem
func (f *FileStorage) GetFolderPath(folder string) string {
	// @todo Validate folder path
	return f.Path + folder + "/"
}
