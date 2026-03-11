package helper

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type ID int64

var ( 
	mu        sync.Mutex
	Epoch     int64 = 1288834974657
	NodeBits  uint8 = 10
	StepBits  uint8 = 12
	nodeMax   int64 = -1 ^ (-1 << NodeBits)
	nodeMask        = nodeMax << StepBits
	stepMask  int64 = -1 ^ (-1 << StepBits)
	timeShift       = NodeBits + StepBits
	nodeShift       = StepBits
)

type Node struct {
	mu        sync.Mutex
	epoch     time.Time
	time      int64
	node      int64
	step      int64
	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

func NewSnowflakeNode(node int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("Remember, you have a total 22 bits to share between Node/Step")
	}
	// re-calc in case custom NodeBits or StepBits were set
	mu.Lock()
	nodeMax = -1 ^ (-1 << NodeBits)
	nodeMask = nodeMax << StepBits
	stepMask = -1 ^ (-1 << StepBits)
	timeShift = NodeBits + StepBits
	nodeShift = StepBits
	mu.Unlock()

	n := Node{}
	n.node = node
	n.nodeMax = -1 ^ (-1 << NodeBits)
	n.nodeMask = n.nodeMax << StepBits
	n.stepMask = -1 ^ (-1 << StepBits)
	n.timeShift = NodeBits + StepBits
	n.nodeShift = StepBits

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(n.nodeMax, 10))
	}

	var curTime = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	n.epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime))

	return &n, nil
}

func Int64ToBase62(n int64) string {
	if n == 0 {
		return "a"
	}

	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base := int64(len(chars))

	buf := make([]byte, 0, 11) // enough for int64 base62
	for n > 0 {
		buf = append(buf, chars[n%base])
		n /= base
	}

	// reverse
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}

func (n *Node) GenerateID() string {

	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Since(n.epoch).Milliseconds()

	if now == n.time {
		n.step = (n.step + 1) & n.stepMask

		if n.step == 0 {
			for now <= n.time {
				now = time.Since(n.epoch).Milliseconds()
			}
		}
	} else {
		n.step = 0
	}

	n.time = now

	r := int64((now)<<n.timeShift | (n.node << n.nodeShift) | (n.step))

	
	result:=Int64ToBase62(r)
	
	return result
}
