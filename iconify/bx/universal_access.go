package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UniversalAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="7.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8z"/><path fill="currentColor" d="M16.5 10.5L16 9l-3 1h-2L8 9l-.5 1.5l3 1V13L9 17.25l1.5.75l1.25-3.5h.5L13.5 18l1.5-.75L13.5 13v-1.5l3-1z"/>`),
		g.Group(children),
	)
}