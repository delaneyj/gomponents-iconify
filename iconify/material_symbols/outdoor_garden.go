package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutdoorGarden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21h6V6L5 3L2 6v15Zm7 0h6V6l-3-3l-3 3v15Zm7 0h6V6l-3-3l-3 3v15Z"/>`),
		g.Group(children),
	)
}