package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopWindowsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20q-.425 0-.713-.288T0 19q0-.425.288-.713T1 18h3v-1q-.825 0-1.413-.588T2 15V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v10q0 .825-.588 1.413T20 17v1h3q.425 0 .713.288T24 19q0 .425-.288.713T23 20H1Zm3-5h16V5H4v10Zm0 0V5v10Z"/>`),
		g.Group(children),
	)
}