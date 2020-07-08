package ServicesImpl

import (
	"context"
	"micro/AppInit"
	"micro/DBModels"
	"micro/Services"
	"time"
)

type UserServices struct {
}

func (this *UserServices) UserReg(ctx context.Context, req *Services.UserModel, resp *Services.RegResponse) error {
	users := DBModels.Users{UserName: req.UserName, UserPwd: req.UserPwd, UserDate: time.Now()}
	if err := AppInit.GetDB().Create(&users).Error; err != nil {
		resp.Status = "error"
		resp.Message = "保存失败"
	} else {
		resp.Status = "success"
		resp.Message = "保存成功"
	}
	defer AppInit.GetDB().Close()
	return nil
}
