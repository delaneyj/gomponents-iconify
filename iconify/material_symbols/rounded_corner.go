package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedCorner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm0-4V3h2v2H3Zm4 16v-2h2v2H7ZM7 5V3h2v2H7Zm4 16v-2h2v2h-2Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Zm0-4v-2h2v2h-2Zm2-4h-2V8q0-1.25-.875-2.125T16 5h-5V3h5q2.075 0 3.538 1.463T21 8v5Z"/>`),
		g.Group(children),
	)
}