// Package goczmq is a go interface to CZMQ
package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"errors"
)

const (
	REQ    = int(C.ZMQ_REQ)
	REP    = int(C.ZMQ_REP)
	DEALER = int(C.ZMQ_DEALER)
	ROUTER = int(C.ZMQ_ROUTER)
	PUB    = int(C.ZMQ_PUB)
	SUB    = int(C.ZMQ_SUB)
	XPUB   = int(C.ZMQ_XPUB)
	XSUB   = int(C.ZMQ_XSUB)
	PUSH   = int(C.ZMQ_PUSH)
	PULL   = int(C.ZMQ_PULL)
	PAIR   = int(C.ZMQ_PAIR)
	STREAM = int(C.ZMQ_STREAM)

	POLLIN  = int(C.ZMQ_POLLIN)
	POLLOUT = int(C.ZMQ_POLLOUT)

	ZMSG_TAG = 0x003cafe
	MORE     = int(C.ZFRAME_MORE)
	REUSE    = int(C.ZFRAME_REUSE)
	DONTWAIT = int(C.ZFRAME_DONTWAIT)

	CURVE_ALLOW_ANY = "*"
)

var (
	ErrActorCmd   = errors.New("error sending actor command")
	ErrSockAttach = errors.New("error attaching zsock")
)

func getStringType(k int) string {
	switch k {
	case REQ:
		return "REQ"
	case REP:
		return "REP"
	case DEALER:
		return "DEALER"
	case ROUTER:
		return "ROUTER"
	case PUB:
		return "PUB"
	case SUB:
		return "SUB"
	case XPUB:
		return "XPUB"
	case XSUB:
		return "XSUB"
	case PUSH:
		return "PUSH"
	case PULL:
		return "PULL"
	case PAIR:
		return "PAIR"
	case STREAM:
		return "STREAM"
	default:
		return ""
	}
}
