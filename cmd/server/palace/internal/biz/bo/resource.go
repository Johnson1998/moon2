package bo

import (
	"github.com/aide-family/moon/pkg/helper/model"
	"github.com/aide-family/moon/pkg/types"
)

type (
	QueryResourceListParams struct {
		Keyword string `json:"keyword"`
		Page    types.Pagination
	}

	ResourceSelectOptionBuild struct {
		*model.SysAPI
	}

	QueryTeamMenuListParams struct {
		TeamID uint32 `json:"teamID"`
	}
)

// NewResourceSelectOptionBuild 构建资源选项构建器
func NewResourceSelectOptionBuild(resource *model.SysAPI) *ResourceSelectOptionBuild {
	return &ResourceSelectOptionBuild{
		SysAPI: resource,
	}
}

// ToSelectOption 转换为选项
func (b *ResourceSelectOptionBuild) ToSelectOption() *SelectOptionBo {
	if types.IsNil(b) || types.IsNil(b.SysAPI) {
		return nil
	}
	return &SelectOptionBo{
		Value:    b.ID,
		Label:    b.Name,
		Disabled: !b.Status.IsEnable(),
	}
}
