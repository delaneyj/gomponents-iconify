package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 19l-7-7l7-7v14Zm4 0V5l7 7l-7 7Zm1.5-3.625L18.875 12L15.5 8.625v6.75Z"/>`),
		g.Group(children),
	)
}