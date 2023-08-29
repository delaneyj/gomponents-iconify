package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperFoldText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 21h12c-1 0-3-.6-3-3v-2H3v2c0 2.4 2 3 3 3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v13c0 1-.6 3-3 3m0 0H6c-1 0-3-.6-3-3v-2h12v2c0 2.4 2 3 3 3zM9 7h8m-8 4h4"/></g>`),
		g.Group(children),
	)
}