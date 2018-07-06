package beacon

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func eventToString(t *testing.T, ev Event) string {
	var buf bytes.Buffer

	inner, ok := ev.(*event)
	assert.True(t, ok)

	err := inner.writeKubernetesBeaconLine(&buf)
	assert.Nil(t, err)

	return buf.String()
}

func TestWriteKubernetesBeaconLine(t *testing.T) {
	ev := NewEvent(
		"THE-LABEL",
		Error,
		Count,
		7.0,
	)

	line := eventToString(t, ev)

	assert.Contains(t, line, "section.io-beacon:{", "prefix")
	assert.Contains(t, line, `"label":"THE-LABEL",`, "label")
	assert.Contains(t, line, `"severity":"error",`, "severity")
	assert.Contains(t, line, `"metric":{"type":"count","value":7}`, "metric")
	assert.Contains(t, line, "}\n", "suffix")

	ev = NewEvent(
		"other-label",
		Info,
		Timing,
		3.14,
	).
		AddContext("account", "Peterson").
		AddAnnotation("timing_unit", "ms").
		SetCorrelationID("abc123")

	line = eventToString(t, ev)
	assert.Contains(t, line, `"label":"other-label",`, "label")
	assert.Contains(t, line, `"severity":"info",`, "severity")
	assert.Contains(t, line, `"metric":{"type":"timing","value":3.14}`, "metric")
	assert.Contains(t, line, `"context":{"account":"Peterson"}`, "context")
	assert.Contains(t, line, `"annotations":{"timing_unit":"ms"}`, "annotations")
	assert.Contains(t, line, `"correlationId":"abc123"`, "correlation id")
}
