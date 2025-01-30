package server

import (
	"fmt"
	"net/http"
	"os"
)

type optErr struct {
	option string
	err    error
}

func (o *optErr) Error() string {
	return fmt.Sprintf("Unable to apply option (%s): %s", o.option, o.err.Error())
}

type Option func(o *opts) error

type opts struct {
	path    string
	address string
}

func WithCWD() Option {
	return func(o *opts) error {
		cwd, err := os.Getwd()
		if err != nil {
			return &optErr{option: "WithCWD", err: err}
		}

		o.path = cwd

		return nil
	}
}

func WithAddress(address string) Option {
	return func(o *opts) error {
		o.address = address

		return nil
	}
}

func New(ofs ...Option) (*http.Server, error) {
	opts := &opts{}

	for _, of := range append([]Option{
		WithCWD(),
		WithAddress("127.0.0.1:8000"),
	}, ofs...) {
		if err := of(opts); err != nil {
			return nil, err
		}
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(opts.path))

	mux.Handle("/", fs)

	return &http.Server{
		Handler: mux,
		Addr:    opts.address,
	}, nil
}
