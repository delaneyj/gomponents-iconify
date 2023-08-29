package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAmd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 16V9c0-.566-.434-1-1-1H8L3 3h17c.566 0 1 .434 1 1v17l-5-5z"/><path d="M11.293 20.707L16 16H9a1 1 0 0 1-1-1V8l-4.707 4.707a1 1 0 0 0-.293.707V20a1 1 0 0 0 1 1h6.586a1 1 0 0 0 .707-.293z"/></g>`),
		g.Group(children),
	)
}