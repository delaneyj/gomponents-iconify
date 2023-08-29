package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrunchDiningOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21v-1h14v1q0 .425-.288.713T15 22H3Zm-1-4v-1q0-.425.288-.713T3 16h4v-1q0-.425.288-.713T8 14h2q.425 0 .713.288T11 15v1h4q.425 0 .713.288T16 17v1H2Zm17 4q-.425 0-.713-.288T18 21v-5.1q-.9-1.025-1.45-2.025T16 11.45V3q0-.425.288-.713T17 2h4q.425 0 .713.288T22 3v8.45q0 1.425-.537 2.438T20 15.9V20h1q.425 0 .713.288T22 21q0 .425-.288.713T21 22h-2ZM18 8h2V4h-2v4Zm1 6.05q.45-.525.725-1.2t.275-1.4V10h-2v1.45q0 .725.25 1.4t.75 1.2Zm0 0Z"/>`),
		g.Group(children),
	)
}