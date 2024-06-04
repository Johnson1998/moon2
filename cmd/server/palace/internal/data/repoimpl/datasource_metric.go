package repoimpl

import (
	"context"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/helper/model/bizmodel"
	"github.com/aide-family/moon/pkg/helper/model/bizmodel/bizquery"
	"github.com/aide-family/moon/pkg/types"

	"gorm.io/gorm/clause"
)

func NewDatasourceMetricRepository(data *data.Data) repository.DatasourceMetric {
	return &datasourceMetricRepositoryImpl{data: data}
}

type datasourceMetricRepositoryImpl struct {
	data *data.Data
}

func (l *datasourceMetricRepositoryImpl) CreateMetrics(ctx context.Context, metrics ...*bizmodel.DatasourceMetric) error {
	q, err := getBizDB(ctx, l.data)
	if !types.IsNil(err) {
		return err
	}
	return q.DatasourceMetric.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(metrics, 10)
}

func (l *datasourceMetricRepositoryImpl) CreateMetricsNoAuth(ctx context.Context, teamId uint32, metrics ...*bizmodel.DatasourceMetric) error {
	bizDB, err := l.data.GetBizGormDB(teamId)
	if !types.IsNil(err) {
		return err
	}
	q := bizquery.Use(bizDB)
	return q.DatasourceMetric.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(metrics, 10)
}
