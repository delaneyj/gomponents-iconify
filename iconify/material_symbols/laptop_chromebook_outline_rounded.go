package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopChromebookOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v13h1q.425 0 .713.288T24 19q0 .425-.288.713T23 20H1q-.425 0-.713-.288T0 19q0-.425.288-.713T1 18h1Zm8.5 0h3q.2 0 .35-.15t.15-.35q0-.2-.15-.35T13.5 17h-3q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15ZM4 15h16V5H4v10Zm0 0V5v10Z"/>`),
		g.Group(children),
	)
}