package v1

import (
	"eshop/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create cart
// @Security ApiKeyAuth
// @Tags Cart
// @Description create cart
// @ID create-cart
// @Accept  json
// @Produce  json
// @Param input body []models.CartItem true "cart info"
// @Success 200 {integer} integer 1
// @Failure default {object} errorResponse
// @Router /api/v1/auth/user/cart [post]
func (h *Handler) createCart(c *gin.Context) {

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

	id, err := h.services.Cart.Create(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for _, item := range input {
		item.CartId = id
		h.services.CartItem.Create(&item)
	}

	c.JSON(http.StatusOK, id)
}

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

	item, err := h.services.Cart.GetByID(int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Delete Cart
// @Security ApiKeyAuth
// @Tags Cart
// @Description Delete Cart
// @ID delete-Cart
// @Accept  json
// @Produce  json
// @Success 200 {string} string "cart deleted"
// @Failure default {object} errorResponse
// @Router /api/v1/auth/user/cart [delete]
func (h *Handler) deleteCart(c *gin.Context) {

	userId, err := GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.Cart.GetByID(int64(userId))
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
// @Router /api/v1/auth/user/cart/addproduct [post]
func (h *Handler) addProductToCart(c *gin.Context) {

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
	cart, err:= h.services.Cart.AddProductToCart(userId,input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
}

// @Summary Delete Product to cart
// @Security ApiKeyAuth
// @Tags Cart
// @Description Delete Product to cart
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param productId path int true "Product ID"
// @Param qty path int true "quantity to be removed from the cart"
// @Success 200 {object} domain.Cart
// @Failure default {object} errorResponse
// @Router /api/v1/auth/user/cart/deleteproduct/{productId}/{qty} [post]
func (h *Handler) deleteProductToCart(c *gin.Context) {

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
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	qty, err := strconv.Atoi(c.Param("qty"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid qty param")
		return
	}
	cart, err := h.services.Cart.DeleteProductFromCart(userId,int64(Id),int64(qty))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	c.JSON(http.StatusOK, cart)
}
