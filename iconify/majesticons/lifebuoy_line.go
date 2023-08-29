package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LifebuoyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m18 6l-3.172 3.172M6 18l3.172-3.172M6 6l3.172 3.172M18 18l-3.172-3.172m0-5.656A3.987 3.987 0 0 0 12 8a3.987 3.987 0 0 0-2.828 1.172m5.656 0A3.984 3.984 0 0 1 16 12a3.987 3.987 0 0 1-1.172 2.828M9.172 9.172A3.987 3.987 0 0 0 8 12c0 1.105.448 2.105 1.172 2.828m0 0A3.987 3.987 0 0 0 12 16a3.987 3.987 0 0 0 2.828-1.172"/></g>`),
		g.Group(children),
	)
}