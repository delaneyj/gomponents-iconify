package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20V4h2v16H4Zm12.45-4.45l-1.4-1.425L16.175 13H8v-2h8.175L15.05 9.875l1.4-1.425L20 12l-3.55 3.55ZM12 20v-5h2v5h-2Zm0-11V4h2v5h-2Z"/>`),
		g.Group(children),
	)
}