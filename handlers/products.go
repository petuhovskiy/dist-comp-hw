package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"

	"github.com/petuhovskiy/dist-comp-hw/modelapi"
	"github.com/petuhovskiy/dist-comp-hw/service"
)

type Products struct {
	products *service.Products
}

func NewProductsV1(p *service.Products) *Products {
	return &Products{
		products: p,
	}
}

// @Summary Create product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body modelapi.ProductReq true "Product to create"
// @Success 200 {object} modelapi.Product
// @Router /v1/product [post]
func (h *Products) Create(w http.ResponseWriter, r *http.Request) {
	var data modelapi.ProductReq
	if err := render.Decode(r, &data); err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp, err := h.products.Create(data)
	if err != nil {
		log.Println("failed to create product, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}

// @Summary Delete product
// @Tags products
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200
// @Router /v1/product/{id} [delete]
func (h *Products) Delete(w http.ResponseWriter, r *http.Request) {
	productID := h.productID(r)

	err := h.products.Delete(productID)
	if err != nil {
		log.Println("failed to delete product, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	// maybe render ok message?
}

// @Summary Show a list of products (plus for pagination)
// @Description All products sorted in the order of decreasing id.
// @Tags products
// @Produce  json
// @Param offset query int false "Page offset"
// @Param limit query int false "Page limit"
// @Success 200 {object} modelapi.Product
// @Router /v1/product/list [get]
func (h *Products) List(w http.ResponseWriter, r *http.Request) {
	limit, offset, err := parseLimitOffset(r)
	if err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	list, err := h.products.ListPage(limit, offset)
	if err != nil {
		log.Println("failed to list products, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	// fix null result
	if list == nil {
		list = []modelapi.Product{}
	}

	render.Respond(w, r, list)
}

// @Summary Get product by id
// @Tags products
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} modelapi.Product
// @Router /v1/product/{id} [get]
func (h *Products) Get(w http.ResponseWriter, r *http.Request) {
	productID := h.productID(r)

	product, err := h.products.Get(productID)
	if err != nil {
		log.Println("failed to get product, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, product)
}

// @Summary Edit product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body modelapi.ProductReq true "Product to put"
// @Success 200 {object} modelapi.Product
// @Router /v1/product/{id} [put]
func (h *Products) Update(w http.ResponseWriter, r *http.Request) {
	productID := h.productID(r)

	var data modelapi.ProductReq
	if err := render.Decode(r, &data); err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp, err := h.products.Update(productID, data)
	if err != nil {
		log.Println("failed to update product, err=", err)
		render.Respond(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}

func (h *Products) ParseProductID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productIDStr := chi.URLParam(r, "productID")

		productID, err := parseUint(productIDStr)
		if err != nil {
			log.Println("failed to parse product id, err=", err)
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), keyProductID, productID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Products) productID(r *http.Request) uint {
	return r.Context().Value(keyProductID).(uint)
}
