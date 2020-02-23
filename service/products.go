package service

import (
	"github.com/petuhovskiy/dist-comp-hw/modelapi"
	"github.com/petuhovskiy/dist-comp-hw/modeldb"
)

type productsRepo interface {
	Save(product modeldb.Product) (modeldb.Product, error)
	Get(id uint) (modeldb.Product, error)
	ListPage(limit, offset uint) ([]modeldb.Product, error)
	Delete(id uint) error
}

type Products struct {
	repo productsRepo
}

func NewProducts(repo productsRepo) *Products {
	return &Products{
		repo: repo,
	}
}

func (s *Products) Create(req modelapi.ProductReq) (modelapi.Product, error) {
	p := modeldb.Product{
		ID:       0,
		Name:     req.Name,
		Code:     req.Code,
		Category: req.Category,
	}

	// saving product to repository
	p, err := s.repo.Save(p)
	if err != nil {
		return modelapi.Product{}, err
	}

	return productDbToApi(p), nil
}

func (s *Products) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *Products) ListPage(limit, offset uint) ([]modelapi.Product, error) {
	list, err := s.repo.ListPage(limit, offset)
	if err != nil {
		return nil, err
	}

	var resp []modelapi.Product
	for _, p := range list {
		resp = append(resp, productDbToApi(p))
	}

	return resp, nil
}

func (s *Products) Get(id uint) (modelapi.Product, error) {
	p, err := s.repo.Get(id)
	if err != nil {
		return modelapi.Product{}, err
	}

	return productDbToApi(p), nil
}

func (s *Products) Update(id uint, req modelapi.ProductReq) (modelapi.Product, error) {
	p := modeldb.Product{
		ID:       id,
		Name:     req.Name,
		Code:     req.Code,
		Category: req.Category,
	}

	resp, err := s.repo.Save(p)
	if err != nil {
		return modelapi.Product{}, err
	}

	return productDbToApi(resp), nil
}

func productDbToApi(p modeldb.Product) modelapi.Product {
	return modelapi.Product{
		ID:       p.ID,
		Name:     p.Name,
		Code:     p.Code,
		Category: p.Category,
	}
}
