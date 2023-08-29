package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="32" x="4" y="8" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="3"/><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M4 11C4 9.34315 5.34315 8 7 8H41C42.6569 8 44 9.34315 44 11V20H4V11Z"/><circle r="2" fill="#fff" transform="matrix(0 -1 -1 0 10 14)"/><circle r="2" fill="#fff" transform="matrix(0 -1 -1 0 16 14)"/></g>`),
		g.Group(children),
	)
}