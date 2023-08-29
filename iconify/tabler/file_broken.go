package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 3v4a1 1 0 0 0 1 1h4"/><path d="M5 7V5a2 2 0 0 1 2-2h7l5 5v2m0 9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2m0-3h.01M5 13h.01M5 10h.01M19 13h.01M19 16h.01"/></g>`),
		g.Group(children),
	)
}