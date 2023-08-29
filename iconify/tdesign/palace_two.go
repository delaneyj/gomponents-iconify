package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalaceTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6.002 6.002 0 0 1 18 9h2V3h2v19H2V3h2v6h2a6.002 6.002 0 0 1 5.002-5.917l.002-1.085l2 .004ZM8 9h8a4 4 0 0 0-8 0Zm-4 2v9h4v-1.646a6.427 6.427 0 0 1 3.553-5.748l.447-.224l.447.224A6.427 6.427 0 0 1 16 18.354V20h4v-9H4Zm10 9v-1.646a4.428 4.428 0 0 0-2-3.702a4.427 4.427 0 0 0-2 3.702V20h4ZM10.998 5.998h2.004v2.004h-2.004V5.998Z"/>`),
		g.Group(children),
	)
}