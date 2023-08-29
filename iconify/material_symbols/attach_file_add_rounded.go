package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachFileAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 22q-2.3 0-3.9-1.6T6 16.5V6q0-1.65 1.175-2.825T10 2q1.65 0 2.825 1.175T14 6v8h-1.5V6q0-1.05-.725-1.775T10 3.5q-1.05 0-1.775.725T7.5 6v10.5q0 1.65 1.175 2.825T11.5 20.5q.725 0 1.363-.238T14 19.6v1.8q-.575.275-1.2.438T11.5 22Zm4.5-4h-2q-.425 0-.713-.288T13 17q0-.425.288-.713T14 16h2v-2q0-.425.288-.713T17 13q.425 0 .713.288T18 14v2h2q.425 0 .713.288T21 17q0 .425-.288.713T20 18h-2v2q0 .425-.288.713T17 21q-.425 0-.713-.288T16 20v-2Zm-4.5-1.5V18q-1.05 0-1.775-.725T9 15.5V6.75q0-.325.213-.537T9.75 6q.325 0 .537.213t.213.537v8.75q0 .425.288.713t.712.287Zm4-5.5V6.75q0-.325.213-.537T16.25 6q.325 0 .537.213T17 6.75V11h-1.5Z"/>`),
		g.Group(children),
	)
}