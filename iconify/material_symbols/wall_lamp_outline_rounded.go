package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallLampOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20v-4q0-.425.288-.713T4 15q.425 0 .713.288T5 16v4q0 .425-.288.713T4 21Zm4.35-10h9.3l-1.8-6h-5.7l-1.8 6Zm0 0h9.3h-9.3ZM7 19q-.425 0-.712-.288T6 18q0-.425.288-.713T7 17h4q.425 0 .713-.288T12 16v-3H7q-.5 0-.8-.388t-.15-.887l2.4-8q.1-.325.35-.525t.6-.2h7.2q.35 0 .6.2t.35.525l2.4 8q.15.5-.15.888T19 13h-5v3q0 1.25-.875 2.125T11 19H7Z"/>`),
		g.Group(children),
	)
}