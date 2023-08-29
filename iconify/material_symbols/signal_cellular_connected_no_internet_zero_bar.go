package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularConnectedNoInternetZeroBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22L22 2v6h-2V6.85L6.85 20H18v2H2Zm19 0q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Zm-1-4v-8h2v8h-2Z"/>`),
		g.Group(children),
	)
}