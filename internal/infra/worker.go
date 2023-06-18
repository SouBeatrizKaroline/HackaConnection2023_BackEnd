package infra

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/julioc98/citi/internal/app"
	"github.com/julioc98/citi/internal/domain"
)

type ShippingWorker struct {
	usecase           *app.ShippingUseCase
	sourceFolder      string
	destinationFolder string
}

func NewShippingWorker(usecase *app.ShippingUseCase, sourceFolder string) *ShippingWorker {
	return &ShippingWorker{
		usecase:           usecase,
		sourceFolder:      sourceFolder,
		destinationFolder: filepath.Join(sourceFolder, "processing"),
	}
}

func (w *ShippingWorker) Run() {
	ctx := context.Background()

	err := w.processFiles(ctx)
	if err != nil {
		log.Println("Error processing files:", err)
	}
}

func (w *ShippingWorker) parseFile(ctx context.Context, filename, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	shipping := &domain.Shipping{}

	err = shipping.FromFile(file)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	err = w.usecase.MainFlow(ctx, filename, shipping)
	if err != nil {
		return fmt.Errorf("failed to run main flow: %w", err)
	}

	return nil
}

func (w *ShippingWorker) processFiles(ctx context.Context) error {
	files, err := ioutil.ReadDir(w.sourceFolder)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()

		filePath := filepath.Join(w.sourceFolder, filename)
		destFilePath := filepath.Join(w.destinationFolder, filename)

		err := moveFile(filePath, destFilePath)
		if err != nil {
			log.Printf("Error moving file '%s': %s\n", filename, err)
			continue
		}

		log.Printf("File '%s' moved to '%s'\n", filename, destFilePath)

		err = w.parseFile(ctx, filename, destFilePath)
		if err != nil {
			log.Printf("Error parsing file '%s': %s\n", filename, err)
			continue
		}
	}

	return nil
}

func moveFile(sourcePath, destPath string) error {
	err := os.Rename(sourcePath, destPath)
	if err != nil {
		return fmt.Errorf("failed to move file: %w", err)
	}

	return nil
}
