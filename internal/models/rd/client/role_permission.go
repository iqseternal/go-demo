package rd_client

// RolePermission 角色权限表
type RolePermission struct {
	RoleID       uint `gorm:"not null;" json:"role_id"`       // 角色 ID，普通字段
	PermissionID uint `gorm:"not null;" json:"permission_id"` // 权限 ID，普通字段
}

func (c *RolePermission) TableName() string {
	return "rapid.client.role_permission"
}
