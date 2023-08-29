package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="8" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="m41 8l-12 9.59V44M10.111 23.25L19 18v10.917L7 41"/></g>`),
		g.Group(children),
	)
}