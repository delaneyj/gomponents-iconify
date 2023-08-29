package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTenAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.75 17.225h1.8v-8.45h-3V10.6h1.2v6.625Zm4.45 0h1.2q.75 0 1.287-.525t.538-1.275V10.6q0-.75-.537-1.287T14.4 8.775h-1.2q-.75 0-1.275.537T11.4 10.6v4.825q0 .75.525 1.275t1.275.525Zm0-1.8V10.6h1.2v4.825h-1.2ZM9 3V1h6v2H9Zm3 19q-1.85 0-3.488-.713T5.65 19.35q-1.225-1.225-1.938-2.863T3 13q0-1.85.713-3.488T5.65 6.65q1.225-1.225 2.863-1.938T12 4q1.55 0 2.975.5t2.675 1.45l1.4-1.4l1.4 1.4l-1.4 1.4Q20 8.6 20.5 10.025T21 13q0 1.85-.713 3.488T18.35 19.35q-1.225 1.225-2.863 1.938T12 22Z"/>`),
		g.Group(children),
	)
}