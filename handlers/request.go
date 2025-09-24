package handlers

import (
	"fmt"
	"strings"
)

var gdatastore *datastore = newDataStore()

type RequestHandlers struct {
	datastore *datastore
	Command   string
}

func NewRequestHandler() *RequestHandlers {
	return &RequestHandlers{
		datastore: gdatastore,
	}
}

func (rh *RequestHandlers) Process() string {
	cmds := strings.Fields(strings.TrimSpace(rh.Command))

	if len(cmds) == 0 {
		return "invalid command"
	}

	switch strings.ToUpper(cmds[0]) {
	case "PING":
		if len(cmds) == 1 {
			return "PONG"
		} else {
			return "invalid command: PING does not take any arguments"
		}
	case "HELP":
		if len(cmds) == 1 {
			return "Commands available:\n1.GET\n2.SET\n3.KEYS\n4.GETALL"
		} else {
			return "invalid command: HELP does not take any arguments as of now"
		}
	case "GET":
		if len(cmds) != 2 {
			return "invalid command: GET syntax: `GET <key_name>`"
		}
		rh.datastore.dslock.RLock()
		defer rh.datastore.dslock.RUnlock()
		if val, exists := rh.datastore.dsdata[cmds[1]]; exists {
			return val
		}
		return fmt.Sprintf("key %s not found", cmds[1])
	case "SET":
		if len(cmds) != 3 {
			return "invalid command: SET syntax: `SET <key_name> <value>`"
		}
		rh.datastore.dslock.Lock()
		defer rh.datastore.dslock.Unlock()
		rh.datastore.dsdata[cmds[1]] = cmds[2]
		return "set success"
	case "KEYS":
		if len(cmds) != 1 {
			return "invalid command: KEYS does not take any arguments"
		}
		var resp strings.Builder
		rh.datastore.dslock.RLock()
		defer rh.datastore.dslock.RUnlock()
		for key := range rh.datastore.dsdata {
			fmt.Fprintf(&resp, "%s\n", key)
		}
		if resp.Len() == 0 {
			return "(empty)"
		}
		return resp.String()
	case "GETALL":
		if len(cmds) != 1 {
			return "invalid command: GETALL does not take any arguments"
		}
		var resp strings.Builder
		rh.datastore.dslock.RLock()
		defer rh.datastore.dslock.RUnlock()
		for key, val := range rh.datastore.dsdata {
			fmt.Fprintf(&resp, "%s - %s\n", key, val)
		}
		if resp.Len() == 0 {
			return "(empty)"
		}
		return resp.String()
	default:
		return "unknown command"
	}
}
