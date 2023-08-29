package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneMissedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.4 20.4l-2.3-2.25q-.3-.3-.3-.7t.3-.7q2.2-2.375 5.075-3.562T12 12q2.95 0 5.813 1.188T22.9 16.75q.3.3.3.7t-.3.7l-2.3 2.25q-.275.275-.638.3t-.662-.2l-2.9-2.2q-.2-.15-.3-.35t-.1-.45v-2.85q-.95-.3-1.95-.475T12 14q-1.05 0-2.05.175T8 14.65v2.85q0 .25-.1.45t-.3.35l-2.9 2.2q-.3.225-.663.2t-.637-.3ZM6 9q-.425 0-.713-.287T5 8V4q0-.425.288-.713T6 3h4q.425 0 .713.288T11 4q0 .425-.288.713T10 5H8.4l3.525 3.525l4.95-4.95q.3-.3.713-.3t.712.3q.3.3.3.713T18.3 5l-4.925 4.925q-.575.575-1.425.575t-1.425-.575L7 6.4V8q0 .425-.288.713T6 9Z"/>`),
		g.Group(children),
	)
}