package infra

import (
	"context"
	"os"
	"path/filepath"

	"github.com/julioc98/citi/internal/domain"
)

type Storage struct {
	sourceFolder string
}

func NewReturnStorage(sourceFolder string) *Storage {
	return &Storage{
		sourceFolder: sourceFolder,
	}
}

func (r *Storage) Save(ctx context.Context, filename string, ret domain.Return) error {

	file, err := ret.ToFile()
	if err != nil {
		return err
	}

	newFile, err := os.Create(filepath.Join(r.sourceFolder, "return", "RETURN_"+filename))
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = newFile.Write(file)
	if err != nil {
		return err
	}

	return nil
}
