package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gmodx/prometheus-data-generator/log"
)

func GenerateSamples_WithUnix(ctx context.Context, templateName, templateFilePath, templateValueFilePath, outputPath string, blockHours int) error {
	valueFile := filepath.Base(templateValueFilePath)
	valueName := strings.TrimSuffix(valueFile, filepath.Ext(valueFile))

	log.Green("%v, value: %v", templateName, valueName)

	log.Green("read template value file...")
	tVals, err := GetTFromFile[[]any](templateValueFilePath)
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	helper := BuildGenHelper_WithUnix()

	outputCurrentDir := fmt.Sprintf("%v/%v_%v", outputPath, templateName, valueName)

	batchSize := 100000
	totalLen := len(tVals)
	for i := 0; i < totalLen; i += batchSize {
		end := i + batchSize
		if end > totalLen {
			end = totalLen
		}

		batch := tVals[i:end]

		progress := int(100 * i / totalLen)
		log.Green("progress: %v%%, current: %v, total: %v, batch size: %v", progress, i, totalLen, batchSize)
		err = helper.Exec(templateFilePath, outputCurrentDir, batch, blockHours)
		if err != nil {
			log.Warn(err.Error())
			return err
		}
	}

	return nil
}
