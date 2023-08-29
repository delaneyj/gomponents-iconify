package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 6.825l-.9.9Q9.825 8 9.412 8T8.7 7.7q-.275-.275-.275-.7t.275-.7l2.6-2.6q.3-.3.7-.3t.7.3l2.6 2.6q.275.275.287.687T15.3 7.7q-.275.275-.7.275t-.7-.275l-.9-.875V20q0 .425-.287.713T12 21q-.425 0-.713-.288T11 20V6.825Z"/>`),
		g.Group(children),
	)
}