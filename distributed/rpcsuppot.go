package distributed

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	log.Printf("Listen on %s", host)

	if err != nil {
		return err
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go jsonrpc.ServeConn(connection)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	connection, err := net.Dial("tcp", host)

	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(connection), nil
}
