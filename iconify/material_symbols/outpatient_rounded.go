package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutpatientRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21q-.425 0-.713-.288T1 20V4q0-.425.288-.713T2 3h12q.425 0 .713.288T15 4v16q0 .425-.288.713T14 21h-4v-5H6v5H2Zm3-8h2v-2H5v2Zm0-4h2V7H5v2Zm4 4h2v-2H9v2Zm0-4h2V7H9v2Zm7 3q0-.425.288-.713T17 11h2.175l-.4-.4q-.275-.275-.275-.688t.3-.712q.275-.275.7-.275t.7.275l2.1 2.1q.3.3.3.7t-.3.7l-2.1 2.1q-.275.275-.687.287T18.8 14.8q-.275-.275-.288-.688t.263-.712l.4-.4H17q-.425 0-.713-.288T16 12Z"/>`),
		g.Group(children),
	)
}