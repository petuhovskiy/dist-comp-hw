package importtool

import (
	"encoding/csv"
	"errors"
	"io"
	"product-import/modelq"
)

var ErrMissingFields = errors.New("not found required fields in csv header")

type CsvReader struct {
	r       *csv.Reader
	mapHead map[string]int

	posUniqId         int
	posProductName    int
	posAmazonCategory int
}

func NewCsvReader(ioReader io.Reader) (*CsvReader, error) {
	r := csv.NewReader(ioReader)
	head := make(map[string]int)

	arr, err := r.Read()
	if err != nil {
		return nil, err
	}

	for i, name := range arr {
		head[name] = i
	}

	res := &CsvReader{
		r:       r,
		mapHead: head,
	}

	if !res.parseHead() {
		return res, ErrMissingFields
	}

	return res, nil
}

func (r *CsvReader) parseHead() bool {
	var ok1, ok2, ok3 bool

	r.posUniqId, ok1 = r.mapHead["uniq_id"]
	r.posProductName, ok2 = r.mapHead["product_name"]
	r.posAmazonCategory, ok3 = r.mapHead["amazon_category_and_sub_category"]

	return ok1 && ok2 && ok3
}

func (r *CsvReader) Read() (modelq.Product, error) {
	arr, err := r.r.Read()
	if err != nil {
		return modelq.Product{}, err
	}

	return modelq.Product{
		Name:     arr[r.posProductName],
		Code:     arr[r.posUniqId],
		Category: arr[r.posAmazonCategory],
	}, nil
}
