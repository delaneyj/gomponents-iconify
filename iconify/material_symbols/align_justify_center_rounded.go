package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustifyCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.425 0-.713-.288T11 21V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v18q0 .425-.288.713T12 22Zm4-5q-.425 0-.713-.288T15 16V8q0-.425.288-.713T16 7h1q.425 0 .713.288T18 8v8q0 .425-.288.713T17 17h-1Zm-9 0q-.425 0-.713-.288T6 16V8q0-.425.288-.713T7 7h1q.425 0 .713.288T9 8v8q0 .425-.288.713T8 17H7Z"/>`),
		g.Group(children),
	)
}