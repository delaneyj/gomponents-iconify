package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularConnectedNoInternetZeroBarOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20v2H2L22 2v6h-2V6.825L6.825 20H18Zm3-2q-.425 0-.713-.288T20 17v-6q0-.425.288-.713T21 10q.425 0 .713.288T22 11v6q0 .425-.288.713T21 18Zm0 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Z"/>`),
		g.Group(children),
	)
}