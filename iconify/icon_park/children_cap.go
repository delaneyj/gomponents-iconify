package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildrenCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 32C6 17 11 8 14 8C17 8 20 14 20 14H28C28 14 31 8 34 8C37 8 42 17 42 32"/><rect width="40" height="8" x="4" y="32" fill="#2F88FF" rx="2"/></g>`),
		g.Group(children),
	)
}