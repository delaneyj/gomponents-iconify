package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 19l-7-7l7-7v14Zm-1.5-3.625v-6.75L5.125 12L8.5 15.375ZM14 19V5l7 7l-7 7Z"/>`),
		g.Group(children),
	)
}