package main

import (
	"context"
	"time"

	"github.com/gmodx/prometheus-data-generator/log"
)

func GenerateSamples_WithoutUnix(ctx context.Context, templateName, templateFilePath, templateValueFilePath, outputPath string, resolutionSeconds int, startTime, endTime time.Time, blockHours int) error {
	log.Green("=== %v, %v -> %v, step: %vs ===", templateName, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339), resolutionSeconds)

	tVals, err := GetTFromFile[[]any](templateValueFilePath)
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	{
		// loop by day
		currentStart := startTime
		for {
			currentEnd := time.Date(currentStart.Year(), currentStart.Month(), currentStart.Day(), 0, 0, 0, 0, currentStart.Location()).AddDate(0, 0, 1)
			if currentEnd.After(endTime) {
				currentEnd = endTime
			}
			if currentStart.After(currentEnd) {
				break
			}

			// logic
			{
				log.Green("=== %v, loop, %v -> %v, step: %vs ===", templateName, currentStart.Format(time.RFC3339), currentEnd.Format(time.RFC3339), resolutionSeconds)
				helper := BuildGenHelper_WithoutUnix(currentStart.Unix(), currentEnd.Unix(), resolutionSeconds)
				err := helper.Exec(templateFilePath, outputPath, tVals)
				if err != nil {
					log.Warn(err.Error())
					return err
				}
			}

			currentStart = currentEnd
		}
	}

	return nil
}
