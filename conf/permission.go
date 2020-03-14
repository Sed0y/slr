package conf

type Permissions struct {
	Permissions []RoutPermission
	Menus       []MenuPermissions
}

func (p *Permissions) GetRoles(URL string) []int {

	for i := 0; i < len(p.Permissions); i++ {
		if p.Permissions[i].URL == URL {
			return p.Permissions[i].Roles
		}
	}

	return []int{}
}

type RoutPermission struct {
	URL   string
	Roles []int
}

type MenuPermissions struct {
	Name  string
	Roles []int
}

// Role
// Роли:
// 0 - просто сотрудник филиала
// 1 - администратор
// 2 - остальные...

func (p *Permissions) Init() {

	p.Permissions = p.Permissions[:0]
	p.Menus = p.Menus[:0]

	p.Permissions = append(p.Permissions, RoutPermission{URL: "/upload/rosfin", Roles: []int{1}})

	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin", Roles: []int{1}})
	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin/statistics", Roles: []int{1}})

	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin/databases", Roles: []int{1}})
	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin/databases/update", Roles: []int{1}})

	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin/rosfin", Roles: []int{1}})
	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin/rosfin/parse", Roles: []int{1}})
	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin/rosfin/parse/progress", Roles: []int{1}})
	p.Permissions = append(p.Permissions, RoutPermission{URL: "/admin/rosfin/delete", Roles: []int{1}})

	p.Permissions = append(p.Permissions, RoutPermission{URL: "/charts", Roles: []int{1, 2}})

	p.Menus = append(p.Menus, MenuPermissions{Name: "admin", Roles: []int{1}})
	//p.Menus = append(p.Menus, MenuPermissions{Name: "admin", Roles: []int{1}})

}
