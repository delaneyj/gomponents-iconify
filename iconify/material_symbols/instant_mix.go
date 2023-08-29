package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstantMix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-7H3v-2h6v2H7v7H5ZM5 9V4h2v5H5Zm4 0V7h2V4h2v3h2v2H9Zm2 11v-9h2v9h-2Zm6 0v-3h-2v-2h6v2h-2v3h-2Zm0-7V4h2v9h-2Z"/>`),
		g.Group(children),
	)
}