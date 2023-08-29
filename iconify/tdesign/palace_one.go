package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalaceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6.002 6.002 0 0 1 18 9h4v13H2V9h4a6.002 6.002 0 0 1 5.002-5.917l.002-1.085l2 .004ZM8 9h8a4 4 0 0 0-8 0Zm2.998-3.002h2.004v2.004h-2.004V5.998ZM4 11v9h1v-3a3 3 0 1 1 6 0v3h2v-3a3 3 0 1 1 6 0v3h1v-9H4Zm13 9v-3a1 1 0 1 0-2 0v3h2Zm-8 0v-3a1 1 0 1 0-2 0v3h2Z"/>`),
		g.Group(children),
	)
}