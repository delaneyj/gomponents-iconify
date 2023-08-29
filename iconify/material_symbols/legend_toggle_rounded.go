package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegendToggleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.925 7.4l-4.4 2.675q-.5.3-1.012.013T4 9.225q0-.275.125-.513t.35-.362L10 5l5 3.55l3.425-2.425q.5-.35 1.038-.075t.537.875q0 .25-.112.475t-.313.35L15 11L9.925 7.4ZM5 15q-.425 0-.713-.288T4 14q0-.425.288-.713T5 13h14q.425 0 .713.288T20 14q0 .425-.288.713T19 15H5Zm0 4q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h14q.425 0 .713.288T20 18q0 .425-.288.713T19 19H5Z"/>`),
		g.Group(children),
	)
}