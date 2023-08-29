package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.986 12.502a9 9 0 1 0-5.973 7.98"/><path d="M12 7v5l3 3m4 1v3m0 3v.01"/></g>`),
		g.Group(children),
	)
}