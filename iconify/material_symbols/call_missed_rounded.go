package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallMissedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 10.425V14q0 .425-.288.713T4 15q-.425 0-.713-.288T3 14V8q0-.425.288-.713T4 7h6q.425 0 .713.288T11 8q0 .425-.288.713T10 9H6.4l5.6 5.6l6.9-6.9q.3-.3.7-.287t.7.312q.275.3.287.7t-.287.7L13.425 16q-.3.3-.675.45t-.75.15q-.375 0-.75-.15t-.675-.45L5 10.425Z"/>`),
		g.Group(children),
	)
}