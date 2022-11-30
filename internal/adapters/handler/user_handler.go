package handler

import (
	"errors"
	"net/http"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/usecase"
	"privy-backend-test/internal/exceptions"
	"privy-backend-test/internal/helpers"
	"privy-backend-test/internal/resources"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
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
	auth, err := h.uc.Login(ctx, &user)
	if err != nil {
		logrus.Error(err)
		errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
		ctx.SecureJSON(errorState.Code, errorState)
		return
	}

	ctx.SecureJSON(http.StatusOK, resources.GlobalWithDataResource(auth))
}
