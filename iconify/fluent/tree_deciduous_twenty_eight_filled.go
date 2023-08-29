package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeDeciduousTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.21 5.562a5.002 5.002 0 0 1 9.58 0a5.001 5.001 0 0 1 4.021 6.303A5 5 0 0 1 20 21h-3v3a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-3H8a5 5 0 0 1-2.811-9.135A5.001 5.001 0 0 1 9.21 5.562ZM12.5 24a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-3h-3v3Z"/>`),
		g.Group(children),
	)
}