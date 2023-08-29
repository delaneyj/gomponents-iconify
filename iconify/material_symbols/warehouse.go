package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warehouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V7l10-4l10 4v14h-6v-8H8v8H2Zm7 0v-2h2v2H9Zm2-3v-2h2v2h-2Zm2 3v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}