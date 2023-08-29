package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiftLockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18v-5H3l9-11l9 11h-5v5H8Zm2-2h4v-5h2.775L12 5.15L7.225 11H10v5Zm2-5.425ZM4 22v-2h16v2H4Z"/>`),
		g.Group(children),
	)
}