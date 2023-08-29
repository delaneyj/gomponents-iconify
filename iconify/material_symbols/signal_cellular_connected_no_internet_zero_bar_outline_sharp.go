package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularConnectedNoInternetZeroBarOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20v2H2L22 2v6h-2V6.825L6.825 20H18Zm2-2v-8h2v8h-2Zm0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}