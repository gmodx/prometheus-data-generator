package main

import (
	"bytes"
	"math/rand"
	"os"
	"path"
	"text/template"
	"time"

	"github.com/gmodx/prometheus-data-generator/log"
)

type GenHelper_WithoutUnix struct {
	GenBaseHelper

	From              int64
	To                int64
	ResolutionSeconds int

	Timestamps []int64
	Items      []any
}

type GenBaseHelper struct {
}

func (h GenBaseHelper) RandomInt(s, e int) int {
	return rand.Intn(e-s+1) + s
}

func (h GenBaseHelper) RandomFloat64(s, e float64) float64 {
	return rand.Float64()*(e-s) + s
}

func BuildGenHelper_WithoutUnix(from, to int64, resolutionSeconds int) GenHelper_WithoutUnix {
	helper := GenHelper_WithoutUnix{
		From:              from,
		To:                to,
		ResolutionSeconds: resolutionSeconds,
		GenBaseHelper:     GenBaseHelper{},
	}

	if resolutionSeconds == 0 {
		return helper
	}

	stepN := int64(resolutionSeconds)
	num := int(((to - from) / stepN) + 1)
	helper.Timestamps = make([]int64, 0, num)
	for t := from; t <= to; t += stepN {
		helper.Timestamps = append(helper.Timestamps, t)
	}

	return helper
}

func (helper GenHelper_WithoutUnix) Exec(templateFilePath, outputDir string, items []any, blockHours int) error {
	helper.Items = items

	log.Green("process template...")
	buf, err := ProcessTemplate(templateFilePath, outputDir, helper)
	if err != nil {
		return err
	}

	resultBytes := buf.Bytes()
	resultBytes = append(resultBytes, []byte("# EOF")...)

	log.Green("create blocks...")
	_ = os.MkdirAll(outputDir, os.ModePerm)
	err = Backfill(5000, resultBytes, outputDir, true, false, time.Duration(blockHours)*time.Hour)
	return err
}

func (helper GenHelper_WithUnix) Exec(templateFilePath, outputDir string, items []any, blockHours int) error {
	helper.Items = items

	log.Green("process template...")
	buf, err := ProcessTemplate(templateFilePath, outputDir, helper)
	if err != nil {
		return err
	}

	resultBytes := buf.Bytes()
	resultBytes = append(resultBytes, []byte("# EOF")...)

	log.Green("create blocks...")
	_ = os.MkdirAll(outputDir, os.ModePerm)
	err = Backfill(5000, resultBytes, outputDir, true, false, time.Duration(blockHours)*time.Hour)
	return err
}

func ProcessTemplate(templateFilePath, outputDir string, tplValue any) (bytes.Buffer, error) {
	name := path.Base(templateFilePath)
	temp := template.Must(template.New(name).ParseFiles(templateFilePath))

	var buf bytes.Buffer
	err := temp.Execute(&buf, tplValue)
	return buf, err
}

type GenHelper_WithUnix struct {
	GenBaseHelper

	Items []any
}

func BuildGenHelper_WithUnix() GenHelper_WithUnix {
	helper := GenHelper_WithUnix{
		GenBaseHelper: GenBaseHelper{},
	}

	return helper
}
