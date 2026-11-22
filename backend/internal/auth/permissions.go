package auth

import (
	"bounty_board/internal/constants"
)

type Permission string

const (
	PermissionGetRequests    Permission = "request:get_requests"
	PermissionApproveRequest Permission = "request:approve_requests"
	PermissionDeclineRequest Permission = "request:decline_requests"

	PermissionTaskCreate          Permission = "task:create"
	PermissionTaskApproveFinished Permission = "task:approve_finished"
	PermissionTaskDeclineFinished Permission = "task:decline_finished"

	PermissionGetTable           Permission = "table:get"
	PermissionTaskRead           Permission = "task:read"
	PermissionTaskOwn            Permission = "task:own"
	PermissionTaskFinish         Permission = "task:finish"
	PermissionTicketDeclineOwned Permission = "task:decline_owned"
)

var RolePermission = map[constants.UserRole]map[Permission]struct{}{
	constants.RoleUser: {
		PermissionGetTable:           {},
		PermissionTaskRead:           {},
		PermissionTaskOwn:            {},
		PermissionTaskFinish:         {},
		PermissionTicketDeclineOwned: {},
	},
	constants.RoleOwner: {
		PermissionGetRequests:    {},
		PermissionApproveRequest: {},
		PermissionDeclineRequest: {},

		PermissionGetTable:           {},
		PermissionTaskRead:           {},
		PermissionTaskOwn:            {},
		PermissionTaskFinish:         {},
		PermissionTicketDeclineOwned: {},

		PermissionTaskCreate:          {},
		PermissionTaskApproveFinished: {},
		PermissionTaskDeclineFinished: {},
	},
}

func HasPermission(role constants.UserRole, permission Permission) bool {
	permissions, exists := RolePermission[role]
	if !exists {
		return false
	}

	_, exists = permissions[permission]

	return exists
}
