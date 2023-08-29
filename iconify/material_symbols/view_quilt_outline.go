package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewQuiltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5h18v14H3V5Zm7.325 2v4H19V7h-8.675ZM19 13h-3.325v4H19v-4Zm-8.675 0v4h3.35v-4h-3.35ZM5 17h3.325V7H5v10Z"/>`),
		g.Group(children),
	)
}