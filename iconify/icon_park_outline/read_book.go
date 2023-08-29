package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 21v23c-3.291-4-13.371-4-18-4V18c9.874 0 16.114 2 18 3Zm0 0v23c3.291-4 13.371-4 18-4V18c-9.874 0-16.114 2-18 3Z"/><circle cx="24" cy="12" r="8"/></g>`),
		g.Group(children),
	)
}