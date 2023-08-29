package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-2h2v-9H2V9l10-7l10 7v2h-2v9h2v2H2Zm4-2h12H6Zm2-2h2v-4l2 3l2-3v4h2v-7h-2l-2 3l-2-3H8v7Zm10 2V8.65l-6-4.2l-6 4.2V20h12Z"/>`),
		g.Group(children),
	)
}