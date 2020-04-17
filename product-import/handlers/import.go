package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"product-import/importtool"
	"product-import/modelq"
	"product-import/service"
	"strconv"

	"github.com/go-chi/render"
)

type Import struct {
	sender *service.QueueSender
}

func NewImport(sender *service.QueueSender) *Import {
	return &Import{
		sender: sender,
	}
}

// @Summary Import data from csv file
// @Tags import
// @Accept  mpfd
// @Produce  text/csv
// @Param file formData file true "File with data"
// @Success 200
// @Router /v1/import [post]
func (h *Import) Import(w http.ResponseWriter, r *http.Request) {
	writer := csv.NewWriter(w)
	_ = writer.Write([]string{"total", "ok", "errors", "batch_error"})
	writer.Flush()

	mpart, err := r.MultipartReader()
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	part, err := mpart.NextPart()
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}
	defer part.Close()

	if part.FormName() != "file" {
		err := fmt.Errorf("unexpected %s form", part.FormName())
		render.Render(w, r, ErrInternal(err))
		return
	}

	const batchSize = 1024
	reader, err := importtool.NewCsvReader(part)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	okCnt := 0
	errorsCnt := 0

	currentBatch := make([]modelq.Product, 0, batchSize)
	for run := true; run; {
		entry, err := reader.Read()
		if err == io.EOF {
			run = false
		}

		if err != nil {
			logrus.WithError(err).Info("csv read error")
			errorsCnt++
		} else {
			okCnt++
			currentBatch = append(currentBatch, entry)
		}

		// send batch
		if len(currentBatch) > 0 && (len(currentBatch) == batchSize || !run) {
			err := h.sender.Send(modelq.ProductImport{
				Products: currentBatch,
			})
			currentBatch = currentBatch[:0]

			batchError := ""
			if err != nil {
				batchError = err.Error()
			}

			_ = writer.Write([]string{
				strconv.Itoa(okCnt + errorsCnt),
				strconv.Itoa(okCnt),
				strconv.Itoa(errorsCnt),
				batchError,
			})
			writer.Flush()
		}
	}
}
