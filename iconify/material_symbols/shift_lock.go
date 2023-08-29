package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiftLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18v-5H3l9-11l9 11h-5v5H8Zm-4 4v-2h16v2H4Z"/>`),
		g.Group(children),
	)
}