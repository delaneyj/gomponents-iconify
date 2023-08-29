package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataArray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 20v-2h3V6h-3V4h5v16h-5ZM4 20V4h5v2H6v12h3v2H4Z"/>`),
		g.Group(children),
	)
}