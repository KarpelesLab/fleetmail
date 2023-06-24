package fleetmail

import (
	"context"
	"io"

	"github.com/KarpelesLab/rest"
)

var Rest Sender = restSender{}

type restSender struct{}

func (f restSender) Send(from string, to []string, msg io.WriterTo) error {
	// perform email sending using rest + fleet
	reader, writer := io.Pipe()
	go msg.WriteTo(writer)
	_, err := rest.Upload(context.Background(), "MTA:send", "POST", map[string]any{"from": from, "to": to}, reader, "message/rfc822")
	return err
}
