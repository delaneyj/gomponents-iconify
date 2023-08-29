package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosswordOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20h4v-4h-4v4Zm-6-6h4v-4H4v4Zm6 0h4v-4h-4v4Zm6 0h4v-4h-4v4Zm0-6h4V4h-4v4ZM8 22v-6H2V8h12V2h8v14h-6v6H8Z"/>`),
		g.Group(children),
	)
}