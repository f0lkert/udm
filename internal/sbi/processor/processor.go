package processor

import (
	"github.com/f0lkert/udm/internal/sbi/consumer"
	"github.com/f0lkert/udm/pkg/app"
)

type ProcessorUdm interface {
	app.App

	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorUdm
}

func NewProcessor(udm ProcessorUdm) (*Processor, error) {
	p := &Processor{
		ProcessorUdm: udm,
	}
	return p, nil
}
