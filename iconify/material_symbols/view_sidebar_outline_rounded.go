package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSidebarOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4ZM17.5 8.675H20V6h-2.5v2.675Zm0 4.65H20v-2.65h-2.5v2.65ZM4 18h11.5V6H4v12Zm13.5 0H20v-2.675h-2.5V18Z"/>`),
		g.Group(children),
	)
}