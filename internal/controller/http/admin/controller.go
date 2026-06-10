package admin

import (
	"github.com/QING520-MAKER/Xiaoqing-blog-api/internal/usecase"
	"github.com/QING520-MAKER/Xiaoqing-blog-api/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type Admin struct {
	logger   logger.Interface
	validate *validator.Validate
	sess     *session.Store
	auth     usecase.AdminAuth
	userAuth usecase.UserAuth
	file     usecase.File
	user     usecase.User
	content  usecase.Content
	comment  usecase.Comment
	feedback usecase.Feedback
	link     usecase.Link
	setting  usecase.Setting
	notify   usecase.Notification
}
