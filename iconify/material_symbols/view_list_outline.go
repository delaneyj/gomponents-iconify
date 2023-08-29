package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18h11v-2.675H9V18ZM4 8.675h3V6H4v2.675Zm0 4.675h3v-2.675H4v2.675ZM4 18h3v-2.675H4V18Zm5-4.65h11v-2.675H9v2.675Zm0-4.675h11V6H9v2.675ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}