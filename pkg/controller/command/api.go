/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package command

import (
	"io"

	"github.com/hyperledger/aries-framework-go/pkg/didcomm/dispatcher"
)

// Exec is controller command execution function type
type Exec func(rw io.Writer, req io.Reader) Error

// MessageHandler maintains registered message services
// and it allows dynamic registration of message services
type MessageHandler interface {
	// Services returns list of available message services in this message handler
	Services() []dispatcher.MessageService
	// Register registers given message services to this message handler
	Register(msgSvcs ...dispatcher.MessageService) error
	// Unregister unregisters message service with given name from this message handler
	Unregister(name string) error
}
