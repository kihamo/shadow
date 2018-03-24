package dashboard

type Menu interface {
	Title() string
	Url() string
	ComponentName() string
	Route() Route
	Icon() string
	Childs() []Menu
	Show(request *Request) bool
}

type HasMenu interface {
	DashboardMenu() Menu
}

type MenuItem struct {
	title         string
	url           string
	componentName string
	route         Route
	icon          string
	childs        []Menu
	show          func(*Request) bool
}

func newMenu(title, url, componentName, icon string, route Route, childs []Menu, show func(*Request) bool) Menu {
	return MenuItem{
		title:         title,
		url:           url,
		componentName: componentName,
		icon:          icon,
		route:         route,
		childs:        childs,
		show:          show,
	}
}

func NewMenuWithUrl(title, url, componentName, icon string, childs []Menu, show func(*Request) bool) Menu {
	return newMenu(title, url, componentName, icon, nil, childs, show)
}

func NewMenuWithRoute(title string, route Route, icon string, childs []Menu, show func(*Request) bool) Menu {
	return newMenu(title, "", route.ComponentName(), icon, route, childs, show)
}

func (m MenuItem) Title() string {
	return m.title
}

func (m MenuItem) Url() string {
	return m.url
}

func (m MenuItem) ComponentName() string {
	return m.componentName
}

func (m MenuItem) Route() Route {
	return m.route
}

func (m MenuItem) Icon() string {
	return m.icon
}

func (m MenuItem) Childs() []Menu {
	return m.childs
}

func (m MenuItem) Show(request *Request) bool {
	if m.show != nil {
		return m.show(request)
	}

	return true
}
