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

func (h *CakeHandler) GetCakes(ctx *gin.Context) {
	res, err := h.uc.GetCakes(ctx)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}
	ctx.SecureJSON(http.StatusOK, resources.CakeResource(res))
}

func (h *CakeHandler) GetCakeByID(ctx *gin.Context) {
	id := ctx.Param("id")

	cakeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}
	res, err := h.uc.GetCakeByID(ctx, cakeID)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}
	ctx.SecureJSON(http.StatusOK, resources.CakeResource(res))
}

func (h *CakeHandler) Store(ctx *gin.Context) {
	var cake domain.Cake

	if err := ctx.ShouldBindJSON(&cake); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			errorState := exceptions.ErrorValidationException(http.StatusNotAcceptable, out)
			ctx.SecureJSON(errorState.Code, errorState)
			return
		}
	}

	err := h.uc.Store(ctx, &cake)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}
	ctx.SecureJSON(http.StatusOK, resources.GlobalResource("Create cake successfull"))
}

func (h *CakeHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	cakeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}

	var cake domain.Cake

	if err := ctx.ShouldBindJSON(&cake); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			errorState := exceptions.ErrorValidationException(http.StatusNotAcceptable, out)
			ctx.SecureJSON(errorState.Code, errorState)
			return
		}
	}

	cake.ID = cakeID
	err = h.uc.Update(ctx, &cake)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}
	ctx.SecureJSON(http.StatusOK, resources.GlobalResource("Update cake successfull"))
}

func (h *CakeHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	cakeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}

	err = h.uc.Delete(ctx, cakeID)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}
	ctx.SecureJSON(http.StatusOK, resources.GlobalResource("Delete cake successfull"))
}

func (h *CakeHandler) UploadImage(ctx *gin.Context) {
	filename, err := h.uc.UploadImage(ctx)
	if err != nil {
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}
	ctx.SecureJSON(http.StatusOK, resources.GlobalWithDataResource(filename))
}
