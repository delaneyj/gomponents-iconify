package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakfastDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.825 0-1.413-.588T4 19v-8.55q-.925-.55-1.463-1.462T2 7q0-1.65 1.175-2.825T6 3h12q1.65 0 2.825 1.175T22 7q0 1.075-.537 1.988T20 10.45V19q0 .825-.588 1.413T18 21H6Zm5.3-4.3q.3.275.713.275t.687-.275l3-3q.3-.3.3-.713t-.3-.687l-3-3q-.275-.3-.687-.3t-.713.3l-3 3q-.275.275-.275.688t.275.712l3 3Zm.7-2.1L10.4 13l1.6-1.6l1.6 1.6l-1.6 1.6Z"/>`),
		g.Group(children),
	)
}