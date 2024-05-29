package bizmodel

import (
	"context"
	"encoding/json"

	"github.com/aide-cloud/moon/pkg/types"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const TableNameMetricLabel = "metric_labels"

// MetricLabel mapped from table <metric_labels>
type MetricLabel struct {
	ID        uint32                  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name      string                  `gorm:"column:name;type:varchar(64);not null;comment:数据源名称" json:"name"`                                   // 数据源名称
	MetricID  uint32                  `gorm:"column:metric_id;type:int unsigned;not null;comment:所属指标" json:"metric_id"`                         // 所属指标
	CreatedAt types.Time              `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt types.Time              `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt int64                   `gorm:"column:deleted_at;type:bigint;not null;comment:删除时间" json:"deleted_at"`                             // 删除时间
	Remark    string                  `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`                                 // 备注
	Labels    []*DatasourceLabelValue `gorm:"foreignKey:LabelID" json:"labels"`
}

// String json string
func (c *MetricLabel) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func (c *MetricLabel) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *MetricLabel) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// Create func
func (c *MetricLabel) Create(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Create(c).Error
}

// Update func
func (c *MetricLabel) Update(ctx context.Context, tx *gorm.DB, conds []gen.Condition) error {
	return tx.WithContext(ctx).Model(c).Where(conds).Updates(c).Error
}

// Delete func
func (c *MetricLabel) Delete(ctx context.Context, tx *gorm.DB, conds []gen.Condition) error {
	return tx.WithContext(ctx).Where(conds).Delete(c).Error
}

// TableName MetricLabel's table name
func (*MetricLabel) TableName() string {
	return TableNameMetricLabel
}