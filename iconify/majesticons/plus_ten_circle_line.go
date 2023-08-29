package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusTenCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12h2m2 0H7m0 0v-2m0 2v2m3-5h2v6m6-1.5v-3A1.5 1.5 0 0 0 16.5 9v0a1.5 1.5 0 0 0-1.5 1.5v3a1.5 1.5 0 0 0 1.5 1.5v0a1.5 1.5 0 0 0 1.5-1.5z"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`),
		g.Group(children),
	)
}