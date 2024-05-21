package build

import (
	"github.com/aide-cloud/moon/api"
	"github.com/aide-cloud/moon/pkg/types"
)

type PageBuild struct {
	types.Pagination
}

func NewPageBuild(page types.Pagination) *PageBuild {
	return &PageBuild{
		Pagination: page,
	}
}

// ToApi 转换为api对象
func (b *PageBuild) ToApi() *api.PaginationReply {
	if types.IsNil(b) || types.IsNil(b.Pagination) {
		return nil
	}
	return &api.PaginationReply{
		PageNum:  int32(b.GetPageNum()),
		PageSize: int32(b.GetPageSize()),
		Total:    int64(b.GetTotal()),
	}
}