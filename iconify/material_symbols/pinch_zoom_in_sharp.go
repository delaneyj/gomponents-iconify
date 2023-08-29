package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinchZoomInSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.55 12.5L.5 11.45L3.425 8.5H1V7h5v5H4.5V9.575L1.55 12.5ZM7 6V1h1.5v2.45L11.425.5L12.5 1.575L9.55 4.5H12V6H7Zm6.575 17L7.6 17l1.575-1.625l2.825.8V7h2v8h1v-4h2v4h1v-3h2v3h1v-1h2v9h-9.425Z"/>`),
		g.Group(children),
	)
}