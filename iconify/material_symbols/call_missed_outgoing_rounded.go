package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallMissedOutgoingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 10.425L13.425 16q-.3.3-.675.45t-.75.15q-.375 0-.75-.15t-.675-.45L3.7 9.125q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3L12 14.6L17.6 9H14q-.425 0-.713-.288T13 8q0-.425.288-.713T14 7h6q.425 0 .713.288T21 8v6q0 .425-.288.713T20 15q-.425 0-.713-.288T19 14v-3.575Z"/>`),
		g.Group(children),
	)
}