package model

type Userinfo struct {
	UserinfoId int64  `gorose:"userinfo_id" json:"userinfo_id"`
	UserId     int64  `gorose:"user_id" json:"user_id"` // user.id
	Avatar     string `gorose:"avatar" json:"avatar"`   // 头像
	Remark     string `gorose:"remark" json:"remark"`   // 说明
}

func (Userinfo) TableName() string {
	return "userinfo"
}

type CasbinRule struct {
	Id    int64  `gorose:"id" json:"id"`
	PType string `gorose:"p_type" json:"p_type"` // perm类型：p,g......
	V0    string `gorose:"v0" json:"v0"`         // 角色名字...
	V1    string `gorose:"v1" json:"v1"`         // 对象资源...
	V2    string `gorose:"v2" json:"v2"`         // 权限值...
	V3    string `gorose:"v3" json:"v3"`         // ext
	V4    string `gorose:"v4" json:"v4"`         // ext
	V5    string `gorose:"v5" json:"v5"`         // ext
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}

type Resource struct {
	Id            int64  `gorose:"id" json:"id"`
	ResourceTitle string `gorose:"resource_title" json:"resource_title"` // 资源名称
	ResourceName  string `gorose:"resource_name" json:"resource_name"`   // 资源中心,如:user_center
	PName         string `gorose:"p_name" json:"p_name"`                 // 父节点名字
	Level         int64  `gorose:"level" json:"level"`                   // 级别:默认1目录,2子目录,3按钮
	UserId        int64  `gorose:"user_id" json:"user_id"`               // 操作人
	Remark        string `gorose:"remark" json:"remark"`                 // 备注
	CreatedAt     string `gorose:"created_at" json:"created_at"`
	Sort          int64  `gorose:"sort" json:"sort"` // 排序,根据数值从小决定目录位置:默认99很大放最后
}

func (Resource) TableName() string {
	return "resource"
}

type Role struct {
	Id        int64  `gorose:"id" json:"id"`
	RoleName  string `gorose:"role_name" json:"role_name"`   // 角色名字
	RoleTitle string `gorose:"role_title" json:"role_title"` // 角色标题
	UserId    int64  `gorose:"user_id" json:"user_id"`       // 操作人
	Remark    string `gorose:"remark" json:"remark"`         // 备注
	CreatedAt string `gorose:"created_at" json:"created_at"`
}

func (Role) TableName() string {
	return "role"
}

type User struct {
	Id         int64  `gorose:"id" json:"id"`
	Username   string `gorose:"username" json:"username"`       // 用户名
	RoleName   string `gorose:"role_name" json:"role_name"`     // role.role_name
	UserStatus int64  `gorose:"user_status" json:"user_status"` // 用户状态:默认0无效,1有效
	Password   string `gorose:"password" json:"password"`       // 密码
	Nickname   string `gorose:"nickname" json:"nickname"`       // 用户的昵称
	CreatedAt  string `gorose:"created_at" json:"created_at"`
}

func (User) TableName() string {
	return "user"
}
