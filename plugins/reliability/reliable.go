package reliability

import (
	"bytes"
	"sync"
	"time"
)

// Message Sending
type ReliableSender struct {
	seq  uint64
	unacked map[uint64][]byte
	sendFunc func([]byte) error
	mu sync.Mutex
	timeout time.Duration
}

func NewReliableSender(sendFunc func([]byte) error, timeout time.Duration)
*ReliableSender {
	return &ReliableSender{
		unacked: make(map[uint64][]byte),
		sendFunc: sendFunc,
		timeout: timeout,
	}
}

func (rs *ReliableSender) AckReceived(seq uint64) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	for _, packet := range rs.unacked {
		rs.sendFunc(packet)
	}
}

func encodeUint64(n uint64) []byte {
	b := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		b[7-i] = byte(n >> (i * 8))
	}
	return b
}

// Message Receiving
type ReliableReceiver struct {
	expectedSeq uint64
	recvCh chan []byte
	mu sync.Mutex
	lastMessage []byte
}

func NewReliableReceiver() *ReliableReceiver {
	return &ReliableReceiver{
		recvCh: make(chan []byte, 100),
	}
}

func (rr *ReliableReceiver) Receive(packet []byte) {
	if len(packet) < 8 {
		return
	}
	seq := decodeUint64(packet[:8])
	rr.mu.Lock()
	defer rr.mu.Unlock()
	if seq == rr.expectedSeq+1 {
		rr.expectedSeq = seq
		rr.lastMessage = packet[8:]
		rr.recvCh <- rr.lastMessage
	}
}

func (rr *ReliableReceiver) GetLastMessage() []byte {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	return rr.lastMessage
}

func decodeUint64(b []byte) uint64 {
	var n uint64
	for i := uint(0); i < 8; i++ {
		n = (n << 8) | uint64(b[i])
	}
	return n
}

// Fragmentation Utilities
const MaxFragmentSize = 1024

func FragmentMessage(msg []byte) [][]byte {
	var fragments [][]byte
	for i := 0; i < len(msg); i += MaxFragmentSize {
		end := i + MaxFragmentSize
		if end > len(msg) {
			end = len(msg)
		}
		fragments = append(fragments, msg[i:end])
	}
	return fragments
}

func ReassembleMessage(fragments [][]byte) []byte {
	return bytes.Join(fragments, nil)
}