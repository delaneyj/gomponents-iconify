package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22v-7h2v7h-2Zm4 0v-7h2v7h-2ZM2 22L22 2v11h-9v9H2Z"/>`),
		g.Group(children),
	)
}