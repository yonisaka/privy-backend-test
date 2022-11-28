package handler

import (
	"errors"
	"net/http"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/usecase"
	"privy-backend-test/internal/exceptions"
	"privy-backend-test/internal/helpers"
	"privy-backend-test/internal/resources"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CakeHandler struct {
	uc usecase.CakeUsecase
}

func NewCakeHandler(uc usecase.CakeUsecase) *CakeHandler {
	return &CakeHandler{uc: uc}
}

func (h *CakeHandler) GetCakes(c *gin.Context) {
	res, err := h.uc.GetCakes(c)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}
	c.SecureJSON(http.StatusOK, resources.CakeResource(res))
}

func (h *CakeHandler) GetCakeByID(c *gin.Context) {
	id := c.Param("id")

	cakeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}
	res, err := h.uc.GetCakeByID(c, cakeID)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}
	c.SecureJSON(http.StatusOK, resources.CakeResource(res))
}

func (h *CakeHandler) Store(c *gin.Context) {
	var cake domain.Cake

	if err := c.ShouldBindJSON(&cake); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			errorState := exceptions.ErrorValidationException(http.StatusNotAcceptable, out)
			c.SecureJSON(errorState.Code, errorState)
			return
		}
	}

	err := h.uc.Store(c, &cake)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}
	c.SecureJSON(http.StatusOK, resources.GlobalResource("Create cake successfull"))
}

func (h *CakeHandler) Update(c *gin.Context) {
	id := c.Param("id")

	cakeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}

	var cake domain.Cake

	if err := c.ShouldBindJSON(&cake); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			errorState := exceptions.ErrorValidationException(http.StatusNotAcceptable, out)
			c.SecureJSON(errorState.Code, errorState)
			return
		}
	}

	cake.ID = cakeID
	err = h.uc.Update(c, &cake)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}
	c.SecureJSON(http.StatusOK, resources.GlobalResource("Update cake successfull"))
}

func (h *CakeHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	cakeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}

	err = h.uc.Delete(c, cakeID)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}
	c.SecureJSON(http.StatusOK, resources.GlobalResource("Delete cake successfull"))
}

func (h *CakeHandler) UploadImage(c *gin.Context) {
	filename, err := h.uc.UploadImage(c)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		c.SecureJSON(errorState.Code, errorState)
		return
	}
	c.SecureJSON(http.StatusOK, resources.GlobalWithDataResource(filename))
}
