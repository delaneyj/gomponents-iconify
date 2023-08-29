package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20V7H9V4h13v3h-5v13h-3Zm-9 0v-8H2V9h9v3H8v8H5Z"/>`),
		g.Group(children),
	)
}