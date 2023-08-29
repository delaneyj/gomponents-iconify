package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustifyFlexEndRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 22q-.425 0-.713-.288T20 21V3q0-.425.288-.713T21 2q.425 0 .713.288T22 3v18q0 .425-.288.713T21 22Zm-6-5q-.425 0-.713-.288T14 16V8q0-.425.288-.713T15 7h1q.425 0 .713.288T17 8v8q0 .425-.288.713T16 17h-1Zm-6 0q-.425 0-.713-.288T8 16V8q0-.425.288-.713T9 7h1q.425 0 .713.288T11 8v8q0 .425-.288.713T10 17H9Z"/>`),
		g.Group(children),
	)
}