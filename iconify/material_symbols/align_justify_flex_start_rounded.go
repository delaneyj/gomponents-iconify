package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustifyFlexStartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21V3q0-.425.288-.713T3 2q.425 0 .713.288T4 3v18q0 .425-.288.713T3 22Zm11-5q-.425 0-.713-.288T13 16V8q0-.425.288-.713T14 7h1q.425 0 .713.288T16 8v8q0 .425-.288.713T15 17h-1Zm-6 0q-.425 0-.713-.288T7 16V8q0-.425.288-.713T8 7h1q.425 0 .713.288T10 8v8q0 .425-.288.713T9 17H8Z"/>`),
		g.Group(children),
	)
}