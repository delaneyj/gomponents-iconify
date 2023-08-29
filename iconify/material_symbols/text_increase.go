package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19L6.25 5h2.5L14 19h-2.4l-1.275-3.575h-5.65L3.4 19H1Zm4.4-5.6h4.2L7.55 7.6h-.1L5.4 13.4ZM18 16v-3h-3v-2h3V8h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}