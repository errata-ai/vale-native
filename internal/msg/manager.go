package msg

import (
	"fmt"
	"log"
	"os"

	"github.com/errata-ai/nativemessaging"
	"github.com/errata-ai/vale-native/internal/cli"
)

// MsgManager is a struct that manages the communication between Vale and the
// browser.
type MsgManager struct {
	decoder nativemessaging.JSONDecoder
	encoder nativemessaging.JSONEncoder
	binMgr  *cli.ValeManager
}

// NewMsgManager creates a new MsgManager.
func NewStdioManager(path string) (*MsgManager, error) {
	bin, err := cli.NewValeManager(path)
	if err != nil {
		return nil, err
	}

	return &MsgManager{
		decoder: nativemessaging.NewNativeJSONDecoder(os.Stdin),
		encoder: nativemessaging.NewNativeJSONEncoder(os.Stdout),
		binMgr:  bin,
	}, nil
}

// Read reads a message from the browser.
func (m *MsgManager) Read() (*Message, error) {
	var message Message

	err := m.decoder.Decode(&message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// Write writes a response to the browser.
func (m *MsgManager) Write(in, out, cmd string) error {
	return m.encoder.Encode(Response{Input: in, Output: out, Command: cmd})
}

// Run runs the MsgManager.
func (m *MsgManager) Run() error {
	for {
		msg, err := m.Read()
		if err != nil {
			return err
		}

		log.Printf("Received message: %v", msg)

		switch msg.Command {
		case "lint":
			result, err := m.binMgr.Lint(msg.Text, msg.Url, msg.Format)
			if err != nil {
				return err
			}
			m.Write(msg.Text, result, "lint")
		case "version":
			result, err := m.binMgr.Version()
			if err != nil {
				return err
			}
			m.Write("", result, "version")
		case "ls-config":
			result, err := m.binMgr.Config()
			if err != nil {
				return err
			}
			m.Write("", result, "ls-config")
		default:
			msg := fmt.Sprintf("unknown command: '%s'", msg.Command)
			v, err := NewEncodedValeError(msg)
			if err != nil {
				return err
			}
			m.Write("", v, "error")
		}
	}
}
