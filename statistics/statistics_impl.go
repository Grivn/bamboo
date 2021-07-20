package statistics

import (
	"fmt"
	"time"
)

type Statistics interface {
	Query() string
	AdvanceReceivedNo()
	AddLatency(duration time.Duration)
	AddCommittedTxCount(count int)
	ProposedBlockSize(size int, duration time.Duration)
	LastView()
	VotedTime()
	RoundTime()
	StartSignal()
	ProposedDuration(duration time.Duration)
	ProcessedDuration(duration time.Duration)
	VotedDuration(duration time.Duration)
}

type statisticsImpl struct {
	/* for monitoring node statistics */
	thrus                string
	lastViewTime         time.Time
	startTime            time.Time
	tmpTime              time.Time
	voteStart            time.Time
	totalCreateDuration  time.Duration
	totalProcessDuration time.Duration
	totalProposeDuration time.Duration
	totalDelay           time.Duration
	totalRoundTime       time.Duration
	totalVoteTime        time.Duration
	totalBlockSize       int
	receivedNo           int
	roundNo              int
	voteNo               int
	totalCommittedTx     int
	latencyNo            int
	proposedNo           int
	processedNo          int
	committedNo          int
}

func NewStatistics() *statisticsImpl {
	return &statisticsImpl{}
}

func (si *statisticsImpl) Query() string {
	aveCreateDuration := float64(si.totalCreateDuration.Milliseconds()) / float64(si.proposedNo)
	aveProcessTime := float64(si.totalProcessDuration.Milliseconds()) / float64(si.processedNo)
	aveVoteProcessTime := float64(si.totalVoteTime.Milliseconds()) / float64(si.voteNo)
	aveBlockSize := float64(si.totalBlockSize) / float64(si.proposedNo)
	aveRoundTime := float64(si.totalRoundTime.Milliseconds()) / float64(si.roundNo)
	latency := float64(si.totalDelay.Milliseconds()) / float64(si.latencyNo)
	si.thrus += fmt.Sprintf("Time: %v s. Throughput: %v txs/s\n",
		time.Now().Sub(si.startTime).Seconds(), float64(si.totalCommittedTx)/time.Now().Sub(si.tmpTime).Seconds())
	si.totalCommittedTx = 0
	si.tmpTime = time.Now()
	//status := fmt.Sprintf("chain status is: %s\nCommitted rate is %v.\nAve. block size is %v.\nAve. trans. delay is %v ms.\nAve. creation time is %f ms.\nAve. processing time is %v ms.\nAve. vote time is %v ms.\nRequest rate is %f txs/s.\nAve. round time is %f ms.\nLatency is %f ms.\nThroughput is %f txs/s.\n", si.Safety.GetChainStatus(), committedRate, aveBlockSize, aveTransDelay, aveCreateDuration, aveProcessTime, aveVoteProcessTime, requestRate, aveRoundTime, latency, throughput)
	status := fmt.Sprintf(""+
		"Ave. block size is %v.\n"+
		"Ave. creation time is %f ms.\n"+
		"Ave. processing time is %v ms.\n"+
		"Ave. vote time is %v ms.\n"+
		"Ave. round time is %f ms.\n"+
		"Latency is %f ms.\n"+
		"Throughput is: \n%v",
		aveBlockSize, aveCreateDuration, aveProcessTime, aveVoteProcessTime, aveRoundTime, latency, si.thrus)
	return status
}

func (si *statisticsImpl) AdvanceReceivedNo() {
	si.receivedNo++
}

func (si *statisticsImpl) AddLatency(duration time.Duration) {
	si.totalDelay += duration
	si.latencyNo++
}

func (si *statisticsImpl) AddCommittedTxCount(count int) {
	si.totalCommittedTx += count
	si.committedNo++
}

func (si *statisticsImpl) ProposedBlockSize(size int, duration time.Duration) {
	si.totalBlockSize += size
	si.proposedNo++
	si.totalCreateDuration += duration
	si.voteStart = time.Now()
}

func (si *statisticsImpl) LastView() {
	si.lastViewTime = time.Now()
}

func (si *statisticsImpl) VotedTime() {
	si.totalVoteTime += time.Now().Sub(si.voteStart)
}

func (si *statisticsImpl) RoundTime() {
	now := time.Now()
	lasts := now.Sub(si.lastViewTime)
	si.totalRoundTime += lasts
	si.roundNo++
	si.lastViewTime = now
}

func (si *statisticsImpl) StartSignal() {
	si.startTime = time.Now()
	si.tmpTime = time.Now()
}

func (si *statisticsImpl) ProposedDuration(duration time.Duration) {
	si.totalProposeDuration += duration
}

func (si *statisticsImpl) ProcessedDuration(duration time.Duration) {
	si.totalProcessDuration += duration
	si.voteStart = time.Now()
	si.processedNo++
}

func (si *statisticsImpl) VotedDuration(duration time.Duration) {
	si.totalVoteTime += duration
	si.voteNo++
}
