package server

type SIGNAL uint8

const (
	RUN  = 0
	EXIT = 1
)

type worcker struct {
	sig chan SIGNAL
}

func MakeWorcker() *worcker {
	return &worcker{
		sig: make(chan SIGNAL, 1),
	}
}
