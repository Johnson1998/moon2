// Code generated by "stringer -type=DatasourceType -linecomment"; DO NOT EDIT.

package vobj

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DataSourceTypeUnknown-0]
	_ = x[DataSourceTypeMetrics-1]
	_ = x[DataSourceTypeTrace-2]
	_ = x[DataSourceTypeLog-3]
}

const _DatasourceType_name = "未知监控指标链路追踪日志"

var _DatasourceType_index = [...]uint8{0, 6, 18, 30, 36}

func (i DatasourceType) String() string {
	if i < 0 || i >= DatasourceType(len(_DatasourceType_index)-1) {
		return "DatasourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DatasourceType_name[_DatasourceType_index[i]:_DatasourceType_index[i+1]]
}

// IsDatasourcetypeunknown 是否是：未知
func (i DatasourceType) IsDatasourcetypeunknown() bool {
	return i == DataSourceTypeUnknown
}

// IsDatasourcetypemetrics 是否是：监控指标
func (i DatasourceType) IsDatasourcetypemetrics() bool {
	return i == DataSourceTypeMetrics
}

// IsDatasourcetypetrace 是否是：链路追踪
func (i DatasourceType) IsDatasourcetypetrace() bool {
	return i == DataSourceTypeTrace
}

// IsDatasourcetypelog 是否是：日志
func (i DatasourceType) IsDatasourcetypelog() bool {
	return i == DataSourceTypeLog
}

// GetValue 获取原始类型值
func (i DatasourceType) GetValue() int {
	return int(i)
}