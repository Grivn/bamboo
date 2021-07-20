package types

import "time"

type View int

type Statistics struct {
	Thrus                string
	LastViewTime         time.Time
	StartTime            time.Time
	TmpTime              time.Time
	VoteStart            time.Time
	TotalCreateDuration  time.Duration
	TotalProcessDuration time.Duration
	TotalProposeDuration time.Duration
	TotalDelay           time.Duration
	TotalRoundTime       time.Duration
	TotalVoteTime        time.Duration
	TotalBlockSize       int
	ReceivedNo           int
	RoundNo              int
	VoteNo               int
	TotalCommittedTx     int
	LatencyNo            int
	ProposedNo           int
	ProcessedNo          int
	CommittedNo          int
}
