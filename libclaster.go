package libclaster

import (
	"errors"
	"net"
	"strings"
)

func connect(domain string) (net.Conn, error) {
	conn, err := net.Dial("tcp", domain+":5222")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Status returns the current server connection status
func Status() {

}

// NewClient connects to an XMPP server
func NewClient(jid string) (net.Conn, error) {
	if strings.Count(jid, "@") != 1 {
		err := errors.New("Invalid JID")
		return nil, err
	}
	parts := strings.Split(jid, "@")
	username, domain := parts[0], parts[1]
	username = username
	conn, err := connect(domain)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
