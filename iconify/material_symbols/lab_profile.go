package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabProfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12h8v-2H8v2Zm0-4h8V6H8v2Zm11.95 12.475L15.9 15.2q-.425-.575-1.05-.887T13.5 14H4V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .125-.013.238t-.037.237ZM6 22q-.825 0-1.412-.588T4 20v-4h9.5q.25 0 .463.113t.362.312l4.2 5.5q-.125.05-.262.063T18 22H6Z"/>`),
		g.Group(children),
	)
}