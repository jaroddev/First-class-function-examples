package oop

import (
	"io"
	"net"
	"sync"
)

// WRONG! this should be safe
// But you probably don't want to do it like this
// Try to avoid immediate access to shared memory with locks and mutex
type WrongMux struct {
	mutex       sync.Mutex
	connections map[net.Addr]net.Conn
}

func (wr *WrongMux) Add(conn net.Conn) {
	wr.mutex.Lock()
	defer wr.mutex.Unlock()

	wr.connections[conn.RemoteAddr()] = conn
}

func (wr *WrongMux) Remove(addr net.Addr) {
	wr.mutex.Lock()
	defer wr.mutex.Unlock()

	delete(wr.connections, addr)
}

func (wr *WrongMux) SendMessage(message string) error {
	wr.mutex.Lock()
	defer wr.mutex.Unlock()

	for _, connection := range wr.connections {
		_, err := io.WriteString(connection, message)
		if err != nil {
			return err
		}
	}

	return nil
}

// Instead share that memory by communicating
// Use go channels to do this
type Mux struct {
	add         chan net.Conn
	remove      chan net.Addr
	sendMessage chan string
}

func (mux *Mux) Add(conn net.Conn) {
	mux.add <- conn
}

func (mux *Mux) Remove(addr net.Addr) {
	mux.remove <- addr
}

// Maybe changing this function return type be a good idea
func (mux *Mux) SendMessage(message string) error {
	mux.sendMessage <- message
	return nil
}

// Here we see that find the same problem as before
// Remember the calculator and terrain examples
// We should aim for better ways to increase behaviors
func (mux *Mux) loop() {
	connections := make(map[net.Addr]net.Conn)

	for {
		select {
		case connection := <-mux.add:
			connections[connection.RemoteAddr()] = connection
		case addr := <-mux.remove:
			delete(connections, addr)
		case message := <-mux.sendMessage:
			for _, connection := range connections {
				io.WriteString(connection, message)
			}
		}

	}
}
