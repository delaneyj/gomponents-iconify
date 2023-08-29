package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.175 10.575L4.7 6.1q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l4.475 4.475l-1.4 1.4ZM15 20q-.425 0-.713-.288T14 19q0-.425.288-.713T15 18h1.6l-3.175-3.175L14.85 13.4L18 16.55V15q0-.425.288-.713T19 14q.425 0 .713.288T20 15v4q0 .425-.288.713T19 20h-4Zm-10.3-.7q-.275-.275-.275-.7t.275-.7L16.6 6H15q-.425 0-.712-.287T14 5q0-.425.288-.713T15 4h4q.425 0 .713.288T20 5v4q0 .425-.288.713T19 10q-.425 0-.713-.288T18 9V7.4L6.1 19.3q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}