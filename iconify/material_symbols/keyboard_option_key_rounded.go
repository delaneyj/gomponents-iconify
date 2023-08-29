package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardOptionKeyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.925 19q-.55 0-1-.263T14.2 18L7.85 7H4q-.425 0-.713-.288T3 6q0-.425.288-.713T4 5h3.85q.55 0 1 .263T9.575 6l6.35 11H20q.425 0 .713.288T21 18q0 .425-.288.713T20 19h-4.075ZM16 7q-.425 0-.713-.288T15 6q0-.425.288-.713T16 5h4q.425 0 .713.288T21 6q0 .425-.288.713T20 7h-4Z"/>`),
		g.Group(children),
	)
}