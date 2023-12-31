// Copyright 2020 EMIT Foundation.
// This file is part of E.M.I.T. .

// E.M.I.T. is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// E.M.I.T. is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with E.M.I.T. . If not, see <http://www.gnu.org/licenses/>.

package tron

import (
	bridge "github.com/emit-technology/emit-cross/bindings/tron/Bridge"
	"github.com/emit-technology/emit-cross/bindings/tron/Handler"
	log "github.com/emit-technology/emit-cross/log"
	metrics "github.com/emit-technology/emit-cross/metrics/types"
)

var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4

type writer struct {
	cfg             Config
	conn            Connection
	trc20HandlerAbi *Handler.Trc20HandlerABI
	bridgeAbi       *bridge.BridgeABI
	log             log.Logger
	stop            <-chan int
	sysErr          chan<- error // Reports fatal error to core
	metrics         *metrics.ChainMetrics
}

// NewWriter creates and returns writer
func NewWriter(conn Connection, cfg *Config, log log.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *writer {
	return &writer{
		cfg:             *cfg,
		conn:            conn,
		trc20HandlerAbi: Handler.NewTrc20HanderABI(),
		bridgeAbi:       bridge.NewBridgeABI(),
		log:             log,
		stop:            stop,
		sysErr:          sysErr,
		metrics:         m,
	}
}
