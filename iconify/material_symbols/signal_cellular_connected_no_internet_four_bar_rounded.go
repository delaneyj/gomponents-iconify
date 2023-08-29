package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularConnectedNoInternetFourBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18q-.425 0-.713-.288T20 17v-6q0-.425.288-.713T21 10q.425 0 .713.288T22 11v6q0 .425-.288.713T21 18Zm0 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22ZM4.425 22q-.675 0-.937-.613T3.7 20.3L20.3 3.7q.475-.475 1.088-.213t.612.938V8h-2q-.825 0-1.413.588T18 10v12H4.425Z"/>`),
		g.Group(children),
	)
}