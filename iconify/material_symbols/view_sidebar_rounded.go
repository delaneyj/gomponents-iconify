package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSidebarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8q-.425 0-.713-.288T18 7V5q0-.425.288-.713T19 4h2q.425 0 .713.288T22 5v2q0 .425-.288.713T21 8h-2Zm0 6q-.425 0-.713-.288T18 13v-2q0-.425.288-.713T19 10h2q.425 0 .713.288T22 11v2q0 .425-.288.713T21 14h-2ZM3 20q-.425 0-.713-.288T2 19V5q0-.425.288-.713T3 4h12q.425 0 .713.288T16 5v14q0 .425-.288.713T15 20H3Zm16 0q-.425 0-.713-.288T18 19v-2q0-.425.288-.713T19 16h2q.425 0 .713.288T22 17v2q0 .425-.288.713T21 20h-2Z"/>`),
		g.Group(children),
	)
}