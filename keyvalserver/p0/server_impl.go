// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"bufio"
	"sync/atomic"
	"io"
	"strings"
)

type keyValueServer struct {
	nl       net.Listener
	conCount int32
	conns    []net.Conn

	//channel chan string
	getChan chan string
	putChan chan string

	doneChan chan bool
}

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	init_db()
	return &keyValueServer{}
}

func (kvs *keyValueServer) Start(port int) error {

	kvs.getChan = make(chan string)
	kvs.putChan = make(chan string)

	aPort := fmt.Sprintf(":%d", port)

	var err error

	kvs.nl, err = net.Listen("tcp", aPort)
	if err != nil {
		return err
	}

	for {
		conn, err := kvs.nl.Accept()
		if err != nil {
			return err
		}
		atomic.AddInt32(&kvs.conCount, 1)
		kvs.conns = append(kvs.conns, conn)
		go kvs.start(conn)
		go kvs.process(conn)

	}

	return nil
}

func (kvs *keyValueServer) Close() {
	for _, conn := range kvs.conns {
		conn.Write([]byte("Closing connections!" + "\n"))
		conn.Close()
	}
	kvs.nl.Close()
}

func (kvs *keyValueServer) Count() int {
	return int(kvs.conCount)
}

func (kvs *keyValueServer) start(conn net.Conn) error {

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err == io.EOF {
			atomic.AddInt32(&kvs.conCount, -1)
			return err
		}

		if err != nil {
			fmt.Printf("Something wrong with the input. Got %v", err)
			break
		}

		fmt.Printf("Message Received: %s", message)

		strs := strings.Split(message, ",")

		method := strs[0]
		if method == "put" {
			kvs.putChan <- message
		} else if method == "get" {
			kvs.getChan <- message
		}
	}
	return nil
}

func (kvs *keyValueServer) process(conn net.Conn) {
	var msg string
	for {
		select {
		case msg = <-kvs.putChan:
			strs := strings.Split(msg, ",")
			key := strs[1]
			value := []byte (strs[2])
			put(key, value)
		case msg = <-kvs.getChan:
			strs := strings.Split(msg, ",")
			key := strings.TrimSuffix(strs[1], "\n")

			val := string(get(key)[:])
			for _, c := range kvs.conns {
				c.Write([]byte(val+"\n"))
			}
		default:
		}
	}
}
