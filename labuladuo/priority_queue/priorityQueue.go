package priority_queue

type PriorityQueue struct {
	Size int
	Pqs  []int
}

func NewPriorityQueue(cap int) *PriorityQueue {
	keys := make([]int, cap+1)
	return &PriorityQueue{0, keys}
}

// InsertMax - 插入元素到最小堆
func (pq *PriorityQueue) InsertMax(key int, val int) {
	pq.Size++
	pq.Pqs[key] = val

}

// InsertMin - 插入元素
func (pq *PriorityQueue) InsertMin() {

}

// DelMax - 删除最大元素并返回
func (pq *PriorityQueue) DelMax() {

}

// DelMin - 删除最小元素并返回
func (pq *PriorityQueue) DelMin() {

}

// SwimMax - 上浮第x个元素，维护最大堆的性质
func (pq *PriorityQueue) swimMax(k int) {

}

// SinkMax - 下沉第x个元素，维护最大堆的性质
func (pq *PriorityQueue) sinkMax(k int) {

}

// SwimMin - 上浮第x个元素，维护最小堆的性质
func (pq *PriorityQueue) swimMin() {

}

// SinkMin - 下沉第x个元素，维护最小堆的性质
func (pq *PriorityQueue) sinkMin(k int) {

}
