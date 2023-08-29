package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivateConnectivitySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q-2.65 0-4.612-1.713T5.075 13H2v-2h3.075q.35-2.575 2.313-4.288T12 5q2.65 0 4.612 1.713T18.925 11H22v2h-3.075q-.35 2.575-2.313 4.288T12 19Zm-3-3.5h6v-5h-1v-.9q0-.875-.575-1.488T12 7.5q-.825 0-1.413.588T10 9.5v1H9v5Zm3-1.75q-.325 0-.537-.213T11.25 13q0-.325.213-.537T12 12.25q.325 0 .537.213t.213.537q0 .325-.213.537T12 13.75Zm-1-3.25v-1q0-.425.288-.713T12 8.5q.425 0 .713.288T13 9.5v1h-2Z"/>`),
		g.Group(children),
	)
}