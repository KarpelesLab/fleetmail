package fleetmail

import "io"

type Sender interface {
	Send(from string, to []string, msg io.WriterTo) error
}
