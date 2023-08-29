package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NorthRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.425 0-.713-.288T11 21V5.825L7.1 9.7q-.275.275-.687.288T5.7 9.7q-.275-.275-.275-.7t.275-.7l5.6-5.6q.15-.15.325-.212T12 2.425q.2 0 .375.063t.325.212l5.6 5.6q.275.275.275.688T18.3 9.7q-.3.3-.713.3t-.712-.3L13 5.825V21q0 .425-.288.713T12 22Z"/>`),
		g.Group(children),
	)
}