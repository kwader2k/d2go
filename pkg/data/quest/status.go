package quest

const (
	StatusRewardGranted Status = 1 << iota
	StatusRewardPending
	StatusStarted
	StatusLeaveTown
	StatusEnterArea
	StatusInProgress1
	StatusInProgress2
	StatusInProgress3
	StatusInProgress4
	StatusInProgress5
	StatusInProgress6
	StatusInProgress7
	StatusUpdateQuestLog
	StatusPrimaryGoalDone
	StatusCompletedNow
	StatusCompletedBefore
)

const StatusMaskCompleted = StatusPrimaryGoalDone |
	StatusRewardGranted |
	StatusRewardPending |
	StatusCompletedNow |
	StatusCompletedBefore |
	StatusUpdateQuestLog
