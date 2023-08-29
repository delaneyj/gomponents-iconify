package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignFlexEndRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h18q.425 0 .713.288T22 21q0 .425-.288.713T21 22H3Zm8.5-4q-.425 0-.713-.288T10.5 17V5q0-.425.288-.713T11.5 4h1q.425 0 .713.288T13.5 5v12q0 .425-.288.713T12.5 18h-1Z"/>`),
		g.Group(children),
	)
}