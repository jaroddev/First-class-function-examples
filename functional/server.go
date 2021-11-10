package functional

import (
	"errors"
	"fmt"
	"io"
	"net"
)

type Clients map[net.Addr]net.Conn
type ServerBehaviors func(Clients)

// Instead share that memory by communicating
// Use go channels to do this
type Mux struct {
	operations chan ServerBehaviors
}

func (mux *Mux) Add(connection net.Conn) {
	mux.operations <- func(connections Clients) {
		connections[connection.RemoteAddr()] = connection
	}
}

func (mux *Mux) Remove(addr net.Addr) {
	mux.operations <- func(connections Clients) {
		delete(connections, addr)
	}
}

// Lack of error handling
// Because the behavior is inside a returned function
// returned function that is called later
// We cannot directly return occuring errors
func (mux *Mux) WrongSendMessage(message string) error {
	mux.operations <- func(connections Clients) {
		for _, connection := range connections {
			io.WriteString(connection, message)
		}
	}
	return nil
}

// The solution to handle this problem is to create an error channel
// Tie it to our current function body
// Wait until behavior completion
// And returns whatever is inside
func (mux *Mux) SendMessage(message string) error {
	result := make(chan error, 1)

	mux.operations <- func(connections Clients) {
		for _, connection := range connections {
			_, err := io.WriteString(connection, message)
			if err != nil {
				result <- err
				return
			}
		}
		result <- nil
	}
	return <-result
}

// Here we see that find the same problem as before
// Remember the calculator and terrain examples
// We should aim for better ways to increase behaviors
func (mux *Mux) loop() {
	connections := make(Clients)

	for operation := range mux.operations {
		operation(connections)
	}
}

// Get back the targeted connection if it exists
// In the case the connection does not exist anymore
// return an error
func (mux *Mux) SendPrivateMessage(addr net.Addr, message string) error {
	result := make(chan net.Conn, 1)

	mux.operations <- func(connections Clients) {
		result <- connections[addr]
	}

	connection := <-result
	if connection == nil {
		return errors.New(
			fmt.Sprintf("client %v not registered", addr),
		)
	}

	_, err := io.WriteString(connection, message)
	return err
}
