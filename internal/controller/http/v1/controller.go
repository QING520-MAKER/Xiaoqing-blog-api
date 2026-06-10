package v1

import (
	"github.com/QING520-MAKER/Xiaoqing-blog-api/internal/usecase"
	authuser "github.com/QING520-MAKER/Xiaoqing-blog-api/internal/usecase/auth/user"
	"github.com/QING520-MAKER/Xiaoqing-blog-api/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type V1 struct {
	logger       logger.Interface
	validate     *validator.Validate
	captcha      usecase.Captcha
	email        usecase.Email
	auth         usecase.UserAuth
	file         usecase.File
	user         usecase.User
	content      usecase.Content
	comment      usecase.Comment
	feedback     usecase.Feedback
	link         usecase.Link
	setting      usecase.Setting
	notification usecase.Notification
	signer       authuser.TokenSigner
}

func optionalUserID(ctx fiber.Ctx) *int64 {
	claims, ok := authuser.AccessClaimsFromContext(ctx.Context())
	if !ok || claims == nil {
		return nil
	}
	uid, err := claims.UserIDInt()
	if err != nil {
		return nil
	}
	return &uid
}
