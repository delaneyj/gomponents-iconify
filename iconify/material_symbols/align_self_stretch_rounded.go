package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignSelfStretchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h18q.425 0 .713.288T22 21q0 .425-.288.713T21 22H3Zm8.5-4.5q-.425 0-.713-.288T10.5 16.5V7q0-.425.288-.713T11.5 6h1q.425 0 .713.288T13.5 7v9.5q0 .425-.288.713t-.712.287h-1ZM3 4q-.425 0-.713-.288T2 3q0-.425.288-.713T3 2h18q.425 0 .713.288T22 3q0 .425-.288.713T21 4H3Z"/>`),
		g.Group(children),
	)
}