package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gmodx/prometheus-data-generator/log"
)

func GenerateSamples_WithoutUnix(ctx context.Context, templateName, templateFilePath, templateValueFilePath, outputPath string, resolutionSeconds int, startTime, endTime time.Time, blockHours int) error {
	log.Green("%v, %v -> %v, step: %vs", templateName, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339), resolutionSeconds)

	tVals, err := GetTFromFile[[]any](templateValueFilePath)
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	{
		stepHours := 2.0
		currentStart := startTime
		for {
			currentDayZero := time.Date(currentStart.Year(), currentStart.Month(), currentStart.Day(), 0, 0, 0, 0, currentStart.Location())
			currentEnd := currentDayZero.Add(time.Duration(stepHours) * time.Hour * time.Duration(int(currentStart.Sub(currentDayZero).Hours()/stepHours)+1))

			if currentEnd.After(endTime) {
				currentEnd = endTime
			}
			if currentEnd.Sub(currentStart).Seconds() <= 0 {
				break
			}

			// logic
			{
				progress := int(100 * currentStart.Sub(startTime).Seconds() / endTime.Sub(startTime).Seconds())
				log.Green("progress: %v%%, %v -> %v", progress, currentStart.Format(time.RFC3339), currentEnd.Format(time.RFC3339))

				helper := BuildGenHelper_WithoutUnix(currentStart.Unix(), currentEnd.Unix(), resolutionSeconds)

				outputCurrentDir := fmt.Sprintf("%v/%v_%v_%v_%v/%v_%v", outputPath, templateName, resolutionSeconds, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339), currentStart.Format(time.RFC3339), currentEnd.Format(time.RFC3339))
				err := helper.Exec(templateFilePath, outputCurrentDir, tVals, blockHours)
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
