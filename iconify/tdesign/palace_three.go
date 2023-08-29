package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalaceThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6.002 6.002 0 0 1 18 9h4v13H2V9h4a6.002 6.002 0 0 1 5.002-5.917l.002-1.085l2 .004ZM12 5a4 4 0 0 0-4 4v2H4v9h4v-1.646a6.427 6.427 0 0 1 3.553-5.748l.447-.224l.447.224A6.427 6.427 0 0 1 16 18.354V20h4v-9h-4V9a4 4 0 0 0-4-4Zm2 15v-1.646a4.428 4.428 0 0 0-2-3.702a4.427 4.427 0 0 0-2 3.702V20h4ZM10.998 7.998h2.004v2.004h-2.004V7.998Z"/>`),
		g.Group(children),
	)
}