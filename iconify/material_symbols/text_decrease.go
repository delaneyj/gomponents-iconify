package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19L6.25 5h2.5L14 19h-2.4l-1.275-3.575h-5.65L3.4 19H1Zm4.4-5.6h4.2L7.55 7.6h-.1L5.4 13.4ZM15 13v-2h8v2h-8Z"/>`),
		g.Group(children),
	)
}