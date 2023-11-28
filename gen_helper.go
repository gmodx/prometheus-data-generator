package main

import (
	"bytes"
	"html/template"
	"math/rand"
	"path"

	"github.com/Masterminds/sprig/v3"
)

type GenHelper_WithoutUnix struct {
	GenBaseHelper

	From              int64
	To                int64
	ResolutionSeconds int

	Timestamps []int64

	Data any
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

func (helper GenHelper_WithoutUnix) Exec(templateFilePath, outputDir string, data any) error {
	helper.Data = data

	buf, err := helper.ProcessTemplate(templateFilePath, outputDir, helper)
	if err != nil {
		return err
	}

	resultBytes := buf.Bytes()
	resultBytes = append(resultBytes, []byte("# EOF")...)

	// fmt.Println(resultBytes)

	return nil
}

func (helper GenHelper_WithoutUnix) ProcessTemplate(templateFilePath, outputDir string, data any) (bytes.Buffer, error) {
	name := path.Base(templateFilePath)
	temp := template.Must(template.New(name).Funcs(sprig.TxtFuncMap()).ParseFiles(templateFilePath))

	var buf bytes.Buffer
	err := temp.Execute(&buf, helper)
	return buf, err
}
