package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15q-.425 0-.713-.288T10 14v-4q0-.425.288-.713T11 9h3q.425 0 .713.288T15 10v4q0 .425-.288.713T14 15h-3Zm.5-1.5h2v-3h-2v3Zm5.5 3.75V15.5q0-.425.288-.713T18 14.5h2v-1h-2.275q-.3 0-.513-.213T17 12.75q0-.3.213-.525T17.75 12h2.75q.425 0 .713.288T21.5 13v1.5q0 .425-.288.713t-.712.287h-2v1h2.25q.3 0 .525.225t.225.525q0 .325-.225.537T20.75 18h-3q-.325 0-.537-.213T17 17.25ZM4 15q-.425 0-.713-.288T3 14v-4q0-.425.288-.713T4 9h3q.425 0 .713.288T8 10v.525q0 .3-.225.513t-.525.212q-.325 0-.537-.213T6.5 10.5h-2v3h2q0-.3.213-.525t.537-.225q.3 0 .525.225T8 13.5v.5q0 .425-.288.713T7 15H4Z"/>`),
		g.Group(children),
	)
}