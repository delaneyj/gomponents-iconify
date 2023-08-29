package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignFlexStartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4q-.425 0-.713-.288T2 3q0-.425.288-.713T3 2h18q.425 0 .713.288T22 3q0 .425-.288.713T21 4H3Zm8.5 16q-.425 0-.713-.288T10.5 19V7q0-.425.288-.713T11.5 6h1q.425 0 .713.288T13.5 7v12q0 .425-.288.713T12.5 20h-1Z"/>`),
		g.Group(children),
	)
}