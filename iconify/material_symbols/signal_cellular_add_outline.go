package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularAddOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22L21.975 2.025V12.65q-.45-.275-.95-.438t-1.05-.262v-5.1L6.825 20H13.2q.2.575.5 1.075t.675.925H2Zm16 0v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM6.825 20l13.15-13.15l-3.413 3.413l-3.025 3.024l-3.087 3.088L6.825 20Z"/>`),
		g.Group(children),
	)
}