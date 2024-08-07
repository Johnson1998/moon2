package user

import (
	"context"

	"github.com/aide-family/moon/api/admin"
	userapi "github.com/aide-family/moon/api/admin/user"
	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

// Service 用户管理服务
type Service struct {
	userapi.UnimplementedUserServer

	userBiz *biz.UserBiz
}

// NewUserService 创建用户服务
func NewUserService(userBiz *biz.UserBiz) *Service {
	return &Service{
		userBiz: userBiz,
	}
}

// CreateUser 创建用户 只允许管理员操作
func (s *Service) CreateUser(ctx context.Context, req *userapi.CreateUserRequest) (*userapi.CreateUserReply, error) {
	pass := types.NewPassword(req.GetPassword())
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	createParams := build.NewBuilder().WithCreateUserBo(req).ToCreateUserBO(claims.GetUser(), pass)
	_, err := s.userBiz.CreateUser(ctx, createParams)
	if !types.IsNil(err) {
		return nil, err
	}
	return &userapi.CreateUserReply{}, nil
}

// UpdateUser 更新用户基础信息， 只允许管理员操作
func (s *Service) UpdateUser(ctx context.Context, req *userapi.UpdateUserRequest) (*userapi.UpdateUserReply, error) {
	updateParams := build.NewBuilder().WithUpdateUserBo(req).ToUpdateUserBO()
	if err := s.userBiz.UpdateUser(ctx, updateParams); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.UpdateUserReply{}, nil
}

// DeleteUser 删除用户 只允许管理员操作
func (s *Service) DeleteUser(ctx context.Context, req *userapi.DeleteUserRequest) (*userapi.DeleteUserReply, error) {
	if err := s.userBiz.DeleteUser(ctx, req.GetId()); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.DeleteUserReply{}, nil
}

// GetUser 获取用户详情
func (s *Service) GetUser(ctx context.Context, req *userapi.GetUserRequest) (*userapi.GetUserReply, error) {
	userDo, err := s.userBiz.GetUser(ctx, req.GetId())
	if !types.IsNil(err) {
		return nil, err
	}
	return &userapi.GetUserReply{
		User: build.NewBuilder().WithAPIUserBo(userDo).ToAPI(),
	}, nil
}

// ListUser 获取用户列表
func (s *Service) ListUser(ctx context.Context, req *userapi.ListUserRequest) (*userapi.ListUserReply, error) {
	queryParams := &bo.QueryUserListParams{
		Keyword: req.GetKeyword(),
		Page:    types.NewPagination(req.GetPagination()),
		Status:  vobj.Status(req.GetStatus()),
		Gender:  vobj.Gender(req.GetGender()),
		Role:    vobj.Role(req.GetRole()),
	}
	userDos, err := s.userBiz.ListUser(ctx, queryParams)
	if !types.IsNil(err) {
		return nil, err
	}
	return &userapi.ListUserReply{
		List: types.SliceTo(userDos, func(user *model.SysUser) *admin.User {
			return build.NewBuilder().WithAPIUserBo(user).ToAPI()
		}),
		Pagination: build.NewPageBuilder(queryParams.Page).ToAPI(),
	}, nil
}

// BatchUpdateUserStatus 批量更新用户状态
func (s *Service) BatchUpdateUserStatus(ctx context.Context, req *userapi.BatchUpdateUserStatusRequest) (*userapi.BatchUpdateUserStatusReply, error) {
	params := &bo.BatchUpdateUserStatusParams{
		Status: vobj.Status(req.GetStatus()),
		IDs:    req.GetIds(),
	}
	if err := s.userBiz.BatchUpdateUserStatus(ctx, params); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.BatchUpdateUserStatusReply{}, nil
}

// ResetUserPassword 重置用户密码
func (s *Service) ResetUserPassword(ctx context.Context, req *userapi.ResetUserPasswordRequest) (*userapi.ResetUserPasswordReply, error) {
	// TODO 发送邮件等相关操作
	return &userapi.ResetUserPasswordReply{}, nil
}

// ResetUserPasswordBySelf 重置用户密码
func (s *Service) ResetUserPasswordBySelf(ctx context.Context, req *userapi.ResetUserPasswordBySelfRequest) (*userapi.ResetUserPasswordBySelfReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	// 查询用户详情
	userDo, err := s.userBiz.GetUser(ctx, claims.GetUser())
	if !types.IsNil(err) {
		return nil, err
	}
	newPass := types.NewPassword(req.GetNewPassword(), userDo.Salt)
	oldPass := userDo.Password
	// 对比旧密码正确
	if oldPass != req.OldPassword {
		return nil, merr.ErrorI18nPasswordErr(ctx)
	}

	// 对比两次密码相同, 相同修改无意义
	if newPass.String() == oldPass {
		return nil, merr.ErrorI18nPasswordSameErr(ctx)
	}

	params := &bo.ResetUserPasswordBySelfParams{
		UserID: claims.GetUser(),
		// 使用新的盐
		Password: types.NewPassword(req.GetNewPassword()),
	}
	if err = s.userBiz.ResetUserPasswordBySelf(ctx, params); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.ResetUserPasswordBySelfReply{}, nil
}

// GetUserSelectList 获取用户下拉列表
func (s *Service) GetUserSelectList(ctx context.Context, req *userapi.GetUserSelectListRequest) (*userapi.GetUserSelectListReply, error) {
	params := &bo.QueryUserSelectParams{
		Keyword: req.GetKeyword(),
		Page:    types.NewPagination(req.GetPagination()),
		Status:  vobj.Status(req.GetStatus()),
		Gender:  vobj.Gender(req.GetGender()),
		Role:    vobj.Role(req.GetRole()),
		IDs:     req.GetIds(),
	}
	userSelectOptions, err := s.userBiz.GetUserSelectList(ctx, params)
	if !types.IsNil(err) {
		return nil, err
	}
	return &userapi.GetUserSelectListReply{
		List: types.SliceTo(userSelectOptions, func(option *bo.SelectOptionBo) *admin.SelectItem {
			return build.NewSelectBuilder(option).ToAPI()
		}),
		Pagination: build.NewPageBuilder(params.Page).ToAPI(),
	}, nil
}

// UpdateUserPhone 更新用户手机号
func (s *Service) UpdateUserPhone(ctx context.Context, req *userapi.UpdateUserPhoneRequest) (*userapi.UpdateUserPhoneReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	// TODO 验证手机号短信验证码
	params := &bo.UpdateUserPhoneRequest{
		UserID: claims.GetUser(),
		Phone:  req.GetPhone(),
	}
	if err := s.userBiz.UpdateUserPhone(ctx, params); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.UpdateUserPhoneReply{}, nil
}

// UpdateUserEmail 更新用户邮箱
func (s *Service) UpdateUserEmail(ctx context.Context, req *userapi.UpdateUserEmailRequest) (*userapi.UpdateUserEmailReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	// TODO 验证邮箱验证码
	params := &bo.UpdateUserEmailRequest{
		UserID: claims.GetUser(),
		Email:  req.GetEmail(),
	}
	if err := s.userBiz.UpdateUserEmail(ctx, params); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.UpdateUserEmailReply{}, nil
}

// UpdateUserAvatar 更新用户头像
func (s *Service) UpdateUserAvatar(ctx context.Context, req *userapi.UpdateUserAvatarRequest) (*userapi.UpdateUserAvatarReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	params := &bo.UpdateUserAvatarRequest{
		UserID: claims.GetUser(),
		Avatar: req.GetAvatar(),
	}
	if err := s.userBiz.UpdateUserAvatar(ctx, params); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.UpdateUserAvatarReply{}, nil
}

// UpdateUserBaseInfo 更新用户基础信息
func (s *Service) UpdateUserBaseInfo(ctx context.Context, req *userapi.UpdateUserBaseInfoRequest) (*userapi.UpdateUserBaseInfoReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	updateParams := &bo.UpdateUserBaseParams{
		ID:       claims.GetUser(),
		Gender:   vobj.Gender(req.GetGender()),
		Remark:   req.GetRemark(),
		Nickname: req.GetNickname(),
	}
	if err := s.userBiz.UpdateUserBaseInfo(ctx, updateParams); !types.IsNil(err) {
		return nil, err
	}
	return &userapi.UpdateUserBaseInfoReply{}, nil
}
