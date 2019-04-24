package libclaster

import (
	"errors"
	"net"
	"strings"
)

type Client struct {
	conn   net.Conn
	jid    string
	domain string
	port   string
}

// TODO:	Remove port hardcoding -- get from SRV record
func connect(domain string) (net.Conn, error) {
	conn, err := net.Dial("tcp", domain+":5222")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Status returns the current server connection status
func ClientStatus() {

}

// NewClient connects to an XMPP server
func NewClient(jid string, password string) (*Client, error) {
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
