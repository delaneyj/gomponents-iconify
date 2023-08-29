package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularPauseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22L22 2v11h-2V6.825L6.825 20H13v2H2Zm13 0v-7h2v7h-2Zm4 0v-7h2v7h-2Zm-5.575-8.6Z"/>`),
		g.Group(children),
	)
}