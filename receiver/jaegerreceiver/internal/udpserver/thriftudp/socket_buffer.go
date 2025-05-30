// Copyright The OpenTelemetry Authors
// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package thriftudp // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver/thriftudp"

import (
	"fmt"
	"net"
	"syscall"
)

func setSocketBuffer(conn *net.UDPConn, bufferSize int) error {
	rawConn, err := conn.SyscallConn()
	if err != nil {
		return fmt.Errorf("failed to get raw connection: %w", err)
	}

	var syscallErr error
	controlErr := rawConn.Control(func(fd uintptr) {
		syscallErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, bufferSize)
	})
	if controlErr != nil {
		return fmt.Errorf("rawconn control failed: %w", controlErr)
	}
	if syscallErr != nil {
		return fmt.Errorf("syscall failed: %w", syscallErr)
	}

	return nil
}
