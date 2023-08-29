package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikeThroughLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 5h-7a3 3 0 0 0-3 3v1a3 3 0 0 0 3 3h7M7 19h7a3 3 0 0 0 3-3v-1M5 12h14"/>`),
		g.Group(children),
	)
}