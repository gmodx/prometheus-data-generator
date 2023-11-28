package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gmodx/prometheus-data-generator/log"
)

func GenerateSamples_WithoutUnix(ctx context.Context, templateName, templateFilePath, templateValueFilePath, outputPath string, days, resolutionSeconds int, endTime time.Time, blockHours int) {
	startTime := endTime.AddDate(0, 0, -days)
	log.Green("=== %v, %v -> %v, step: %vs ===", templateName, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339), resolutionSeconds)

	tVals, err := GetTFromFile[[]any](templateValueFilePath)
	if err != nil {
		log.Warn(err.Error())
		return
	}

	fmt.Println(tVals)
}
