package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LowDensity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9V7h2v2H3Zm0-4V3h2v2H3Zm4 8v-2h2v2H7Zm0-8V3h2v2H7Zm4 8v-2h2v2h-2Zm0-4V7h2v2h-2ZM3 21V11h2v8h14V5h-8V3h10v18H3Z"/>`),
		g.Group(children),
	)
}