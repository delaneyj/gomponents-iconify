package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14q-.825 0-1.413-.588T5 12q0-.825.588-1.413T7 10q.825 0 1.413.588T9 12q0 .825-.588 1.413T7 14Zm0 4q-2.5 0-4.25-1.75T1 12q0-2.5 1.75-4.25T7 6q1.675 0 3.038.825T12.2 9h8.375q.2 0 .388.075t.337.225l2 2q.15.15.212.325t.063.375q0 .2-.063.375t-.212.325l-3.175 3.175q-.125.125-.3.2t-.35.1q-.175.025-.35-.025t-.325-.175L17.5 15l-1.425 1.075q-.125.1-.275.15t-.3.05q-.15 0-.313-.05t-.287-.15L13.375 15H12.2q-.8 1.35-2.163 2.175T7 18Zm0-2q1.4 0 2.463-.85T10.875 13H14l1.45 1.025v.013v-.013L17.5 12.5l1.775 1.375L21.15 12h-.013h.013l-1-1v-.013V11h-9.275q-.35-1.3-1.413-2.15T7 8Q5.35 8 4.175 9.175T3 12q0 1.65 1.175 2.825T7 16Z"/>`),
		g.Group(children),
	)
}