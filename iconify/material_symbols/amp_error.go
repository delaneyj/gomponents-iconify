package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmpError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.275 21L3 15.725v-7.45L8.275 3h7.45L21 8.275v7.45L15.725 21Zm2.075-1.5h.8L17 10.35h-4.2l.85-5.85h-.8l-5.9 9.15h4.2Z"/>`),
		g.Group(children),
	)
}