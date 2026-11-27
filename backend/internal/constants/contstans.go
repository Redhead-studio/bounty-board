package constants

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleOwner UserRole = "owner"
)

type TaskStatus string

const (
	StatusNew     TaskStatus = "new"
	StatusPending TaskStatus = "pending"
	StatusRewiev  TaskStatus = "review"
	StatusClosed  TaskStatus = "closed"
)

type VerdictStatus string

const (
	StatusAwaiting VerdictStatus = "awaiting"
	StatusApproved VerdictStatus = "approved"
	StatusDeclined VerdictStatus = "declined"
)

type RoomMemberRole string

const (
	RoomRoleMember RoomMemberRole = "member"
	RoomRoleMaster RoomMemberRole = "master"
)

type RoomMemberStatus string

const (
	RoomMemberStatusActive  RoomMemberStatus = "active"
	RoomMemberStatusBlocked RoomMemberStatus = "blocked"
)

type JoinRequestStatus string

const (
	JoinRequestStatusPending  JoinRequestStatus = "pending"
	JoinRequestStatusApproved JoinRequestStatus = "approved"
	JoinRequestStatusDeclined JoinRequestStatus = "declined"
)

type BountyStatus string

const (
	BountyStatusOpen       BountyStatus = "open"
	BountyStatusInProgress BountyStatus = "in_progress"
	BountyStatusReview     BountyStatus = "review"
	BountyStatusCompleted  BountyStatus = "completed"
	BountyStatusRejected   BountyStatus = "rejected"
)

type SubmissionStatus string

const (
	SubmissionStatusPending  SubmissionStatus = "pending"
	SubmissionStatusApproved SubmissionStatus = "approved"
	SubmissionStatusRejected SubmissionStatus = "rejected"
)
