package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CategoryManagement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="36" height="14" x="6" y="28" stroke="#000" stroke-width="4" rx="4"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M20 7H10C7.79086 7 6 8.79086 6 11V17C6 19.2091 7.79086 21 10 21H20"/><circle cx="34" cy="14" r="8" fill="#2F88FF" stroke="#000" stroke-width="4"/><circle cx="34" cy="14" r="3" fill="#fff"/></g>`),
		g.Group(children),
	)
}