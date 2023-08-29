package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerPlugOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 21v-3L6 14.5V7.075h1.075v2.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4l-4.95-4.95l-.35.35v3h-5Zm7.65-6.7L8 5.15V3h2v4h4V3h2v4h2v6.45l-.85.85Z"/>`),
		g.Group(children),
	)
}