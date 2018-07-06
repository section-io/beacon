package beacon

import (
	"bytes"
	"encoding/json"
)

type Severity int

const (
	Info Severity = iota
	Error
	Warn
	Debug
)

var severityById = map[Severity]string{
	Info:  "info",
	Error: "error",
	Warn:  "Warn",
	Debug: "debug",
}

var severityByString = map[string]Severity{
	"info":  Info,
	"error": Error,
	"warn":  Warn,
	"debug": Debug,
}

func (s Severity) String() string {
	return severityById[s]
}

func (s Severity) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(severityById[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *Severity) UnmarshalJSON(b []byte) error {
	var k string
	err := json.Unmarshal(b, &k)
	if err != nil {
		return err
	}
	*s = severityByString[k]
	return nil
}
