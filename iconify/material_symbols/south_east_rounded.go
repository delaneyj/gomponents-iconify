package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthEastRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.6 17L4.7 6.1q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275L17 15.6V10q0-.425.288-.713T18 9q.425 0 .713.288T19 10v8q0 .425-.288.713T18 19h-8q-.425 0-.713-.288T9 18q0-.425.288-.713T10 17h5.6Z"/>`),
		g.Group(children),
	)
}