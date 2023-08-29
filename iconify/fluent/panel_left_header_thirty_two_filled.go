package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M30 23.5a4.5 4.5 0 0 1-4.5 4.5h-19A4.5 4.5 0 0 1 2 23.5v-15A4.5 4.5 0 0 1 6.5 4h19A4.5 4.5 0 0 1 30 8.5v15Zm-2-15A2.5 2.5 0 0 0 25.5 6H12v5h16V8.5Zm0 4.5H12v13h13.5a2.5 2.5 0 0 0 2.5-2.5V13Z"/>`),
		g.Group(children),
	)
}