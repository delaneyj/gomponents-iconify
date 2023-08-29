package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4a1 1 0 0 1 2 0v24a1 1 0 1 1-2 0V4Zm7.5 1A3.5 3.5 0 0 0 7 8.5v3a3.5 3.5 0 0 0 3.5 3.5h15a3.5 3.5 0 0 0 3.5-3.5v-3A3.5 3.5 0 0 0 25.5 5h-15Zm0 12A3.5 3.5 0 0 0 7 20.5v3a3.5 3.5 0 0 0 3.5 3.5h9a3.5 3.5 0 0 0 3.5-3.5v-3a3.5 3.5 0 0 0-3.5-3.5h-9Z"/>`),
		g.Group(children),
	)
}