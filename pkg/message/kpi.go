
package message


type KpiType int32

const (
	KpiType_CPUInfo  KpiType = 0
	KpiType_TempInfo KpiType = 1
	KpiType_RamInfo  KpiType = 2
)

type AdditionalMsg struct {
	InstanceId           int32
	InstanceName         string
	Util                 float32
}

type Kpi struct {
	ResourceId           int32
	ParentId             int32
	KpiType              KpiType
	RaisedTs             float32
	ReportedTs           float32
	ResourceType         string
	AdditionalMessage    *AdditionalMsg
}

