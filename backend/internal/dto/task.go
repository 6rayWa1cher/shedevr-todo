package dto

type Task struct {
	ID        int64
	Title     string
	Text      string
	Completed CompletedStatusType
}

type CompletedStatusType string

const (
	CompletedYes       CompletedStatusType = "yes"
	CompletedNo        CompletedStatusType = "no"
	CompletedCancelled CompletedStatusType = "cancelled"
)
