package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 22v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM2 22L21.975 2.025V12.65q-.65-.4-1.425-.588T19 11.875q-2.55 0-4.338 1.787T12.875 18q0 1.15.375 2.138T14.375 22H2Z"/>`),
		g.Group(children),
	)
}