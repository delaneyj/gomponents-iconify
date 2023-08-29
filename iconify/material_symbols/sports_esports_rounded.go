package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsEsportsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.55 19q-1.275 0-1.975-.888T2.05 15.95l1.05-7.5q.225-1.5 1.338-2.475T7.05 5h9.9q1.5 0 2.613.975T20.9 8.45l1.05 7.5q.175 1.275-.525 2.163T19.45 19q-.525 0-.975-.188t-.825-.562L15.4 16H8.6l-2.25 2.25q-.375.375-.825.563T4.55 19ZM17 13q.425 0 .713-.288T18 12q0-.425-.288-.713T17 11q-.425 0-.713.288T16 12q0 .425.288.713T17 13Zm-2-3q.425 0 .713-.288T16 9q0-.425-.288-.713T15 8q-.425 0-.713.288T14 9q0 .425.288.713T15 10Zm-7.25 1.25v1q0 .325.213.537T8.5 13q.325 0 .537-.213t.213-.537v-1h1q.325 0 .537-.213T11 10.5q0-.325-.213-.537t-.537-.213h-1v-1q0-.325-.213-.537T8.5 8q-.325 0-.537.213t-.213.537v1h-1q-.325 0-.537.213T6 10.5q0 .325.213.537t.537.213h1Z"/>`),
		g.Group(children),
	)
}