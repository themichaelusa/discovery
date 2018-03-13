package main 

import "websocket"
import "test.go/Response"

import "bytes" 
import "encoding/binary" 
import "encoding/json"

import (
	"websocket"
	"test/Response"
	"bytes" 
	"encoding/binary"
	"encoding/json"
)

const (
	BAD_KEY = 4
	BAD_TAG = 3
	FATAL_DECODE = 2
	BAD_HEADER = 1
	CLEAN_DECODE = 0
)

type valid struct {
	alias string //tag
	outbound []byte
}

type reciever struct {
	socket websocket
	allocSize int
	incoming []Response
	validOuts []valid

	keyspace map[string]bool
	tagspace map[int]string
	masterKey string
}

func (r reciever) loop() {
	// only break on SIGINT, run teardown protocols
	for {
		r.listen()
	}
}

func (r reciever) listen() {
// retrieve packets sent downstream
}

//concurrent decode (usa bytes to reg)
func (r reciever) decode(idx int) (int status) {
	toDecode := &(r.incoming[idx].stream)

	//convert byte[] to int && check header
	buff := bytes.newBuffer(toDecode) //type Buffer *
	headerTag, err := binary.ReadVarint(buff)

	if err != nil {
		return FATAL_DECODE
	}
	//if valid header 
	if !((headerTag >> 32) & 0xAFFA) {
		return BAD_HEADER
	} 
	//verify in tagspace 
	tag := (headerTag << 32) >> 32
	alias := tagspace[tag] 
	if alias == "" {
		return BAD_TAG
	}
	//verify in keyspace 
	if !(keyspace[key]) {
		return BAD_KEY
	}

	//finish decode 
	//codedSize := (len(toDecode)/8) - 16
	for i := 16; i < len(toDecode); i++ {
		
	}

}

func (r reciever) addSender(dest string) {
	//do something with socket, specific decoded
	//capture add command
}