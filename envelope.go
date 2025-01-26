package models

import (
	"go.opentelemetry.io/otel/propagation"
)

type Envelope struct {
	Carrier propagation.MapCarrier `json:"Carrier"`
}
