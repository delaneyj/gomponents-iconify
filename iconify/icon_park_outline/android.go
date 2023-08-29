package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Android(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M43.901 36H4.1C5.103 25.893 13.63 18 24 18c10.372 0 18.899 7.893 19.902 18Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m14 20l-4-7m23 7l4-7"/><circle cx="15" cy="29" r="2" fill="currentColor"/><circle cx="33" cy="29" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}