package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 5.5A3.5 3.5 0 0 1 12.5 2h7A3.5 3.5 0 0 1 23 5.5v21a3.5 3.5 0 0 1-3.5 3.5h-7A3.501 3.501 0 0 1 9 26.5v-21Zm2 .5v2h3a1 1 0 1 0 0-2h-3Zm0 4.5v2h5a1 1 0 1 0 0-2h-5Zm0 4.5v2h3a1 1 0 1 0 0-2h-3Zm0 4.5v2h5a1 1 0 1 0 0-2h-5Zm0 4.5v2h3a1 1 0 1 0 0-2h-3Z"/>`),
		g.Group(children),
	)
}