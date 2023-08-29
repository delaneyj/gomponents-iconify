package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarehouseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19h2v-8h12v8h2V8.35l-8-3.2l-8 3.2V19Zm-2 2V7l10-4l10 4v14h-6v-8H8v8H2Zm7 0v-2h2v2H9Zm2-3v-2h2v2h-2Zm2 3v-2h2v2h-2ZM6 11h12H6Z"/>`),
		g.Group(children),
	)
}