package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 21l-3.5-3.45l1.45-1.4L5 17.2V4h2v13.2l1.05-1.05l1.4 1.4L6 21Zm4-3.9v-1.9l2.8-.95V9.8l-2.8-1V6.9L21 11v2l-11 4.1Zm4.4-3.45l4.55-1.6v-.1l-4.55-1.6v3.3Z"/>`),
		g.Group(children),
	)
}