package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrontHandRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 23q-3.55 0-6.025-2.475T4 14.5V5q0-.425.288-.713T5 4q.425 0 .713.288T6 5v6h2V3q0-.425.288-.713T9 2q.425 0 .713.288T10 3v8h2V2q0-.425.288-.713T13 1q.425 0 .713.288T14 2v9h2V4q0-.425.288-.713T17 3q.425 0 .713.288T18 4v5.75q-.725.525-1.113 1.325T16.5 12.75V14h-1.25q-1.575 0-2.663 1.088T11.5 17.75V19q0 .325.213.537t.537.213q.325 0 .537-.213T13 19v-1.25q0-.95.65-1.6t1.6-.65H17q.425 0 .713-.288T18 14.5v-1.75q0-.95.65-1.6t1.6-.65q.325 0 .537.213t.213.537v3.25q0 3.55-2.475 6.025T12.5 23Z"/>`),
		g.Group(children),
	)
}