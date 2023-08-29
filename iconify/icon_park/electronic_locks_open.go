package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectronicLocksOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="24" height="18" x="12" y="20" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#000" stroke-linecap="round" d="M18 20V14C18 10.3181 20.6863 8 24 8C25.8493 8 27.5032 8.72195 28.6038 10C29.0889 10.5634 29.4666 11.2348 29.7061 12"/><path stroke="#fff" stroke-linecap="round" d="M24 28V30"/><path stroke="#000" stroke-linecap="round" d="M6 18V30"/><path stroke="#000" stroke-linecap="round" d="M42 18V30"/></g>`),
		g.Group(children),
	)
}