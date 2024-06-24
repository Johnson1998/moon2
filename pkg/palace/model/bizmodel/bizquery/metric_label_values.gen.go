// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package bizquery

import (
	"context"

	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newMetricLabelValue(db *gorm.DB, opts ...gen.DOOption) metricLabelValue {
	_metricLabelValue := metricLabelValue{}

	_metricLabelValue.metricLabelValueDo.UseDB(db, opts...)
	_metricLabelValue.metricLabelValueDo.UseModel(&bizmodel.MetricLabelValue{})

	tableName := _metricLabelValue.metricLabelValueDo.TableName()
	_metricLabelValue.ALL = field.NewAsterisk(tableName)
	_metricLabelValue.ID = field.NewUint32(tableName, "id")
	_metricLabelValue.Name = field.NewString(tableName, "name")
	_metricLabelValue.LabelID = field.NewUint32(tableName, "label_id")
	_metricLabelValue.CreatedAt = field.NewField(tableName, "created_at")
	_metricLabelValue.UpdatedAt = field.NewField(tableName, "updated_at")
	_metricLabelValue.DeletedAt = field.NewUint(tableName, "deleted_at")

	_metricLabelValue.fillFieldMap()

	return _metricLabelValue
}

type metricLabelValue struct {
	metricLabelValueDo

	ALL       field.Asterisk
	ID        field.Uint32
	Name      field.String
	LabelID   field.Uint32
	CreatedAt field.Field
	UpdatedAt field.Field
	DeletedAt field.Uint

	fieldMap map[string]field.Expr
}

func (m metricLabelValue) Table(newTableName string) *metricLabelValue {
	m.metricLabelValueDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m metricLabelValue) As(alias string) *metricLabelValue {
	m.metricLabelValueDo.DO = *(m.metricLabelValueDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *metricLabelValue) updateTableName(table string) *metricLabelValue {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint32(table, "id")
	m.Name = field.NewString(table, "name")
	m.LabelID = field.NewUint32(table, "label_id")
	m.CreatedAt = field.NewField(table, "created_at")
	m.UpdatedAt = field.NewField(table, "updated_at")
	m.DeletedAt = field.NewUint(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *metricLabelValue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *metricLabelValue) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["name"] = m.Name
	m.fieldMap["label_id"] = m.LabelID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
}

func (m metricLabelValue) clone(db *gorm.DB) metricLabelValue {
	m.metricLabelValueDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m metricLabelValue) replaceDB(db *gorm.DB) metricLabelValue {
	m.metricLabelValueDo.ReplaceDB(db)
	return m
}

type metricLabelValueDo struct{ gen.DO }

type IMetricLabelValueDo interface {
	gen.SubQuery
	Debug() IMetricLabelValueDo
	WithContext(ctx context.Context) IMetricLabelValueDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMetricLabelValueDo
	WriteDB() IMetricLabelValueDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMetricLabelValueDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMetricLabelValueDo
	Not(conds ...gen.Condition) IMetricLabelValueDo
	Or(conds ...gen.Condition) IMetricLabelValueDo
	Select(conds ...field.Expr) IMetricLabelValueDo
	Where(conds ...gen.Condition) IMetricLabelValueDo
	Order(conds ...field.Expr) IMetricLabelValueDo
	Distinct(cols ...field.Expr) IMetricLabelValueDo
	Omit(cols ...field.Expr) IMetricLabelValueDo
	Join(table schema.Tabler, on ...field.Expr) IMetricLabelValueDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMetricLabelValueDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMetricLabelValueDo
	Group(cols ...field.Expr) IMetricLabelValueDo
	Having(conds ...gen.Condition) IMetricLabelValueDo
	Limit(limit int) IMetricLabelValueDo
	Offset(offset int) IMetricLabelValueDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMetricLabelValueDo
	Unscoped() IMetricLabelValueDo
	Create(values ...*bizmodel.MetricLabelValue) error
	CreateInBatches(values []*bizmodel.MetricLabelValue, batchSize int) error
	Save(values ...*bizmodel.MetricLabelValue) error
	First() (*bizmodel.MetricLabelValue, error)
	Take() (*bizmodel.MetricLabelValue, error)
	Last() (*bizmodel.MetricLabelValue, error)
	Find() ([]*bizmodel.MetricLabelValue, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*bizmodel.MetricLabelValue, err error)
	FindInBatches(result *[]*bizmodel.MetricLabelValue, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*bizmodel.MetricLabelValue) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMetricLabelValueDo
	Assign(attrs ...field.AssignExpr) IMetricLabelValueDo
	Joins(fields ...field.RelationField) IMetricLabelValueDo
	Preload(fields ...field.RelationField) IMetricLabelValueDo
	FirstOrInit() (*bizmodel.MetricLabelValue, error)
	FirstOrCreate() (*bizmodel.MetricLabelValue, error)
	FindByPage(offset int, limit int) (result []*bizmodel.MetricLabelValue, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMetricLabelValueDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m metricLabelValueDo) Debug() IMetricLabelValueDo {
	return m.withDO(m.DO.Debug())
}

func (m metricLabelValueDo) WithContext(ctx context.Context) IMetricLabelValueDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m metricLabelValueDo) ReadDB() IMetricLabelValueDo {
	return m.Clauses(dbresolver.Read)
}

func (m metricLabelValueDo) WriteDB() IMetricLabelValueDo {
	return m.Clauses(dbresolver.Write)
}

func (m metricLabelValueDo) Session(config *gorm.Session) IMetricLabelValueDo {
	return m.withDO(m.DO.Session(config))
}

func (m metricLabelValueDo) Clauses(conds ...clause.Expression) IMetricLabelValueDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m metricLabelValueDo) Returning(value interface{}, columns ...string) IMetricLabelValueDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m metricLabelValueDo) Not(conds ...gen.Condition) IMetricLabelValueDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m metricLabelValueDo) Or(conds ...gen.Condition) IMetricLabelValueDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m metricLabelValueDo) Select(conds ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m metricLabelValueDo) Where(conds ...gen.Condition) IMetricLabelValueDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m metricLabelValueDo) Order(conds ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m metricLabelValueDo) Distinct(cols ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m metricLabelValueDo) Omit(cols ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m metricLabelValueDo) Join(table schema.Tabler, on ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m metricLabelValueDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m metricLabelValueDo) RightJoin(table schema.Tabler, on ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m metricLabelValueDo) Group(cols ...field.Expr) IMetricLabelValueDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m metricLabelValueDo) Having(conds ...gen.Condition) IMetricLabelValueDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m metricLabelValueDo) Limit(limit int) IMetricLabelValueDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m metricLabelValueDo) Offset(offset int) IMetricLabelValueDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m metricLabelValueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMetricLabelValueDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m metricLabelValueDo) Unscoped() IMetricLabelValueDo {
	return m.withDO(m.DO.Unscoped())
}

func (m metricLabelValueDo) Create(values ...*bizmodel.MetricLabelValue) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m metricLabelValueDo) CreateInBatches(values []*bizmodel.MetricLabelValue, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m metricLabelValueDo) Save(values ...*bizmodel.MetricLabelValue) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m metricLabelValueDo) First() (*bizmodel.MetricLabelValue, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.MetricLabelValue), nil
	}
}

func (m metricLabelValueDo) Take() (*bizmodel.MetricLabelValue, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.MetricLabelValue), nil
	}
}

func (m metricLabelValueDo) Last() (*bizmodel.MetricLabelValue, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.MetricLabelValue), nil
	}
}

func (m metricLabelValueDo) Find() ([]*bizmodel.MetricLabelValue, error) {
	result, err := m.DO.Find()
	return result.([]*bizmodel.MetricLabelValue), err
}

func (m metricLabelValueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*bizmodel.MetricLabelValue, err error) {
	buf := make([]*bizmodel.MetricLabelValue, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m metricLabelValueDo) FindInBatches(result *[]*bizmodel.MetricLabelValue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m metricLabelValueDo) Attrs(attrs ...field.AssignExpr) IMetricLabelValueDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m metricLabelValueDo) Assign(attrs ...field.AssignExpr) IMetricLabelValueDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m metricLabelValueDo) Joins(fields ...field.RelationField) IMetricLabelValueDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m metricLabelValueDo) Preload(fields ...field.RelationField) IMetricLabelValueDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m metricLabelValueDo) FirstOrInit() (*bizmodel.MetricLabelValue, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.MetricLabelValue), nil
	}
}

func (m metricLabelValueDo) FirstOrCreate() (*bizmodel.MetricLabelValue, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*bizmodel.MetricLabelValue), nil
	}
}

func (m metricLabelValueDo) FindByPage(offset int, limit int) (result []*bizmodel.MetricLabelValue, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m metricLabelValueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m metricLabelValueDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m metricLabelValueDo) Delete(models ...*bizmodel.MetricLabelValue) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *metricLabelValueDo) withDO(do gen.Dao) *metricLabelValueDo {
	m.DO = *do.(*gen.DO)
	return m
}