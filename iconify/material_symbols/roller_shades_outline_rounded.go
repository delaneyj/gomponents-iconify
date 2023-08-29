package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollerShadesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19V5q0-.825.588-1.413T6 3h12q.825 0 1.413.588T20 5v14h1q.425 0 .713.288T22 20q0 .425-.288.713T21 21H3q-.425 0-.713-.288T2 20q0-.425.288-.713T3 19h1Zm2-8h12V5H6v6Zm0 8h12v-6h-5v1.8q.35.25.55.625t.2.825q0 .725-.513 1.238T12 18q-.725 0-1.238-.513t-.512-1.237q0-.45.2-.813t.55-.612V13H6v6ZM6 5h12H6Z"/>`),
		g.Group(children),
	)
}