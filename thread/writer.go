package thread

import (
	"bufio"
	"bytes"
	"github.com/MAYLEAF/Sommelier/json"
	"github.com/MAYLEAF/Sommelier/logger"
	"sync"
)

type writer struct {
	lock sync.Mutex
}

func (e *writer) write(thread *Handler, message []byte) error {
	thread.lock.Lock()
	defer thread.lock.Unlock()

	msg := make(map[string]interface{})
	byteReader := bytes.NewReader(message)
	bufWriter := bufio.NewWriter(thread.conn)

	if err := json.Decode(byteReader, msg); err != nil {
		logger.Error("Thread Writer Decode error occur. Err: %v, Conn: %v", err, thread.value[0])
		return err
	}

	msg["uid"] = thread.id
	last_msg := json.Read(msg)
	defer logger.Info("Request Message:" + string(last_msg))

	if _, err := bufWriter.Write(last_msg); err != nil {
		logger.Error("Thread Writer Write error occur. Err: %v, Conn: %v", err, thread.value[0])
		return err
	}
	if err := bufWriter.Flush(); err != nil {
		logger.Error("Thread Writer Flush error occur. Err: %v, Conn: %v", err, thread.value[0])
		return err
	}
	return nil
}
