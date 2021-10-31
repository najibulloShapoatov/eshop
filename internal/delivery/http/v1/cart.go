package v1

import (
	"eshop/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get All Cart
// @Security ApiKeyAuth
// @Tags Cart
// @Description get all Cart
// @ID get-all-cart
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.Cart
// @Failure default {object} errorResponse
// @Router /api/v1/auth/manager/carts [get]
func (h *Handler) getAllCarts(c *gin.Context) {
	items, err := h.services.Cart.GetList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get Cart
// @Security ApiKeyAuth
// @Tags Cart
// @Description get Cart by id
// @ID get-cart
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Cart
// @Failure default {object} errorResponse
// @Router /api/v1/auth/user/cart [get]
func (h *Handler) getCart(c *gin.Context) {

	userId, err := GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.Cart.Get(int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Add Product to cart
// @Security ApiKeyAuth
// @Tags Cart
// @Description Add Product to cart
// @ID add-product-to-cart
// @Accept  json
// @Produce  json
// @Param input body []models.CartItem true "cart info"
// @Success 200 {object} domain.Cart
// @Failure default {object} errorResponse
// @Router /api/v1/auth/user/cart/product [post]
func (h *Handler) productToCart(c *gin.Context) {

	var input []models.CartItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	cart, err := h.services.Cart.SaveProductToCart(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
}
