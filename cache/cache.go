package cache

import (
	"errors"
	"fmt"
)

const (
	get action = "GET"
	del action = "DEL"
	stop action = "STOP"
)

var ErrKeyNotFound = errors.New("ErrKeyNotFound")
var store cache

func init() {
	store = New()
	go store.HandleIncoming()
}

type action string

type message struct {
	action action
	data   data
	reply  chan reply
}

type reply struct {
	value any
	err   error
}

type data struct {
	key   string
	value any
}

type cache struct {
	command chan message
	datamap map[string]any
}

func New() cache {
	return cache{
		datamap: make(map[string]any),
		command: make(chan message, 100),
	}
}

func (k *cache) HandleIncoming() {
	for m := range k.command {
		switch m.action {
		case stop:
			return
		case get:
			if value, ok := k.datamap[m.data.key]; ok {
				m.reply <- reply{value: value, err: nil}
			} else {
				m.reply <- reply{value: nil, err: fmt.Errorf("%w: %s", ErrKeyNotFound, m.data.key)}
			}
		case del:
			if _, ok := k.datamap[m.data.key]; ok {
				delete(k.datamap, m.data.key)
				m.reply <- reply{value: nil, err: nil}
			} else {
				m.reply <- reply{value: nil, err: fmt.Errorf("%w: %s", ErrKeyNotFound, m.data.key)}
			}
		}
	}
}

func KillStore(){
	m := message {
		action: stop,
		reply: make(chan reply),
	}
	store.command <- m
}

func Get(key string) (any, error) {
	m := message{
		action: get,
		data: data{
			key: key,
		},
		reply: make(chan reply),
	}
	store.command <- m
	r := <- m.reply
	return r.value, r.err
}

func Del(key string) error {
	m := message{
		action: del,
		data: data{
			key: key,
		},
		reply: make(chan reply),
	}
	store.command <- m
	r := <- m.reply
	return r.err
}
