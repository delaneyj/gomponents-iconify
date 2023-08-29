package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BigClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 4C12.954 4 4 12.865 4 23.8V44h40V23.8C44 12.865 35.046 4 24 4Z"/><circle cx="24" cy="24" r="12"/><path d="M24 18v6l4 4"/></g>`),
		g.Group(children),
	)
}