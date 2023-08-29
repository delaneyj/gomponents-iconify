package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaseSensitive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 15l4-8l4 8m-7-2h6"/><circle cx="18" cy="12" r="3"/><path d="M21 9v6"/></g>`),
		g.Group(children),
	)
}