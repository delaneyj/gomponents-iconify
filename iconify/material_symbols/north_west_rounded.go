package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NorthWestRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 8.4V14q0 .425-.288.713T6 15q-.425 0-.713-.288T5 14V6q0-.425.288-.713T6 5h8q.425 0 .713.288T15 6q0 .425-.288.713T14 7H8.4l10.9 10.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L7 8.4Z"/>`),
		g.Group(children),
	)
}