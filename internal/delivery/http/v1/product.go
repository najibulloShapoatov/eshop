package v1

import (
	"eshop/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create product
// @Security ApiKeyAuth
// @Tags Product
// @Description create product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param input body models.Product true "product info"
// @Success 200 {integer} integer 1
// @Failure default {object} errorResponse
// @Router /api/v1/manager/products [post]
func (h *Handler) createProduct(c *gin.Context) {

	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Product.Create(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

// @Summary Get All Products 
// @Security ApiKeyAuth
// @Tags Product
// @Description get all Products
// @ID get-all-Products
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.Product
// @Failure default {object} errorResponse
// @Router /api/v1/auth/products [get]
func (h *Handler) getAllProducts(c *gin.Context) {
	items, err := h.services.Product.GetList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}



// @Summary Get Product By Id
// @Security ApiKeyAuth
// @Tags Product
// @Description get Product by id
// @ID get-product-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} domain.Product
// @Failure default {object} errorResponse
// @Router /api/v1/auth/products/{id} [get]
func (h *Handler) getProductById(c *gin.Context) {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.Product.GetByID(int64(Id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update product
// @Security ApiKeyAuth
// @Tags Product
// @Description Update product
// @ID Update-product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param input body models.Product true "product info"
// @Success 200 {string} string "product updated"
// @Failure default {object} errorResponse
// @Router /api/v1/auth/manager/products/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {

	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	err = h.services.Product.Update(int64(Id), &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "product updated")
}



// @Summary Delete product
// @Security ApiKeyAuth
// @Tags Product
// @Description Delete product
// @ID Delete-product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param input body models.Product true "product info"
// @Success 200 {string} string "product deleted"
// @Failure default {object} errorResponse
// @Router /api/v1/auth/manager/products/{id} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	err = h.services.Product.Delete(int64(Id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "product deleted")
}
