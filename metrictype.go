package beacon

import (
	"bytes"
	"encoding/json"
)

type MetricType int

const (
	Count MetricType = iota
	Timing
)

var metricTypeById = map[MetricType]string{
	Count:  "count",
	Timing: "timing",
}

var metricTypeByString = map[string]MetricType{
	"count":  Count,
	"timing": Timing,
}

func (m MetricType) String() string {
	return metricTypeById[m]
}

func (m MetricType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(metricTypeById[m])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (m *MetricType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*m = metricTypeByString[s]
	return nil
}
