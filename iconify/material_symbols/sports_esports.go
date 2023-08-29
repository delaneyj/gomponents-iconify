package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsEsports(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.55 19q-1.275 0-1.975-.888T2.05 15.95l1.05-7.5q.225-1.5 1.338-2.475T7.05 5h9.9q1.5 0 2.613.975T20.9 8.45l1.05 7.5q.175 1.275-.525 2.163T19.45 19q-.525 0-.975-.188t-.825-.562L15.4 16H8.6l-2.25 2.25q-.375.375-.825.563T4.55 19ZM17 13q.425 0 .713-.288T18 12q0-.425-.288-.713T17 11q-.425 0-.713.288T16 12q0 .425.288.713T17 13Zm-2-3q.425 0 .713-.288T16 9q0-.425-.288-.713T15 8q-.425 0-.713.288T14 9q0 .425.288.713T15 10Zm-7.25 3h1.5v-1.75H11v-1.5H9.25V8h-1.5v1.75H6v1.5h1.75V13Z"/>`),
		g.Group(children),
	)
}