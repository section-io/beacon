package beacon // import "github.com/section-io/beacon"

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type Event interface {
	AddContext(key, value string) Event
	AddAnnotation(key, value string) Event
	SetCorrelationID(id string) Event
	WriteToStderr() error
}

type event struct {
	Label         string            `json:"label"`
	Severity      Severity          `json:"severity"`
	Metric        metric            `json:"metric"`
	Context       map[string]string `json:"context,omitempty"`
	Annotations   map[string]string `json:"annotations,omitempty"`
	CorrelationID string            `json:"correlationId,omitempty"`
}

type metric struct {
	Type  MetricType `json:"type"`
	Value float64    `json:"value"`
}

func NewEvent(label string, severity Severity, metricType MetricType, metricValue float64) Event {
	return &event{
		Label:    label,
		Severity: severity,
		Metric: metric{
			Type:  metricType,
			Value: metricValue,
		},
	}
}

func (e *event) AddContext(key, value string) Event {
	if e.Context == nil {
		e.Context = make(map[string]string)
	}
	e.Context[key] = value
	return e
}

func (e *event) AddAnnotation(key, value string) Event {
	if e.Annotations == nil {
		e.Annotations = make(map[string]string)
	}
	e.Annotations[key] = value
	return e
}

func (e *event) SetCorrelationID(id string) Event {
	e.CorrelationID = id
	return e
}

func (e event) writeKubernetesBeaconLine(w io.Writer) error {
	// Use a Buffer to ensure we only call w.Write() once.
	buf := bytes.NewBufferString("section.io-beacon:")
	enc := json.NewEncoder(buf)

	// Encode writes ... to the stream, followed by a newline character.
	// https://golang.org/pkg/encoding/json/#Encoder.Encode
	err := enc.Encode(e)
	if err != nil {
		return err
	}

	_, err = w.Write(buf.Bytes())
	return err
}

func (e event) WriteToStderr() error {
	return e.writeKubernetesBeaconLine(os.Stderr)
}
