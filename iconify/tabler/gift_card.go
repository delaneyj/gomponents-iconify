package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 8a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3z"/><path d="m7 16l3-3l3 3m-5-3c-.789 0-2-.672-2-1.5S6.711 10 7.5 10c1.128-.02 2.077 1.17 2.5 3c.423-1.83 1.372-3.02 2.5-3c.789 0 1.5.672 1.5 1.5S12.789 13 12 13H8z"/></g>`),
		g.Group(children),
	)
}