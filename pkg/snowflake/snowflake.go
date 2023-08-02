package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	startTimeStr = "2023-08-01"

	timestampBits  = 41
	workerIDBits   = 10
	sequenceIDBits = 12

	timestampMax  = -1 ^ (-1 << timestampBits)  //求数据范围最大值
	workerIDMax   = -1 ^ (-1 << workerIDBits)   //求数据范围最大值
	sequenceIDMax = -1 ^ (-1 << sequenceIDBits) //求数据范围最大值
	t
	sequenceIDShift = 0 //存储位置
	workerIDShift   = sequenceIDShift + sequenceIDBits
	timestampShift  = workerIDShift + workerIDBits
)

var startTime, _ = time.Parse("2006-01-02", startTimeStr)

type Worker struct {
	mu sync.Mutex //同步
	//  也可以使用int64 定义，因为最高因为不使用，不会超出数据范围
	timestamp uint64 //时间戳
	workerID  uint64 //节点ID
	sequence  uint64 //序列号
}

func NewWorker(workerID uint64) (worker *Worker, err error) {
	if workerID < 0 || workerID > workerIDMax {
		return worker, errors.New("worker ID excess of quantity")
	}
	return &Worker{
		workerID: workerID,
	}, err
}
func (w *Worker) GetID() uint64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	//为什么不能直接使用time.Now().UnixMilli()？
	currentTimestamp := uint64(time.Now().UnixNano() / 1e6)
	if currentTimestamp == w.timestamp {
		w.sequence++
		if w.sequence > sequenceIDMax {
			//如果超过了最大值，则进入死循环，直接到下一ms
			for currentTimestamp != w.timestamp {
				w.sequence = 0
				w.timestamp = currentTimestamp
			}
		}
	} else {
		w.sequence = 0
		w.timestamp = currentTimestamp
	}
	//生成ID
	ID := w.timestamp<<timestampShift | w.workerID<<workerIDShift | w.sequence<<sequenceIDShift
	return ID
}
func (w *Worker) GetIDHex() string {
	return strconv.FormatInt(int64(w.GetID()), 16)
}
